package dbops

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	fmt.Println("init")
	dbConn, err = sql.Open("mysql", "root:seecha888@tcp(192.168.2.174)/video_server?charset=utf8")
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	fmt.Println("success")
}
