package main

import (
	"fmt"
	"net/http"
)

type ThreeChainHandler struct {}

func (t ThreeChainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w , "Hello , I am three-chain-handler!")
}

func threeChainLog(h http.Handler) (http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter , r *http.Request) {
		fmt.Println("three chain log is occered!")
		h.ServeHTTP(w, r)
	})
}

func protect(h http.Handler) (http.Handler) {
	return http.HandlerFunc(func(w http.ResponseWriter , r *http.Request) {
		fmt.Println("protect function is occured!")
		h.ServeHTTP(w , r)
	})
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:7774",
	}

	hello := ThreeChainHandler{}
	http.Handle("/hello",protect(threeChainLog(hello)))
	http.Handle("/three",threeChainLog(hello))
	http.Handle("/test" , protect(hello))
	http.Handle("/normal" , hello)
	server.ListenAndServe()
	}

