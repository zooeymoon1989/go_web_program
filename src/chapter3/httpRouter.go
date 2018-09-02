package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func helloRouter(w http.ResponseWriter , r *http.Request , p httprouter.Params)  {

	fmt.Fprintf(w , "Hello , %s!\n" , p.ByName("name"))//获取get传入的参数

}

func main() {

	mux := httprouter.New()
	mux.GET("/hello/:name" , helloRouter)

	server := http.Server{
		Addr: "127.0.0.1:7773",
		Handler:mux,
	}

	server.ListenAndServe()

}
