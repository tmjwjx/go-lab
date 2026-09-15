package util

import (
	"fmt"
)

type Animal interface {
	Eat(food Food)
}

type Food interface {
	getName() string
	getSize() int
}

type Bread struct {
	Name string
	Size int
}

func (bread Bread) getName() string {
	return bread.Name
}

func (bread Bread) getSize() int {
	return bread.Size
}

type Dog struct {
	Name   string
	Weight int
}

func (dog *Dog) Eat(food Food) {
	fmt.Println(dog.Name, "eat:", food.getName())
	dog.Weight += food.getSize()
	fmt.Println(dog.Name, "weight:", dog.Weight)
	return
}

//func main() {
//
//	d := Dog{
//		UserName:   "李壮",
//		Weight: 100,
//	}
//	b := Bread{
//		UserName: "break",
//		Size: 10,
//	}
//	d.Eat(b)
//
//}
