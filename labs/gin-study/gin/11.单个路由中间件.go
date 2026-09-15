package main

//func m1(c *gin.Context) {
//	fmt.Println("m1...in")
//	//c.Next()
//	c.Abort()
//	fmt.Println("m1...out")
//	//c.Abort() //终止，后面的就没了
//}
//func m2(c *gin.Context) {
//	fmt.Println("m2...in")
//	//c.Next()
//	//c.Abort()
//	fmt.Println("m2...out")
//}
//func m3(c *gin.Context) {
//	fmt.Println("m3...in")
//	//c.Next()
//	fmt.Println("m3...out")
//}
//
//func main() {
//	router := gin.Default()
//
//	router.GET("/", m1, m2, m3)
//
//	router.Run(":8080")
//
//}