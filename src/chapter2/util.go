package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

func session(w http.ResponseWriter , r *http.Request) (sess data.Session , err error) {

	//若果没有cookie的话，返回参数
	cookie , err := r.Cookie("_cookie")
	if err == nil {
		sess = data.Session{
			Uuid: cookie.Value ,
		}

		if ok , _ := sess.Check(); !ok { // 如果验证失败，输出错误信息
			err = errors.New("invalid session")
		}
	}
	return

}

/**
 生成模板页面
 */
func generateHTML(w http.ResponseWriter ,data interface{} , fn ... string)  {

	var files []string
	for _ , file := range fn{
		files = append(files , fmt.Sprintf("templates/%s.html",file))
	}

	templates := template.Must(template.ParseFiles(files ...))
	templates.ExecuteTemplate(w , "layout" , data)

}
