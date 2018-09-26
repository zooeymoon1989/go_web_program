package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


type PostSqlx struct {
	Id int
	Content string
	Author string
}

var DbSqlx *sqlx.DB


func init()  {
	var err error
	DbSqlx , err = sqlx.Open("mysql" , "root:2863186@/gwp")
	if err != nil {
		panic(err)
	}

}

func GetPostSqlx(id int) (post PostSqlx , err error) {
	err = DbSqlx.QueryRowx("SELECT id , content , author FROM posts WHERE id = ?" , id ).StructScan(&post)//.Scan(&post.Id , &post.Content , &post.AuthorName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer DbSqlx.Close()

	return

}

func (post PostSqlx)Create() (err error)  {
	result := DbSqlx.QueryRowx("INSERT INTO posts (content , author) VALUES ( ? , ? )" , post.Content , post.Author)
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println(result.Scan())
	return
}


func main() {

	post := PostSqlx{}

	post , _ = GetPostSqlx(1)
	fmt.Println(post)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(post)
	//post := PostSqlx{Content: "Hello World!", AuthorName: "Sau Sheong"}
	//post.Create()
	//fmt.Println(post)
}