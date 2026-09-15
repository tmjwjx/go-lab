package main

//type UserInfo struct {
//	Name string `json:"name" form:"name"`
//	Age  int    `json:"age" form:"age"`
//	Sex  string `json:"sex" form:"sex"`
//}
//
//func main() {
//	router := gin.Default()
//
//	router.POST("/query", func(c *gin.Context) {
//		fmt.Println(c.GetHeader("Content-Type"))
//
//		var userInfo UserInfo
//		err := c.ShouldBindQuery(&userInfo)
//		if err != nil {
//			c.JSON(200, gin.H{"msg": "你错了"})
//			return
//		}
//		c.JSON(200, userInfo)
//	})
//
//	router.Run(":8080")
//}