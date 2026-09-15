package test

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func TestExample1(t *testing.T) {
	// 打开文件
	file, err := os.Open("../file/file1.txt")
	if err != nil {
		fmt.Println("open file err =", err)
	}

	// 输出文件，看看文件是什么
	fmt.Printf("file=%v\n", file)

	// 创建一个*Reader，带缓存
	reader := bufio.NewReader(file)

	// 循环读取文件内容
	for {
		str, err := reader.ReadString('\n') // 读到换行结束
		fmt.Print(str)

		if err == io.EOF {
			break
		}
	}
	fmt.Println("文件读取结束...")

	// 关闭文件
	defer file.Close() // 及时关闭，否则会内存泄漏
}

func TestExample2(t *testing.T) {
	// 使用 ioutil.ReadFile 一次性将文件读取到内存
	file := "../file/file1.txt"
	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err=%v", err)
	}

	// 输出文件内容
	fmt.Printf("%v\n", content)       // 字节形式
	fmt.Printf("%v", string(content)) // 字符串形式
}
