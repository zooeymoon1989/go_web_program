package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func parseHtml(w http.ResponseWriter , r *http.Request)  {
	
	t , err := template.ParseFiles("src/chapter5/templates/tmpl.html"  , "src/chapter5/templates/t2.html")

	if err != nil{
		fmt.Println(err)
		return
	}

	t.Execute(w, "Hello World!")

}

func main() {

	server := http.Server{
		Addr:"127.0.0.1:9988",
	}

	http.HandleFunc("/parseHtml" , parseHtml)

	server.ListenAndServe()

}
