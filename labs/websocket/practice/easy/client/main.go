package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	websocket "websocket-main"
)

func main() {
	dl := websocket.Dialer{}
	conn, _, err := dl.Dial("ws://127.0.0.1:8888", nil)
	if err != nil {
		log.Println(err)
		return
	}
	conn.WriteMessage(websocket.TextMessage, []byte("你好啊,我来了"))
	go send(conn)
	for {
		m, p, e := conn.ReadMessage()
		if e != nil {
			break
		}

		fmt.Println(m, string(p))
	}
}
func send(conn *websocket.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		l, _, _ := reader.ReadLine()
		conn.WriteMessage(websocket.TextMessage, l)
	}
}