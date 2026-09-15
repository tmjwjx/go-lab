package main

//func main() {
//	router := gin.Default()
//
//	router.POST("/upload", func(c *gin.Context) {
//		//form-data 文件流 上传文件
//		file, _ := c.FormFile("file")
//		readerFile, _ := file.Open()
//		writerFile, _ := os.Create("./uploads/13.png") //要确保有 uploads 这个文件夹
//		defer writerFile.Close()
//		n, _ := io.Copy(writerFile, readerFile)
//		fmt.Println(n)
//
//		fmt.Println(file.Filename)
//		fmt.Println(file.Size / 1024)
//		c.JSON(200, gin.H{"msg": "上传成功"})
//	})
//	router.POST("/uploads", func(c *gin.Context) {
//		form, _ := c.MultipartForm()
//		files := form.File["upload[]"]
//		for _, file := range files {
//			c.SaveUploadedFile(file, "./uploads/"+file.Filename)
//
//		}
//		c.JSON(200, gin.H{"msg": fmt.Sprintf("成功上传 %d 个文件", len(files))})
//	})
//
//	router.Run(":8080")
//}