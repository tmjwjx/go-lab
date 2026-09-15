package main

import (
	"fmt"
	"unsafe"
)

func main() {
	sli1 := []int{0, 1, 2, 3, 4, 5, 6}
	sli2 := sli1[2:5]
	fmt.Println(sli2)
	fmt.Println(len(sli2), cap(sli2))
	fmt.Println(&sli1[2] == &sli2[0])
	p := unsafe.Pointer(&sli2[2])
	i := 1
	// 地址下一位
	add := uintptr(p) + uintptr(i)*unsafe.Sizeof(sli2[0])
	fmt.Println(*(*int)(unsafe.Pointer(add)))
	fmt.Println(sli2[2])
	sli1 = append(sli1, 7)
	fmt.Println(&sli1[2] == &sli2[0])
	fmt.Println(sli2[3])

}

//func main() {
//	//runtime.GOMAXPROCS(1)
//	go func(s string) {
//		for i := 0; i < 2; i++ {
//			fmt.Println(s)
//		}
//	}("world")
//	// 主协程
//	for i := 0; i < 2; i++ {
//		// 切一下，再次分配任务
//		runtime.Gosched()
//		fmt.Println("hello")
//	}
//	time.Sleep(1 * time.Second)
//}

//func main() {
//	go func() {
//		defer fmt.Println("A.defer")
//		func() {
//			defer fmt.Println("B.defer")
//			runtime.Goexit()
//			defer fmt.Println("C.defer")
//			fmt.Println("B")
//		}()
//		fmt.Println("A")
//	}()
//
//	fmt.Println("main")
//	time.Sleep(1 * time.Second)
//}
