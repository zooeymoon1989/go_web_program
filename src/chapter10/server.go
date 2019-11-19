package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"path"
	"strconv"
)

var Db *sql.DB


type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func init()  {
	var err error
	Db , err = sql.Open("mysql" , "root:@/gwp")
	if err != nil {
		panic(err)
	}

}


func main() {
	server := http.Server{
		Addr: "127.0.0.1:9898",
	}

	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	var post = new(Post)
	id, err := getUrlId(r)
	if err != nil {
		return
	}

	err = post.Retrieve(id)
	if err != nil {
		return
	}
	err = post.Delete()
	if err != nil {
		return
	}

	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	var post = new(Post)
	id, err := getUrlId(r)

	if err != nil {
		return
	}

	err = post.Retrieve(id)
	if err != nil {
		return
	}

	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &post)
	err = post.Update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var post Post
	json.Unmarshal(body, &post)
	err = post.Create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	var post = new(Post)
	id, err := getUrlId(r)
	if err != nil {
		return
	}

	err = post.Retrieve(id)
	if err != nil {
		return
	}

	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func getUrlId(r *http.Request) (id int, err error) {
	return strconv.Atoi(path.Base(r.URL.Path))
}

func (post *Post)Retrieve(id int) ( err error) {
	err = Db.QueryRow("select id, content, author from posts where id = ?", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post)Create() (err error) {
	statement := "INSERT INTO posts (content , author) values (? , ?)"
	stmt , err :=Db.Prepare(statement)
	if err != nil {
		return
	}
	defer func() {
		err = stmt.Close()
		if err != nil{
			return
		}
	}()

	result , err := stmt.Exec(post)
	if err != nil {
		fmt.Println(err)
		return
	}

	id , _ := result.LastInsertId()

	post.Id = int(id)
	return

}

func (post *Post)Update() (err error) {
	_ , err = Db.Exec("UPDATE posts SET content = ? , author = ? where id = ? ", &post.Content , &post.Author , &post.Id)
	return
}

func (post *Post)Delete() (err error) {
	_ , err = Db.Exec("DELETE FROM  posts WHERE id = ?" , &post.Id)
	return
}
