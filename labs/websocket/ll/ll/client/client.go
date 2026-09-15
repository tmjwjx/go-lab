package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

type User struct {
	Username      string
	OtherUsername string
	Msg           string
	ServerMsg     string
}

var (
	user = new(User)
	wg   sync.WaitGroup
)

func main() {
	wg.Add(1)
	//fmt.Println("请输入用户名：")
	//fmt.Scanln(&user.Username)
	user.Username = "ll"
	//fmt.Println("请输入要给谁发送信息")
	//fmt.Scanln(&user.OtherUsername)
	user.OtherUsername = "tjj"

	conn, _ := net.Dial("tcp4", "localhost:8899")

	//发送消息
	go func() {
		fmt.Println("请输入您要发送的消息，输入'exit'退出程序（只提示一次）:")
		for {
			fmt.Scanln(&user.Msg)
			if user.Msg == "exit" {
				conn.Close()
				wg.Done()
				os.Exit(0)
			}
			conn.Write([]byte(fmt.Sprintf("%s-%s-%s-%s", user.Username, user.OtherUsername, user.Msg, user.ServerMsg)))
		}
	}()

	go func() {
		for true {
			b := make([]byte, 1024)
			count, _ := conn.Read(b)
			array := strings.Split(string(b[:count]), "-")
			user2 := new(User)
			user2.Username = array[0]
			user2.OtherUsername = array[1]
			user2.Msg = array[2]
			user2.ServerMsg = array[3]
			if user2.ServerMsg != "" {
				fmt.Println("\t\t服务器的消息：", user2.ServerMsg)
			} else {
				fmt.Println("田家杰：", user2.Msg)
			}

		}
	}()

	wg.Wait()
}