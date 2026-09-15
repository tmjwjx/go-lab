package main

import (
	"database/sql"
	"fmt"
	_ "mysql-master"
)

func main() {
	//打开连接
	db, _ := sql.Open("mysql", "root:9437498Tjj@tcp(localhost:3306)/first")
	//db.Ping()
	defer func() {
		if db != nil {
			db.Close()
		}
	}()
	//预处理sql
	stmt, err := db.Prepare("select * from people")
	if err != nil {
		fmt.Println("结果获取失败")
		return
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("获取结果失败")
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	for rows.Next() {
		var id int
		var name string
		var address string
		rows.Scan(&id, &name, &address)
		fmt.Println(id, name, address)
	}
}
