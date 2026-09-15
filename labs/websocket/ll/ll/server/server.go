package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Username      string
	OtherUsername string
	Msg           string
	ServerMsg     string
}

var (
	userMap = make(map[string]net.Conn)
	user    = new(User)
)

func main() {
	lis, _ := net.Listen("tcp", "localhost:8899")
	for {
		conn, _ := lis.Accept()
		go func() {
			for {
				b := make([]byte, 1024)
				count, _ := conn.Read(b)
				array := strings.Split(string(b[:count]), "-")
				if user == nil {
					user = new(User)
				}
				user.Username = array[0]
				user.OtherUsername = array[1]
				user.Msg = array[2]
				user.ServerMsg = array[3]
				userMap[user.Username] = conn

				if v, ok := userMap[user.OtherUsername]; ok && v != nil {
					n, err := v.Write([]byte(fmt.Sprintf("%s-%s-%s-%s", user.Username, user.OtherUsername, user.Msg, user.ServerMsg)))
					if n <= 0 || err != nil {
						delete(userMap, user.OtherUsername)
						conn.Close()
						v.Close()
						fmt.Println("if......")
						break
					}
					fmt.Println("消息发送成功")
				} else {
					user.ServerMsg = "对方不在线"
					conn.Write([]byte(fmt.Sprintf("%s-%s-%s-%s", user.Username, user.OtherUsername, user.Msg, user.ServerMsg)))
				}
			}
		}()

	}

}