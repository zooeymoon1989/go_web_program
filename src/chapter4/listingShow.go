package main

import (
	"fmt"
	"net/http"
)

func header(w http.ResponseWriter , r *http.Request)  {
	h := r.Header
	fmt.Fprintln(w,h)
}

func main() {

	server :=http.Server{
		Addr:"127.0.0.1:7777",
	}

	http.HandleFunc("/headers",header)

	server.ListenAndServe()

}
