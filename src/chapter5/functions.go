package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func formatDate(time time.Time)  string{
	layout := "2018-09-17"
	return time.Format(layout)
}


func processF(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap { "fdate": formatDate }
	t := template.New("src/chapter5/templates/functions.html").Funcs(funcMap)
	t , err := template.ParseFiles("src/chapter5/templates/functions.html" )

	if err != nil{
		fmt.Println(err)
		return
	}
	t.Execute(w, time.Now())
}

func processFunc(w http.ResponseWriter , r *http.Request)  {
	funcMap := template.FuncMap{
		"fdate" : formatDate,
	}

	t := template.New("src/chapter5/templates/functions.html").Funcs(funcMap)
	t , err := template.ParseFiles("src/chapter5/templates/functions.html" )

	if err != nil{
		fmt.Println(err)
		return
	}

	t.Execute(w, time.Now())
}

func main() {

	server := http.Server{
		Addr:"127.0.0.1:9988",
	}

	http.HandleFunc("/parseFunc" , processF)

	server.ListenAndServe()

}