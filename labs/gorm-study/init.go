package gorm_study

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Initialization() *gorm.DB {
	username := "root"       //账号
	password := "9437498Tjj" //密码
	host := "127.0.0.1"      //数据库地址，可以是Ip或者域名
	port := 3306             //数据库端口
	Dbname := "gorm"         //数据库名
	timeout := "10s"         //连接超时，10秒
	
	// root:root@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//跳过默认事务，能获得 30-60% 的性能提升
		SkipDefaultTransaction: true,
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix:   "f_", // 表名前缀
		//	SingularTable: true, // 单数表名 开关
		//	NoLowerCase:   true, // 关闭小写转换 开关
		//},
		
		////显示日志
		//Logger: mysqlLogger,
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 连接成功
	fmt.Println("连接成功")
	return db
}