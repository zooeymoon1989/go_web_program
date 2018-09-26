package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type PostGWP struct {
	Id     int
	Content string
	Author  string
	Comment []CommentGWP
}

type CommentGWP struct {
	Id int
	Content string
	Author string
	Post *PostGWP
}


var Db *sql.DB

func init()  {
	var err error
	Db , err = sql.Open("mysql" , "root:@/gwp")
	if err != nil {
		panic(err)
	}

}

func (post *PostGWP) Create()  (err error){
	statement := "insert into posts (content , author) values (? , ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer  stmt.Close()

	result , err := stmt.Exec(post.Content, post.Author)
	if err != nil {
		fmt.Println(err)
		return
	}

	id , _ := result.LastInsertId()

	post.Id = int(id)

	return
}

func (comment *CommentGWP) CreateComment() (err error)  {
	if comment.Post == nil {
		err = errors.New("Post not found")
		return
	}

	result , err  := Db.Exec("INSERT INTO comments (content , author , post_id ) VALUES ( ? , ? , ?)" , comment.Content , comment.Author , comment.Post.Id)
	if err != nil{
		fmt.Println(err)
		return
	}

	id , _ := result.LastInsertId()

	comment.Id = int(id)

	defer  Db.Close()
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
		return
	}

	rows , err := Db.Query("SELECT id , content , author FROM comments WHERE post_id = ?" , id)

	if err != nil {
		fmt.Println(err)
		return
	}

	for rows.Next() {
		comment := CommentGWP{Post:&post}
		err = rows.Scan(&comment.Id , &comment.Content , &comment.Author)
		if err != nil {
			fmt.Println(err)
			return
		}

		post.Comment = append(post.Comment , comment)
	}

	rows.Close()
	return


	Db.Close()
	return
}

func main() {

	//post := PostGWP{
	//	Content : "hello world!",
	//	Author : "liwenqiang",
	//}
	//
	//post.Create()
	//
	//comment := CommentGWP{Content: "Good post!", Author: "Joe", Post: &post}
	//fmt.Println(*comment.Post)
	//comment.CreateComment()

	//posts, _ := Posts(2)
	//fmt.Println(posts)

	getPost , _ := GetPost(22)

	fmt.Println(*getPost.Comment[0].Post)

}
