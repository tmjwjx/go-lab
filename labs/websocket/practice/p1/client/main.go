package main

import "net"

func main() {
	//1.创建服务器地址
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")
	//2.创建连接
	conn, _ := net.DialTCP("tcp4", nil, addr)
	//3.发送数据
	conn.Write([]byte("客户端发送的数据"))
	//4.关闭连接
	conn.Close()
}