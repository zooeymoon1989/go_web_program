package main

import (
	json2 "encoding/json"
	"fmt"
	"net/http"
)

type jsonPost struct {
	User string
	Threads []string
}

func writeExample(w http.ResponseWriter , r *http.Request)  {
	str := `<html>
				<head><title>Go Web Programming</title></head>
				<body><h1>Hello World</h1></body>
			</html>`

	//如果在调用write方法之前制定content-type
	//前512字节会检测应该输出什么样的content-type
	w.Write([]byte(str))

}

func writeHeaderExample(w http.ResponseWriter , r *http.Request)  {
	//设置http response的状态码
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

func headerExample(w http.ResponseWriter , r *http.Request)  {
	//跳转方法
	w.Header().Set("Location" , "https://www.baidu.com")
	w.WriteHeader(302)
}

//写一个json相应
func jsonExample(w http.ResponseWriter , r *http.Request)  {
	w.Header().Set("Content-Type" , "application/json")
	post :=&jsonPost{
		User:"Li Wen Qiang",
		Threads: []string{"foo" , "bar"},
	}

	json , _ := json2.Marshal(post)

	w.Write(json)
}

func main() {
	server := http.Server{
		Addr:"127.0.0.1:9630",
	}

	http.HandleFunc("/write" , writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}
