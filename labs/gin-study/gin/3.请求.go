package main

//func _query(c *gin.Context) {
//	fmt.Println(c.Query("user"))
//	fmt.Println(c.GetQuery("user"))
//	fmt.Println(c.QueryArray("user"))
//}
//func _param(c *gin.Context) {
//	fmt.Println(c.Param("user_id"))
//	fmt.Println(c.Param("book_id"))
//	c.String(200, c.Param("user_id"))
//	c.String(200, c.Param("book_id"))
//}
//func _form(c *gin.Context) {
//	fmt.Println(c.PostForm("name"))
//	fmt.Println(c.PostFormArray("name"))
//	//如果用户没有传，使用默认值
//	fmt.Println(c.DefaultPostForm("addr", "河南省"))
//	//接收所有的参数，包括文件
//	forms, err := c.MultipartForm()
//	fmt.Println(forms, err)
//}
//func bindJson(c *gin.Context, obj any) (err error) {
//	body, _ := c.GetRawData()
//	contentType := c.GetHeader("Content-Type")
//	fmt.Println(contentType)
//	switch contentType {
//	case "application/json":
//		err := json.Unmarshal(body, &obj)
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//	}
//	return nil
//}
//func _raw(c *gin.Context) { //获取原始参数
//
//	type User struct {
//		UserName string `json:"username"`
//		Password string `json:"password"`
//	}
//	var user User
//	err := bindJson(c, &user)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(user)
//
//	//body, _ := c.GetRawData()
//	//contentType := c.GetHeader("Content-Type")
//	//fmt.Println(contentType)
//	//switch contentType {
//	//case "application/json":
//	//	type User struct {
//	//		UserName string `json:"username"`
//	//		Password string `json:"password"`
//	//	}
//	//	var user User
//	//	err := json.Unmarshal(body, &user)
//	//	if err != nil {
//	//		fmt.Println(err.Error())
//	//	}
//	//	fmt.Println(user)
//	//}
//}
//
//func main() {
//	router := gin.Default()
//	router.GET("/query", _query)
//	router.GET("/param/:user_id", _param)
//	router.GET("/param/:user_id/:book_id", _param)
//	router.POST("/form", _form)
//	router.POST("/raw", _raw)
//
//	router.Run(":8080")
//
//}