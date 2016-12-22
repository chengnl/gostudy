package dao

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Tdb *sql.DB

func init() {
	//初始化
	log.Println("初始化数据库连接开始。。。。。")
	Tdb, _ = sql.Open("mysql", "root:vrvim@tcp(192.168.0.59:3306)/IM_TIMESTAMP?charset=utf8")
	Tdb.SetMaxOpenConns(10)
	Tdb.SetMaxIdleConns(5)
	err := Tdb.Ping()
	if err != nil {
		log.Printf("初始化数据库连接失败：%v\n", err)
		panic(err)
	}
	log.Println("初始化数据库连接成功。")
}
