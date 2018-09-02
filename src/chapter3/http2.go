package main

import (
	"fmt"
    "golang.org/x/net/http2"
	"net/http"
)

type Http2Handler struct {}

func (h *Http2Handler) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"Hello World, http2!")
}

func main() {

	handler := Http2Handler{}
	server := http.Server{
		Addr:"127.0.0.1:7771",
		Handler:&handler,
	}



}



