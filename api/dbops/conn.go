package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err error
)

func init() {
	dbConn, err = sql.Open("mysql","root:FF558234ttpk@tcp(127.0.0.1:3306)/imooc181120?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}