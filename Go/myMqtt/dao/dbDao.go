package dao

import (
	"database/sql"
	"fmt"
)

func sql1() {

	db, err := sql.Open("mysql", "root:123456@/db_test?charset=utf8")
	checkErr(err)

	// 插入数据
	stmt, err := db.Prepare("INSERT user SET name=?,password=?")
	checkErr(err)
	res, err := stmt.Exec("test", "qwe123")
	fmt.Println("res == ", res)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	// 更新数据
	stmt, err = db.Prepare("update user set name=? where id=?")
	checkErr(err)
	res, err = stmt.Exec("test", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	// 查询数据
	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)
	for rows.Next() {
		var id int
		var name string
		var password string

		err = rows.Scan(&id, &name, &password)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(name)
		fmt.Println(password)
	}

	// 删除数据
	stmt, err = db.Prepare("delete from user where id=?")
	checkErr(err)
	//res, err = stmt.Exec(id)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	db.Close()
}
