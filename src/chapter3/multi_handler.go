package main

import (
	"fmt"
	"net/http"
)

type BasicHandler struct{}

func (b *BasicHandler) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w , "I created this function!")
}

func hello(w http.ResponseWriter , r *http.Request)  {
	fmt.Fprintf(w , "Hello!")
}

func world(w http.ResponseWriter , r *http.Request)  {
	fmt.Fprintf(w , "World!")
}



func main() {

	basic := BasicHandler{}

	server := http.Server{
		Addr:	"127.0.0.1:7776",
	}

	http.Handle("/haha" , &basic)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world" , world)
	http.HandleFunc("/test" , func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer , "test")
	})

	server.ListenAndServe()
}