package main

import (
	"fmt"
	"log"
	"net/http"
	websocket "websocket-main"
)

var UP = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var conns []*websocket.Conn

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := UP.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	conns = append(conns, conn)
	for {
		m, p, err := conn.ReadMessage() // m 为消息类型，p 为消息内容，err 为错误
		if err != nil {
			break
		}
		for i := range conns {
			conns[i].WriteMessage(websocket.TextMessage, []byte("你说的是：\""+string(p)+"\"吗"))
		}
		fmt.Println(m, string(p))
	}
	defer conn.Close()
	log.Println("服务关闭")

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8888", nil)

}