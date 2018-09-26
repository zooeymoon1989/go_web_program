package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


type PostSqlx struct {
	Id int
	Content string
	AuthorName string `db: author`
}

var DbSqlx *sqlx.DB
var Row *sqlx.Row


func init()  {
	var err error
	DbSqlx , err = sqlx.Open("mysql" , "root:@/gwp")
	if err != nil {
		panic(err)
	}

}

func GetPostSqlx(id int) (post PostSqlx , err error) {
	Row = DbSqlx.QueryRowx("SELECT id , content , author FROM posts WHERE id = ?" , id )
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	aa , _ := Row.SliceScan()


	i := 0

	for i < len(aa) {

		fmt.Println(aa[i])
		i ++

	}


	if err != nil {
		fmt.Println(err)
		return
	}


	defer DbSqlx.Close()

	return
	//fmt.Println(rows.Err())
	//err = DbSqlx.Get(&post , "SELECT id , content , author FROM posts WHERE id = ?" , post.Id)

}

func (post PostSqlx)Create() (err error)  {
	result := DbSqlx.QueryRowx("INSERT INTO posts (content , author) VALUES ( ? , ? )" , post.Content , post.AuthorName)
	if err != nil {

		fmt.Println(err)
	}
	fmt.Println(result.Scan())
	return
}


func main() {

	GetPostSqlx(22)

	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(post)
	//post := PostSqlx{Content: "Hello World!", AuthorName: "Sau Sheong"}
	//post.Create()
	//fmt.Println(post)
}