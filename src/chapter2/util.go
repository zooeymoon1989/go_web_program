package main

import (
	"errors"
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
