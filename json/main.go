package main

import (
	"encoding/json"
	"fmt"
)

type userInfo struct {
	Name  string
	Age   int `json:"age"`
	Hobby []string
}

func main() {
	// 创建一个结构体实例
	user := userInfo{
		Name:  "小王子",
		Age:   18,
		Hobby: []string{"篮球", "足球", "双色球"},
	}

	buf, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println("buf = ", buf)
	fmt.Println("string(buf) = ", string(buf))

	buf, err = json.MarshalIndent(user, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println("string(buf) = ", string(buf))

	var b userInfo
	err = json.Unmarshal(buf, &b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("b = %+v\n", b)
	fmt.Printf("b = %#v\n", b)
}
