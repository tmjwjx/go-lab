package main

import "fmt"

func main() {
	diySlice := make([]int, 0, 2)
	diySlice = append(diySlice, 8)

	//观察diySlice3
	diySlice3 := append(diySlice, 1)
	//diySlice 变化
	//查看输出切片的变化，为什么和直接输出结果不一样
	fmt.Println("diySlice内容下标", diySlice[0:2])
	//查看输出切片的变化
	fmt.Println("diySlice 内容", diySlice)
	//查看长度和容量
	fmt.Printf("diySlice-->容量%d 长度%d\n", cap(diySlice), len(diySlice))

	fmt.Println("diySlice3 内容", diySlice3)
	fmt.Printf("diySlice3-->容量%d 长度%d\n", cap(diySlice3), len(diySlice3))

	//观察diySlice2
	diySlice2 := append(diySlice, 8)
	fmt.Println("diySlice2 内容", diySlice2)
	fmt.Printf("diySlice2-->容量%d 长度%d\n", cap(diySlice2), len(diySlice2))
	fmt.Println(diySlice[0:2])

}