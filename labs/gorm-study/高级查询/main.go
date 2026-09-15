package main

import (
	"fmt"
	"gorm-study"
)

type Student struct {
	ID     uint   `gorm:"size:3"`
	Name   string `gorm:"size:8"`
	Age    int    `gorm:"size:3"`
	Gender bool
	Email  *string `gorm:"size:32"`
}

func main() {
	DB := gorm_study.Initialization()
	DB.AutoMigrate(&Student{})
	
	type AggeGroup struct {
		Gender int
		Count  int    `gorm:"column:count(id)"`
		Name   string `gorm:"column:group_concat(name)"`
	}
	
	var agge []AggeGroup
	// 查询男生的个数和女生的个数
	DB.Table("students").Select("count(id)", "gender", "group_concat(name)").Group("gender").Scan(&agge)
	fmt.Println(agge)
	
	//var ageList []int
	//DB.Table("students").Select("age").Distinct("age").Scan(&ageList)
	//fmt.Println(ageList)
	
	//var users []Student
	//// 一页两条，第1页
	//DB.Limit(2).Offset(0).Find(&users)
	//fmt.Println(users)
	//// 第2页
	//DB.Limit(3).Offset(2).Find(&users)
	//fmt.Println(users)
	//// 第3页
	//DB.Limit(2).Offset(5).Find(&users)
	//fmt.Println(users)
	
	//var users Student
	//// 会过滤零值
	//DB.Where(&Student{Name: "李元芳", Age: 0}).Find(&users)
	//fmt.Println(users)
	
	//var users []Student
	//// 查询用户名是枫枫的
	//DB.Where("name = ?", "枫枫").Find(&users)
	//fmt.Println(users)
	//// 查询用户名不是枫枫的
	//DB.Where("name <> ?", "枫枫").Find(&users)
	//fmt.Println(users)
	//// 查询用户名包含 如燕，李元芳的
	//DB.Where("name in ?", []string{"如燕", "李元芳"}).Find(&users)
	//fmt.Println(users)
	//// 查询姓李的
	//DB.Where("name like ?", "李%").Find(&users)
	//fmt.Println(users)
	//// 查询年龄大于23，是qq邮箱的
	//DB.Where("age > ? and email like ?", "23", "%@qq.com").Find(&users)
	//fmt.Println(users)
	//// 查询是qq邮箱的，或者是女的
	//DB.Where("gender = ? or email like ?", false, "%@qq.com").Find(&users)
	//fmt.Println(users)
	
}