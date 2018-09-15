package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func processRandom(w http.ResponseWriter , r *http.Request)  {

	t , err := template.ParseFiles("src/chapter5/templates/tmpl.html","src/chapter5/templates/t2.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("I am calling")
	t.Execute(w, "hello")

}

func main() {
	server := http.Server{
		Addr:"127.0.0.1:9977",
	}

	http.HandleFunc("/random" , processRandom)
	server.ListenAndServe()
}
