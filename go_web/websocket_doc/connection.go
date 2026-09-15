package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

// connection 结构体表示一个 WebSocket 连接
type connection struct {
	ws   *websocket.Conn // WebSocket 连接
	sc   chan []byte     // 发送消息的通道
	data *Data           // 连接相关的数据
}

// WebSocket 升级器，配置读写缓冲区大小和跨域检查
var wu = &websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// WebSocket 处理函数
func myws(w http.ResponseWriter, r *http.Request) {
	// 将 HTTP 连接升级为 WebSocket 连接
	ws, err := wu.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	// 创建一个新的 connection 实例
	c := &connection{
		sc:   make(chan []byte, 256),
		ws:   ws,
		data: &Data{},
	}
	// 将连接加入到 hub 的注册通道
	h.r <- c
	// 启动 writer 协程
	go c.writer()
	// 启动 reader 函数
	c.reader()
	defer func() {
		// 连接关闭时，处理登出逻辑
		c.data.Type = "logout"
		user_list = del(user_list, c.data.User)
		c.data.UserList = user_list
		c.data.Content = c.data.User
		data_b, _ := json.Marshal(c.data)
		h.b <- data_b
		//h.r <- c
		h.u <- c
	}()
}

// writer 函数从 sc 通道读取消息并发送到 WebSocket 连接
func (c *connection) writer() {
	for message := range c.sc {
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
	c.ws.Close()
}

// 全局用户列表
var user_list = []string{}

// reader 函数从 WebSocket 连接读取消息并处理
func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			//h.r <- c // todo 为什么是 h.r <- c    不理解
			h.u <- c
			break
		}
		json.Unmarshal(message, &c.data)
		switch c.data.Type {
		case "login":
			// 处理登录消息
			c.data.User = c.data.Content
			c.data.From = c.data.User
			user_list = append(user_list, c.data.User)
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
		case "user":
			// 处理用户消息
			c.data.Type = "user"
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
		case "logout":
			// 处理登出消息
			c.data.Type = "logout"
			user_list = del(user_list, c.data.User)
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
			//h.r <- c
			h.u <- c
		default:
			fmt.Print("========default================")
		}
	}
}

// del 函数从用户列表中删除指定用户
func del(slice []string, user string) []string {
	count := len(slice)
	if count == 0 {
		return slice
	}
	if count == 1 && slice[0] == user {
		return []string{}
	}
	var n_slice []string
	for i := range slice {
		if slice[i] == user && i == count {
			return slice[:count]
		} else if slice[i] == user {
			n_slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	fmt.Println(n_slice)
	return n_slice
}
