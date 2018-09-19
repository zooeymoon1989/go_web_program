package main

import (
	"html/template"
	"net/http"
)

func nest(w http.ResponseWriter, r *http.Request) {

	t , _ := template.ParseFiles("src/chapter5/templates/nest.html")
	t.ExecuteTemplate( w, "layout" , "")

}

func main() {

	server := http.Server{
		Addr:"127.0.0.1:9988",
	}

	http.HandleFunc("/nest" , nest)

	server.ListenAndServe()

}
