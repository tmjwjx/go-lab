package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	// 创建一个新的路由器
	router := mux.NewRouter()
	// 启动 hub 的运行 goroutine
	go h.run()
	// 设置 WebSocket 路由
	router.HandleFunc("/ws", myws)
	// 启动 HTTP 服务器，监听 127.0.0.1:8080
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		// 如果服务器启动失败，打印错误信息
		fmt.Println("err:", err)
	}
}
