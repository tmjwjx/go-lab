package main

//type ArticleModel struct {
//	Title   string `json:"title"`
//	Content string `json:"content"`
//}
//type Response struct {
//	Code int    `json:"code"`
//	Date any    `json:"date"`
//	Msg  string `json:"msg"`
//}
//
//func _bindJson(c *gin.Context, obj any) (err error) {
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
//func _getList(c *gin.Context) {
//
//	articleList := []ArticleModel{
//		{"Go语言title", "Go语言content"},
//		{"Python语言title", "Python语言content"},
//		{"Java语言title", "Java语言content"},
//	}
//
//	c.JSON(200, Response{0, articleList, "成功"})
//}
//func _getDetail(c *gin.Context) {
//	//获取param中的id
//	fmt.Println(c.Param("id"))
//	article := ArticleModel{"Go语言title", "Go语言content"}
//
//	c.JSON(200, Response{0, article, "成功"})
//}
//func _create(c *gin.Context) {
//	//接收前端传来的json
//	var article ArticleModel
//	err := _bindJson(c, &article)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	c.JSON(200, Response{0, article, "添加成功"})
//}
//func _update(c *gin.Context) {
//	fmt.Println(c.Param("id"))
//	var article ArticleModel
//	err := _bindJson(c, &article)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	c.JSON(200, Response{0, article, "添加成功"})
//}
//func _delete(c *gin.Context) {
//	fmt.Println(c.Param("id"))
//	c.JSON(200, Response{0, map[string]string{}, "删除成功"})
//}
//
//func main() {
//	router := gin.Default()
//
//	router.GET("/articles", _getList)       //文章列表
//	router.GET("/articles/:id", _getDetail) //文章详情
//	router.POST("/articles", _create)       //添加文章
//	router.PUT("/articles/:id", _update)    //编辑文章
//	router.DELETE("/articles/:id", _delete) //删除文章
//
//	router.Run(":80")
//}