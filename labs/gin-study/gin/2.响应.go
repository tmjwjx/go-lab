package main

//// 最重要的 json
//func _json(c *gin.Context) {
//	//json响应结构体 最多的最常用
//	type UserInfo struct {
//		UserName string `json:"user_name"`
//		Age      int    `json:"age"`
//		Password string `json:"-"` //忽略转化为json
//	}
//	user := UserInfo{"tjj", 18, "123123"}
//	c.JSON(200, user)
//
//	////json响应map
//	//userMap := map[string]string{
//	//	"user_name": "tjj",
//	//	"age":       "18",
//	//}
//	//c.JSON(200, userMap)
//
//	////直接响应json
//	//c.JSON(200, gin.H{"username": "tjj", "age": "18"})
//}
//func _string(c *gin.Context) {
//	c.String(200, "你好啊")
//}
//func _xml(c *gin.Context) {
//	c.XML(200, gin.H{"user": "tjj", "message": "hello", "status": http.StatusOK})
//}
//func _yaml(c *gin.Context) {
//	c.YAML(200, gin.H{"user": "tjj", "message": "hello", "status": http.StatusOK})
//}
//func _html(c *gin.Context) {
//	type UserInfo struct {
//		UserName string `json:"username"`
//		Age      int    `json:"age"`
//		Password string `json:"-"` //忽略转化为json
//	}
//	user := UserInfo{"tjj", 18, "123123"}
//	c.HTML(200, "index.html", user)
//}
//func _redirect(c *gin.Context) {
//	//永久重定向
//	//c.Redirect(301, "https://www.baidu.com")
//	//临时重定向
//	c.Redirect(302, "https://cn.bing.com")
//}
//
//func main() {
//	router := gin.Default()
//
//	//加载模板目录下所有的模板文件
//	router.LoadHTMLGlob("templates/*")
//	//配置网页请求静态目录前缀
//	router.StaticFS("/static", http.Dir("static/static"))
//	//配置单个文件 网络请求的路由 文件的路径
//	router.StaticFile("/girl.png", "static/girl.png")
//
//	router.GET("/string", _string)
//	router.GET("/json", _json)
//	router.GET("/xml", _xml)
//	router.GET("/yaml", _yaml)
//	router.GET("/html", _html)
//	router.GET("/redirect", _redirect)
//	router.Run(":8080")
//}