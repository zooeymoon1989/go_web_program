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
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}

	defer  stmt.Close()

	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)

	return
}

func Posts(limit int) (posts []PostGWP , err error) {
	rows , err := Db.Query("SELECT id , content , author  FROM posts LIMIT ?" , limit)
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next(){
		post := PostGWP{}
		err = rows.Scan(&post.Id , &post.Content , &post.Author)

		if err != nil {
			return
		}

		posts = append(posts, post)
	}

	rows.Close()

	return
}

func GetPost(id int) (post PostGWP , err error) {
	post = PostGWP{}
	err = Db.QueryRow("SELECT id , content , author FROM posts where id = ?" , id).Scan(&post.Id , &post.Content , &post.Author)
	if err != nil {
		fmt.Println(err)
	}
	Db.Close()
	return
}

func main() {

	//post := PostGWP{
	//	Content : "hello world!",
	//	Author : "liwenqiang",
	//}
	//
	//fmt.Println(post)
	//post.Create()
	//fmt.Println(post)

	//posts, _ := Posts(2)
	//fmt.Println(posts)

	getPost , _ := GetPost(2)
	fmt.Println(getPost)
}
