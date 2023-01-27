package service

import (
	"database/sql"
	"fmt"
	_"github.com/go-sql-driver/mysql"
)

var Db *sql.DB

var err error

func init() {

	fmt.Println("This will get called on main initialization")
	fmt.Println("starts")
	Db, err = sql.Open("mysql", "root:#Anjali123@tcp(127.0.0.1:3306)/balance")
	if err != nil {
		fmt.Println("error", err)
	}
	fmt.Println("dababase connected")

	//defer db.Close()

}
