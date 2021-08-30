package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var dbcon = ""

func DBInit(debugmode bool) { //初始化数据库
	if debugmode == true {
		dbcon = "root:root@tcp(127.0.0.1:3306)/?charset=utf8"
	} else {
		dbcon = "root:yTtqVWWNuS6qzbfL@tcp(127.0.0.1:3306)/?charset=utf8"
	}
	_db, err := sql.Open("mysql", dbcon)
	DB = _db
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}