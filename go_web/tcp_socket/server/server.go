package main

import (
	"fmt"
	"net"
)

func main() {

	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			continue
		}

		go func(c net.Conn) {
			fmt.Fprintln(c, "Hello, Client!")
			fmt.Println("客户端ip地址为：", c.RemoteAddr())
			c.Close()
		}(conn)

	}
}
