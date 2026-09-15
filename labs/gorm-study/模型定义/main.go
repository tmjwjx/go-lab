package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var mysqlLogger logger.Interface

func init() {
	username := "root"       //账号
	password := "9437498Tjj" //密码
	host := "127.0.0.1"      //数据库地址，可以是Ip或者域名
	port := 3306             //数据库端口
	Dbname := "gorm"         //数据库名
	timeout := "10s"         //连接超时，10秒
	
	mysqlLogger = logger.Default.LogMode(logger.Info)
	
	// root:root@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//跳过默认事务，能获得 60% 的性能提升
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
	DB = db
	//fmt.Println(db)
	fmt.Println("连接成功")
}

type Student struct {
	ID   uint
	Name string
	Age  int
}

func main() {
	//fmt.Println(DB)
	
	//显示日志
	DB = DB.Session(&gorm.Session{
		Logger: mysqlLogger,
	})
	
	//Debug显示日志 AutoMigrate自动生成表结构
	//DB.Debug().AutoMigrate(&Student{})
	DB.AutoMigrate(&Student{}) //只新增，不删除，不修改（大小会修改）
	//常识：小写属性是不会生成字段的
}