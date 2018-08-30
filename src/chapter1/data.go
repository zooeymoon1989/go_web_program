package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"
)

import _ "github.com/go-sql-driver/mysql"

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}


func handler(writer http.ResponseWriter, request *http.Request) {

	var post = Post{}

	db , err := sql.Open("mysql" , "root:@/chitchat")

	if err != nil {
		fmt.Fprintln(writer , err)
	}

	rows , err := db.Query("SELECT * FROM posts " )

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {

		if err = rows.Scan(&post.Id,&post.Uuid,&post.Body,&post.UserId,&post.ThreadId,&post.CreatedAt) ; err != nil {
			fmt.Println(err)
		}

	}

	rows.Close()

	fmt.Fprintf(writer, "Hello World, %s!", post.Body)
	//fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8989", nil)
}