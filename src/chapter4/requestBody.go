package main

import (
	"fmt"
	"net/http"
)

func body(w http.ResponseWriter , r *http.Request)  {
	length := r.ContentLength
	body :=make([]byte , length)
	r.Body.Read(body)
	fmt.Println("I am calling !")
	fmt.Println(length)
	fmt.Fprintln(w , string(body))

}

func main() {

	server := http.Server{
		Addr:"127.0.0.1:7766",
	}

	http.HandleFunc("/body" , body)

	server.ListenAndServe()

}
