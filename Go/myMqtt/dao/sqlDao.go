package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
)

var sourceName string

//初始化数据库
func Init(name string, password string, dbName string) {
	//root:123456@/db_test?charset=utf8
	sourceName = name + ":" + password + "@/" + dbName + "charset=utf8"
	fmt.Println("数据库连接信息是： ", sourceName)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/**
增加数据
 */
func Insert() {

	db, err := sql.Open("mysql", sourceName)
	checkErr(err)

	stmt, err := db.Prepare("INSERT user SET name=?,password=?")
	checkErr(err)
	res, err := stmt.Exec("test", "qwe123")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	db.Close()
}

/**
删除数据
 */
func Delete(id int) {
	db, err := sql.Open("mysql", "root:123456@/db_test?charset=utf8")
	checkErr(err)

	// 删除数据
	stmt, err := db.Prepare("delete from user where id=?")
	checkErr(err)
	res, err := stmt.Exec(id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	db.Close()

	db.Close()
}

/**
更新数据
 */
func Update(id int) {
	db, err := sql.Open("mysql", "root:123456@/db_test?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare("update user set name=? where id=?")
	checkErr(err)
	res, err := stmt.Exec("test", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	db.Close()
}

/**
查询数据
 */
func Select() {
	db, err := sql.Open("mysql", "root:123456@/db_test?charset=utf8")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)
	fmt.Println("查询的数据： ",rows)
	for rows.Next() {
		var id int
		var name string
		var password string

		err = rows.Scan(&id, &name, &password)
		checkErr(err)
		fmt.Print(id, " ", name, " ", password, "\n")
	}

	db.Close()
}
