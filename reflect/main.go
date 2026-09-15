package main

import (
	"fmt"
	"reflect"
)

type f float64

type User struct {
	Id   int    `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
	Age  int    `json:"age" yaml:"age"`
}

func main() {

	var x f = 3.4
	t := reflect.TypeOf(x)
	fmt.Println("type:", t)

	v := reflect.ValueOf(x)
	fmt.Println("value:", v)

	k := t.Kind()
	fmt.Println("kind:", k)

	k = v.Kind()
	fmt.Println(k)

}

func Poni(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("类型：", t)
	fmt.Println("字符串类型：", t.Name())
	// 获取值
	v := reflect.ValueOf(o)
	fmt.Println(v.Kind())
	fmt.Println(v.Type())
	fmt.Println(v)
	// 可以获取所有属性
	// 获取结构体字段个数：t.NumField()
	for i := 0; i < t.NumField(); i++ {
		// 取每个字段
		f := t.Field(i)
		fmt.Printf("%s : %v\t", f.Name, f.Type)
		// 获取字段的值信息
		// Interface()：获取字段对应的值
		val := v.Field(i).Interface()
		fmt.Println("val :", val)
		fmt.Println(v.Field(i))
	}
	fmt.Println("=================方法====================")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
		fmt.Println(m.Type)
	}
}

func (u User) Hello(i int) (err error) {
	fmt.Println("Hello")
	return
}

func reflect_set_value(a interface{}) {
	v := reflect.ValueOf(a)
	k := v.Kind()
	switch k {
	case reflect.Float64:
		// 修改反射值
		v.SetFloat(6.9)
		fmt.Println("a is ", v.Float())
	case reflect.Ptr:
		// Elem()获取地址指向的值
		v.Elem().SetFloat(7.9)
		fmt.Println("case:", v.Elem().Float())
		// 地址
		fmt.Println(v.Pointer())
	}
}

func reflect_type(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Println("类型是：", t)
	// kind可以获取具体的类型
	k := t.Kind()
	fmt.Println("具体类型：", k)
	switch k {
	case reflect.Float64:
		fmt.Println("a is float64")
	case reflect.String:
		fmt.Println("a is string")
	}
}
