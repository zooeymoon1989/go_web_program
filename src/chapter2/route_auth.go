package main

import "net/http"

func authenticate(w http.ResponseWriter , r *http.Request)  {

	//解析表单
	r.ParseForm()
	//通过给定的email获取用户的struct
	user , _ := data.UserByEmail(r.PostFormValue("email"))
	//如果用户输入的密码和数据库里面一致
	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		//创建session
		session := user.CreateSession()
		//创建cookie
		cookie := http.Cookie{
			Name:	"_cookie",
			Value: 	session.Uuid,
			HttpOnly: 	true,//只有http和https才能获取cookie
		}
		//设置cookie
		http.SetCookie(w, &cookie)
		//重定向到首页
		http.Redirect(w , r , "/" , 302)

	}else {
		//如果登录失败，重定向到login页面
		http.Redirect(w, r, "/login", 302)
	}

}