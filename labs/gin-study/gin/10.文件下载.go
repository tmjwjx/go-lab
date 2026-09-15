package main

//func main() {
//	router := gin.Default()
//	router.GET("/download", func(c *gin.Context) {
//		c.Header("Content-Type", "application/octet-stream")              // 表示是文件流，唤起浏览器下载，一般设置了这个，就要设置文件名
//		c.Header("Content-Disposition", "attachment; filename="+"牛逼.png") // 用来指定下载下来的文件名
//		c.Header("Content-Transfer-Encoding", "binary")                   // 表示传输过程中的编码形式，乱码问题可能就是因为它
//		c.File("uploads/12.png")
//
//	})
//	router.Run(":8080")
//
//}