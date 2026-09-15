package main

import (
	"fmt"
	"library/chat_room/server/model"
	"net"
	"time"
)

// process
// @Description: 和客户端保持通信
// @param        conn net.Conn
// @Author tianjiajie 2024-12-03 21:27:21
func process(conn net.Conn) {
	//这里需要延时关闭conn
	defer conn.Close()

	// 消费队列
	go ListenMessageQueue()

	processor := &Processor{
		Conn: conn,
	}

	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务端通讯协程错误err=", err)
		return
	}

}

func init() {
	//当服务器启动时，我们就去初始化我们的redis的连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
}

// 这里编写一个函数，完成对UserDao的初始化任务
func initUserDao() {
	//这里的pool本身就是一个全局变量（在redis.go中定义了）
	//这里需要注意一个初始化顺序问题
	//initPool，再initUserDao
	model.MyUserDao = model.NewUserDao(pool)
}

// main
// @Description: 服务器端
// @Author tianjiajie 2024-12-03 21:27:34
func main() {

	fmt.Println("服务器在8888端口监听")
	listen, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close()

	go serverController()

	//监听成功，等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器......")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		//连接成功，启动一个协程和客户端保持通信
		go process(conn)
	}
}
