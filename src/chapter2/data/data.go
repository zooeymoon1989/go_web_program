package data

import (
	"database/sql"
	"log"
)
import _ "github.com/go-sql-driver/mysql"

var Db *sql.DB

func init()  {
	var err error
	Db ,err = sql.Open("mysql" , "root:2863186@/chitchat")

	if err != nil {
		log.Fatal(err)
	}

	return

}
