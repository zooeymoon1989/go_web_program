package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type PostGWP struct {
	Id     int
	Content string
	Author  string
}


var Db *sql.DB

func init()  {
	var err error
	Db , err = sql.Open("mysql" , "root:2863186@/gwp")
	if err != nil {
		panic(err)
	}

}

func (post *PostGWP) Create()  (err error){
	statement := "insert into posts (content , author) values (? , ?)"
	stmt , err = Db.Prepare(statement)
	if err != nil {
		return
	}

	defer  stmt.Close()

	err = stmt.QueryRow(post.Content , post.Author).Scan(&post.Id)
	return
}

func main() {

	post := PostGWP{
		Content : "hello world!",
		Author : "liwenqiang",
	}

	fmt.Println(post)
	post.Create()
	fmt.Println(post)
}
