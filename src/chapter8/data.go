package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init()  {
	var err error
	Db , err = sql.Open("mysql" , "root:2863186@/gwp")
		if err != nil {
			panic(err)
	}

}


func Retrieve(id int) (post Post, err error) {
	err = Db.QueryRow("select id, content, author from posts where id = ?", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}