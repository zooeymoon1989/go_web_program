package main

import (
	"html/template"
	"net/http"
)

func main() {

	//创建一个multiplexer，这个叫多路复用器，其实我也没有明白
	mux := http.NewServeMux()
	//这个方法从指定的文件夹下提供文件服务
	files := http.FileServer(http.Dir("/public"))
	//把这个文件handler放入到多路复用器中
	//使用http的StripPrefix去掉url前缀
	//这个方法的作用时如果发现请求的url地址去除url前缀之后是以static开头，我们应该到public文件夹下查找对应的文件
	mux.Handle("/static",http.StripPrefix("/static/",files))

	//重定向到root URL使用一个HandleFunc方法
	mux.HandleFunc("/",index)

	server := &http.Server{
		Addr: "0.0.0.0:9877",
		Handler: mux ,
	}

	server.ListenAndServe()

}


func index(w http.ResponseWriter, r *http.Request) {

	var templates *template.Template

	threads , err := data.Threads(); if err == nil{

		_ , err := session(w , r)

		//未授权的页面显示
		publicTmplFiles := []string{
			"templates/layout.html",
			"templates/public.navbar.html",
			"templates/index.html",
		}
		//授权的页面显示
		privateTmplFiles := []string{
			"templates/layout.html",
			"templates/private.navbar.html",
			"templates/index.html",
		}

		//Must在解析文件出现错误的时候，会调用panic方法
		if err != nil {
			templates = template.Must(template.ParseFiles(privateTmplFiles ...))
		}else{
			templates = template.Must(template.ParseFiles(publicTmplFiles ...))
		}

		templates.ExecuteTemplate(w , "layout" , threads)
	}

}


