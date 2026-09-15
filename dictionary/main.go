package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

//func main() {
//
//	// 检查命令行参数的数量是否为 2（程序名和单词）
//	if len(os.Args) != 2 {
//		// 如果参数数量不为 2，打印用法提示并退出程序
//		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
//example: simpleDict hello
//    `)
//		os.Exit(1)
//	}
//
//	// 获取命令行参数中的单词
//	word := os.Args[1]
//
//	// 查询单词的字典信息
//	fmt.Println(query(word))
//
//}

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

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		b, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		input := string(b)
		output := query(input)
		fmt.Println(output)

		_, err = conn.Write([]byte(output))
		if err != nil {
			break
		}
	}
}
