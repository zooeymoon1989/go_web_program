package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func chainHello(w http.ResponseWriter , r *http.Request)  {

	fmt.Fprintf(w , "Hello!")

}

func log(h http.HandlerFunc)  (http.HandlerFunc){
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w , r)
	}
}

func protectChain(h http.HandlerFunc) (http.HandlerFunc) {

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello world!")
		h(w , r)
	}

}


func main() {



	server := http.Server{
		Addr: "127.0.0.1:7775",
	}
	http.HandleFunc("/hello" , protectChain(log(chainHello)))
	server.ListenAndServe()
}