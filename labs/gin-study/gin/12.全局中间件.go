package main

//type User struct {
//	Name string `json:"name"`
//	Age  int    `json:"age"`
//}
//
//func m1(c *gin.Context) {
//	fmt.Println("m1...in")
//	//name, _ := c.Get("name")
//	user, _ := c.Get("user")
//	c.JSON(200, user.(User).Name)
//	fmt.Println("m1...out")
//}
//func m2(c *gin.Context) {
//	fmt.Println("m2...in")
//	fmt.Println("m2...out")
//}
//func m3(c *gin.Context) {
//	fmt.Println("m3...in")
//	c.Set("name", "tom")
//	c.Set("user", User{
//		Name: "tom",
//		Age:  18,
//	})
//	fmt.Println("m3...out")
//}
//
//func main() {
//	router := gin.Default()
//
//	router.Use(m3)
//
//	router.GET("/m1", m1)
//	router.GET("/m2", m2)
//	router.GET("/m3", m3)
//
//	router.Run(":8080")
//}