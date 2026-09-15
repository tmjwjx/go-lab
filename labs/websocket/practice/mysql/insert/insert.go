package main

import (
	"database/sql"
	"fmt"
	_ "mysql-master"
)

// 向数据库的表中添加数据
func main() {
	//打开连接
	db, err := sql.Open("mysql", "root:9437498Tjj@tcp(localhost:3306)/first")
	db.Ping()
	defer db.Close()
	if err != nil {
		fmt.Println("数据库连接失败")
		return
	}

	//预处理sql
	stmt, err := db.Prepare("insert into people values (default,?,?)")
	defer stmt.Close()
	if err != nil {
		fmt.Println("预处理失败！", err)
		return
	}

	r, err := stmt.Exec("张三", "海淀")
	if err != nil {
		fmt.Println("预处理失败！")
		return
	}

	//获取结果
	count, err := r.RowsAffected()
	if err != nil {
		fmt.Println("结果获取失败")
		return
	}
	if count > 0 {
		fmt.Println("新增成功")
	} else {
		fmt.Println("新增失败")
	}

	//可能需要获取到新增时主键的值
	id, _ := r.LastInsertId()
	fmt.Println(id)

}
