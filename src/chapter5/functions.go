package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func formatedDate(time time.Time)  string{
	layout := "2018-09-17"
	return time.Format(layout)
}


func processF(w http.ResponseWriter, r *http.Request) {
	tFuncMap := template.FuncMap {
		"fdate": formatedDate ,
	}
	t , err :=  template.New("funcs").Funcs(tFuncMap).ParseFiles("src/chapter5/templates/functions.html")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(r.Header["Connection"])
	t.Execute(w , time.Now())
}

func main() {

	server := http.Server{
		Addr:"127.0.0.1:9988",
	}

	http.HandleFunc("/parseFunc" , processF)

	server.ListenAndServe()

}