package main

import (
	"fmt"
	"os"
)

func process() {

	// 获取命令行参数
	fmt.Println(os.Args)

	// 获取 GOROOT 环境变量
	fmt.Println("get =", os.Getenv("GOROOT"))

	// 设置环境变量 AA 的值为 BB， 仅在程序运行期间有效，并且只影响当前进程
	fmt.Println("set =", os.Setenv("AA", "BB"))

	//// 不懂什么意思
	//buf, err := exec.Command("grep", "127.0.0.1", "/etc/hosts").CombinedOutput()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(string(buf))

}
