package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter , r *http.Request)  {
	fmt.Fprintf(w , "Hello!")
}

func world(w http.ResponseWriter , r *http.Request)  {
	fmt.Fprintf(w , "World!")
}

func main() {

	server := http.Server{
		Addr:	"127.0.0.1:7776",
	}

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world" , world)
	http.HandleFunc("/test" , func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer , "test")
	})

	server.ListenAndServe()
}