package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter , r *http.Request)  {

	var cookie1 = http.Cookie{
		Name:"foo",
		Value:"bar",
		HttpOnly:true,
	}

	var cookie2 = http.Cookie{
		Name:"foo1",
		Value:"bar1",
		HttpOnly:true,
	}

	w.Header().Set("Set-Cookie" , cookie1.String())
	//w.Header().Add("Set-Cookie" , cookie2.String())

	http.SetCookie( w , &cookie2)

}

//获取cookie方法
func getCookies(w http.ResponseWriter , r *http.Request)  {
	//httpCookie := r.Header["Cookie"]
	////返回是一个切片字符串数组
	//fmt.Fprintln(w , httpCookie)


	cookieAll := r.Cookies()
	cookie1 , err:= r.Cookie("foo")

	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}

	fmt.Fprintln(w , cookieAll)
	fmt.Fprintln(w , cookie1)


}

func main() {
	server := http.Server{
		Addr:"127.0.0.1:8888",
	}

	http.HandleFunc("/cookie",setCookie)

	http.HandleFunc("/getCookie" , getCookies)

	server.ListenAndServe()
}
