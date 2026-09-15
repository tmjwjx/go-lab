package main

//type UserInfo struct {
//	Name string `json:"name"`
//	Age  int    `json:"age"`
//}
//type Response struct {
//	Code int    `json:"code"`
//	Data any    `json:"data"`
//	Msg  string `json:"msg"`
//}
//
//func UserListView(c *gin.Context) {
//	var userList []UserInfo = []UserInfo{{"tom", 18}, {"李白", 100}}
//	c.JSON(200, Response{0, userList, "请求成功"})
//}
//
//func main() {
//	// gin.Default()默认使用了Logger和Recovery中间件
//	router := gin.Default()
//
//	api := router.Group("api")
//	{ //一组一般放到一个大括号内 看起来直观一点
//		api.GET("/user", UserListView)
//	}
//
//	router.Run()
//}