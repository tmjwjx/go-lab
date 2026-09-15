package main

import (
	"database/sql"
	"fmt"
	_ "mysql-master"
)

// 修改表中的数据
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
	stmt, _ := db.Prepare("update people set name=?,address=? where id=?")
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	//获取结果
	r, _ := stmt.Exec("李四", "朝阳", 3)
	count, _ := r.RowsAffected()
	if count > 0 {
		fmt.Println("修改成功")
	} else {
		fmt.Println("修改失败")

	}
}
