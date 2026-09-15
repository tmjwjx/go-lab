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
	stmt, _ := db.Prepare("delete from people where id=?")
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	//获取结果
	r, _ := stmt.Exec(2)
	count, _ := r.RowsAffected()
	if count > 0 {
		fmt.Println("删除成功")
	} else {
		fmt.Println("删除失败")
	}
}
