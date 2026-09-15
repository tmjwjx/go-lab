package main

import (
	"fmt"
	"strconv"
)

type Commodity map[string]map[string]string

func (commodity *Commodity) AddToCart() {
	var name, color string
	var priceFloat float64
	var numberInt int64
	fmt.Print("输入商品名称：")
	fmt.Scanln(&name)
	fmt.Print("输入商品价格：")
	fmt.Scanln(&priceFloat)
	fmt.Print("输入商品数量：")
	fmt.Scanln(&numberInt)
	fmt.Print("输入商品颜色：")
	fmt.Scanln(&color)
	if (*commodity)[name] != nil {
		number := strconv.FormatInt(numberInt, 64)
		//m := (*commodity)[name]
		//m["number"] = number
		(*commodity)[name]["number"] = number
		fmt.Println("商品已存在，数量已更新")
		return
	}
	price := strconv.FormatFloat(priceFloat, 'f', 2, 64)
	number := strconv.FormatInt(numberInt, 10)
	temp := map[string]string{
		"price":  price,
		"number": number,
		"color":  color,
	}
	(*commodity)[name] = temp
	
}
func (commodity *Commodity) RemoveFromCart() {
	var name string
	fmt.Print("输入商品名称：")
	fmt.Scanln(&name)
	if (*commodity)[name] == nil {
		fmt.Println("没有找到该商品")
		return
	}
	delete(*commodity, name)
}
func (commodity *Commodity) CalculateTotalPrice() {
	sumPrice := 0.0
	var sumNumber int64 = 0
	for _, value := range *commodity {
		priceFloat, _ := strconv.ParseFloat(value["price"], 64)
		numberInt, _ := strconv.ParseInt(value["number"], 10, 64)
		sumPrice += priceFloat * float64(numberInt)
		sumNumber += numberInt
	}
	fmt.Println("商品总数：", sumNumber)
	fmt.Println("商品总价：", sumPrice)
}
func (commodity *Commodity) EmptyCart() {
	for key := range *commodity {
		delete(*commodity, key)
	}
}
func (commodity *Commodity) ShowCart() {
	for key, value := range *commodity {
		fmt.Println("商品名称：", key, "\t价格：", value["price"], "\t数量：", value["number"], "\t颜色：", value["color"])
	}
}

func main() {
	commodity := make(Commodity)
	temp := map[string]string{
		"price":  "2499.00",
		"number": "4",
		"color":  "pink",
	}
	commodity["huawei"] = temp
	var sel int
	for {
		fmt.Println("------------------")
		fmt.Println("选择操作选项")
		fmt.Println("1.添加商品到购物车")
		fmt.Println("2.从购物车中移除商品")
		fmt.Println("3.计算购物车商品信息")
		fmt.Println("4.清空购物车")
		fmt.Println("5.展示购物车内容")
		fmt.Println("0.退出")
		fmt.Scanln(&sel)
		switch sel {
		case 1:
			commodity.AddToCart()
		case 2:
			commodity.RemoveFromCart()
		case 3:
			commodity.CalculateTotalPrice()
		case 4:
			commodity.EmptyCart()
		case 5:
			commodity.ShowCart()
		case 0:
			fmt.Println("已退出")
			return
		default:
			fmt.Println("选项无效，请重新输入")
			
		}
	}
	
}