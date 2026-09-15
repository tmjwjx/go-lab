package main

//func Index(context *gin.Context) {
//	context.String(200, "Hello World!")
//
//}
//
//func main() {
//	//创建一个默认的路由
//	router := gin.Default()
//
//	// 绑定路由规则和路由函数,访问index的路由，将由对应的函数去处理
//	router.GET("/index", Index)
//
//	//启动监听，127.0.0.0.0
//	router.Run(":8080")
//
//	//用原生http服务的方式启动,router.Run本质就是http.ListenAndServe的进一步封装
//	//http.ListenAndServe(":8080", nil)
//
//}