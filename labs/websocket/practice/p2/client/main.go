package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")
	conn, _ := net.DialTCP("tcp4", nil, addr)
	_, _ = conn.Write([]byte("客户端发送的数据"))
	b := make([]byte, 1024)
	count, _ := conn.Read(b)
	fmt.Println("服务器发送回来的消息为：", string(b[:count]))
	_ = conn.Close()
}
