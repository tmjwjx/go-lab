package main

import "encoding/json"

// 全局 hub 实例
var h = hub{
	c: make(map[*connection]bool),
	u: make(chan *connection),
	b: make(chan []byte),
	r: make(chan *connection),
}

// hub 结构体管理所有 WebSocket 连接
type hub struct {
	c map[*connection]bool // 活跃连接集合
	b chan []byte          // 广播消息通道
	r chan *connection     // 注册连接通道
	u chan *connection     // 注销连接通道
}

// run 方法处理 hub 的所有事件
func (h *hub) run() {
	for {
		select {
		case conn := <-h.r:
			// 处理新连接
			h.c[conn] = true
			conn.data.Ip = conn.ws.RemoteAddr().String()
			conn.data.Type = "handshake"
			conn.data.UserList = user_list
			data_b, _ := json.Marshal(conn.data)
			conn.sc <- data_b
		case conn := <-h.u:
			// 处理连接注销
			if _, ok := h.c[conn]; ok {
				delete(h.c, conn)
				close(conn.sc)
			}
		case data := <-h.b:
			// 处理广播消息
			for conn := range h.c {
				select {
				case conn.sc <- data:
				default:
					delete(h.c, conn)
					close(conn.sc)
				}
			}
		}
	}
}
