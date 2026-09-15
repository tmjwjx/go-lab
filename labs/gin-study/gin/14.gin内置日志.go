package main

//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//)
//
//func LogGormatterParams(params gin.LogFormatterParams) string {
//	return fmt.Sprintf(
//		"[TOM] %s  |%d|  %s  %s  %s  %s\n",
//		params.TimeStamp.Format("2006-01-02 15:04:05"),
//		params.StatusCode,
//		params.MethodColor(), params.Method, params.ResetColor(),
//		params.Path,
//	)
//
//}
//func main() {
//	//file, _ := os.Create("gin.log")
//	//gin.DefaultWriter = io.MultiWriter(file)
//
//	//gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
//	//	log.Printf(
//	//		"[田家杰] %s %s %s %d \n",
//	//		httpMethod,
//	//		absolutePath,
//	//		handlerName,
//	//		nuHandlers,
//	//	)
//	//}
//
//	gin.SetMode(gin.ReleaseMode)
//	router := gin.New()
//
//	//router.Use(gin.LoggerWithFormatter(LogGormatterParams))
//	router.Use(gin.LoggerWithFormatter(LogGormatterParams))
//
//	router.GET("/index", func(c *gin.Context) {})
//	router.GET("/users", func(c *gin.Context) {})
//	router.GET("/articles", func(c *gin.Context) {})
//	router.GET("/articles/:id", func(c *gin.Context) {})
//
//	router.Run()
//}