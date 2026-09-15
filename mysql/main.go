package main

import (
	"database/sql"
	"fmt"
	// 空导入 mysql驱动
	_ "github.com/go-sql-driver/mysql"
)

func insert() {
	// 1.打开连接
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/first")
	db.Ping()
	defer func() {
		if db != nil {
			db.Close()
		}
	}()
	if err != nil {
		fmt.Println("数据连接失败！ err:", err)
		return
	}

	// 2.预处理SQL
	// ?表示占位符
	stmt, err := db.Prepare("insert into people values(default, ?, ?)")
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	if err != nil {
		fmt.Println("预处理失败！ err:", err)
		return
	}

	// 参数和占位符对应
	r, err := stmt.Exec("张三", "海淀")
	if err != nil {
		fmt.Println("sql执行失败！ err:", err)
		return
	}

	// 3.获取结果
	count, err := r.RowsAffected()
	if err != nil {
		fmt.Println("获取结果失败！ err:", err)
		return
	}
	if count > 0 {
		fmt.Println("新增成功！")
	} else {
		fmt.Println("新增失败！")
	}

	// 4.关闭预处理

	// 获取到新增组件的值
	lastId, err := r.LastInsertId()
	if err != nil {
		fmt.Println("获取新增ID失败！ err:", err)
		return
	}
	fmt.Println("新增ID：", lastId)
}

func update() {
	db, _ := sql.Open("mysql", "root:root@tcp(localhost:3306)/first")
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	stmt, _ := db.Prepare("update people set name = ?, address = ? where id = ?")
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	r, _ := stmt.Exec("李四", "朝阳", 1)
	count, _ := r.RowsAffected()
	if count > 0 {
		fmt.Println("更新成功！")
	} else {
		fmt.Println("更新失败！")
	}
}

func del() {
	db, _ := sql.Open("mysql", "root:root@tcp(localhost:3306)/first")
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	stmt, _ := db.Prepare("delete from people where id = ?")
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	r, _ := stmt.Exec(1)
	count, _ := r.RowsAffected()
	if count > 0 {
		fmt.Println("删除成功！")
	} else {
		fmt.Println("删除失败！")
	}

}

func query() {
	db, _ := sql.Open("mysql", "root:root@tcp(localhost:3306)/first")
	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	stmt, _ := db.Prepare("select * from people")
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	rows, _ := stmt.Query()
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

func main() {
	//insert()
	//update()
	//del()
	query()
}
