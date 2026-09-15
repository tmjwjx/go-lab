package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 设置系统环境变量的命令
	varName := "AAAAAAAA"
	varValue := "BBBBBBBB"

	// 使用 reg 命令来设置系统环境变量
	cmd := exec.Command("reg", "add", "HKLM\\SYSTEM\\CurrentControlSet\\Control\\Session Manager\\Environment",
		"/v", varName, "/t", "REG_SZ", "/d", varValue, "/f")

	err := cmd.Run()
	if err != nil {
		fmt.Println("设置系统环境变量失败:", err)
		return
	}

	fmt.Println("成功设置系统环境变量", varName, "为", varValue)
}
