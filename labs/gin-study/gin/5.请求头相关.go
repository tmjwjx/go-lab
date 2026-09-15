package main

//func hello(c *gin.Context) {
//	//首字母大小写不区分 单词之间用 - 连接
//	//如果时使用Get方法或是GetHeader，可以不用区分大小写，并且返回第一个value
//	//如果是用map的取值方式，请注意大小写问题
//	fmt.Println(c.GetHeader("User-Agent"))
//	fmt.Println(c.Request.Header)
//	fmt.Println(c.Request.Header.Get("User-Agent"))
//	fmt.Println(c.Request.Header["User-Agent"])
//
//	c.JSON(200, gin.H{"msg": "成功"})
//}
//func identify(c *gin.Context) {
//	userAgent := c.GetHeader("User-Agent")
//	// 用正则去匹配
//	// 字符串的包含匹配
//	if strings.Contains(userAgent, "python") {
//		c.JSON(0, gin.H{"data": "这是响应给爬虫的数据"})
//		return
//	}
//	c.JSON(0, gin.H{"data": "这是响应给用户的数据"})
//}
//func response(c *gin.Context) {
//	c.Header("Token", "What is Token?")
//	c.Header("Content-Type", "application/text; charset=utf-8")
//	c.JSON(0, gin.H{"data": "看看响应头"})
//}
//
//func main() {
//	router := gin.Default()
//
//	router.GET("/", hello)
//	//爬虫和用户区别对待
//	router.GET("/python", identify)
//	//设置响应头
//	router.GET("/response", response)
//
//	router.Run(":8080")
//}