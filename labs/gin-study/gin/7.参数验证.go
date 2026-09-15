package main

//type SignUserInfo struct {
//	Name       string `json:"name"`                                //用户名
//	Age        int    `json:"age"`                                 //年龄
//	Sex        string `json:"sex" binding:"required" msg:"性别校验失败"` //性别
//	Password   string `json:"password"`                            //密码
//	RePassword string `json:"re_password"`                         //确认密码
//}
//
//// GetValidMsg 获取结构体中的msg参数
//func GetValidMsg(err error, obj any) string {
//	getObj := reflect.TypeOf(obj)
//	//把err接口断言为具体类型
//	if errs, ok := err.(validator.ValidationErrors); ok {
//		//断言成功
//		for _, e := range errs {
//			//循环每一个错误信息
//			//根据报错字段获取结构体的具体字段
//			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
//				msg := f.Tag.Get("msg")
//				return msg
//			}
//		}
//	}
//	//return err.Error()
//	return ""
//}
//
//func main() {
//	router := gin.Default()
//
//	router.POST("/", func(c *gin.Context) {
//
//		var user SignUserInfo
//		err := c.ShouldBindJSON(&user)
//		if err != nil {
//
//			c.JSON(200, gin.H{"msg": GetValidMsg(err, &user)})
//			return
//		}
//
//		c.JSON(200, gin.H{"msg": "成功"})
//		return
//	})
//
//	_ = router.Run(":8080")
//
//}