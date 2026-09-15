package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

const (
	socks5Ver = 0x05
	cmdBind   = 0x01
	atypeIPV4 = 0x01
	atypeHOST = 0x03
	atypeIPV6 = 0x04
)

func process(conn net.Conn) {
	defer conn.Close()              // 确保函数退出时关闭连接
	reader := bufio.NewReader(conn) // 为连接创建一个新的缓冲读取器

	err := auth(reader, conn) // 执行认证
	if err != nil {
		log.Printf("client %v auth failed:%v", conn.RemoteAddr(), err) // 记录认证失败
		return
	}
	log.Println("auth success") // 记录认证成功

}

// auth 用于socks5协议的认证
func auth(reader *bufio.Reader, conn net.Conn) (err error) {
	// +----+----------+----------+
	// |VER | NMETHODS | METHODS  |
	// +----+----------+----------+
	// | 1  |    1     | 1 to 255 |
	// +----+----------+----------+
	// VER: 协议版本，socks5为0x05
	// NMETHODS: 支持认证的方法数量
	// METHODS: 对应NMETHODS，NMETHODS的值为多少，METHODS就有多少个字节。RFC预定义了一些值的含义，内容如下:
	// X’00’ NO AUTHENTICATION REQUIRED
	// X’02’ USERNAME/PASSWORD

	// 读取并解析客户端发送的协商数据
	ver, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read ver failed:%w", err)
	}
	// 检查协议版本是否为socks5
	if ver != socks5Ver {
		return fmt.Errorf("not supported ver:%v", ver)
	}
	// 读取并解析客户端支持的认证方式
	methodSize, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read methodSize failed:%w", err)
	}
	// 读取客户端支持的认证方式
	method := make([]byte, methodSize)
	_, err = io.ReadFull(reader, method)
	if err != nil {
		return fmt.Errorf("read method failed:%w", err)
	}
	log.Println("ver", ver, "method", method)
	// +----+--------+
	// |VER | METHOD |
	// +----+--------+
	// | 1  |   1    |
	// +----+--------+
	_, err = conn.Write([]byte{socks5Ver, 0x00})
	if err != nil {
		return fmt.Errorf("write failed:%w", err)
	}
	return nil
}

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err)
	}
	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}
		go process(client)
	}
}
