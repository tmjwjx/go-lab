package main

import (
	gorm_study "gorm-study"
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
	
	//email := "xxx@qq.com"
	//// 创建记录
	//student := Student{
	//	Name:   "tom",
	//	Age:    21,
	//	Gender: true,
	//	Email:  &email,
	//}
	//DB.Create(&student)
	//fmt.Println("添加成功")
	
	//var test Student
	//DB.Take(&test)
	//fmt.Println(test)
	//// SELECT * FROM `students` LIMIT 1
	//test = Student{}
	//DB.First(&test)
	//fmt.Println(test)
	//// SELECT * FROM `students` ORDER BY `students`.`id` LIMIT 1
	//test = Student{}
	//DB.Last(&test)
	//fmt.Println(test)
	//// SELECT * FROM `students` ORDER BY `students`.`id` DESC LIMIT 1
	
	//var student Student
	//DB.Take(&student, 2)
	//fmt.Println(student)
	//
	//student = Student{} // 重新赋值
	//DB.Take(&student, "4")
	//fmt.Println(student)
	
	//studentList := []Student{}
	//count := DB.Find(&studentList).RowsAffected
	//fmt.Println(count)
	
	//var studentList []Student
	//DB.Find(&studentList)
	//for _, student := range studentList {
	//	fmt.Println(student)
	//}
	
	// 由于email是指针类型，所以看不到实际的内容
	// 但是序列化之后，会转换为我们可以看得懂的方式
	//var studentList []Student
	//DB.Find(&studentList)
	//for _, student := range studentList {
	//
	//	data, _ := json.Marshal(student)
	//	fmt.Println(string(data))
	//}
	
	//var student Student
	////student.ID = 4
	//student.Age = 50
	//DB.Select("age").Save(&student)
	
	var studentList []Student
	DB.Find(&studentList).Delete(&studentList)
	studentList = []Student{
		{ID: 1, Name: "李元芳", Age: 32, Email: PtrString("lyf@yf.com"), Gender: true},
		{ID: 2, Name: "张武", Age: 18, Email: PtrString("zhangwu@lly.cn"), Gender: true},
		{ID: 3, Name: "枫枫", Age: 23, Email: PtrString("ff@yahoo.com"), Gender: true},
		{ID: 4, Name: "刘大", Age: 54, Email: PtrString("liuda@qq.com"), Gender: true},
		{ID: 5, Name: "李武", Age: 23, Email: PtrString("liwu@lly.cn"), Gender: true},
		{ID: 6, Name: "李琦", Age: 14, Email: PtrString("liqi@lly.cn"), Gender: false},
		{ID: 7, Name: "晓梅", Age: 25, Email: PtrString("xiaomeo@sl.com"), Gender: false},
		{ID: 8, Name: "如燕", Age: 26, Email: PtrString("ruyan@yf.com"), Gender: false},
		{ID: 9, Name: "魔灵", Age: 21, Email: PtrString("moling@sl.com"), Gender: true},
	}
	DB.Create(&studentList)
	
}
func PtrString(email string) *string {
	return &email
}