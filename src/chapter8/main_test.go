package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestDecode(t *testing.T)  {
	post , err := decode("something.json")
	if err !=nil{
		t.Error(err)
	}

	if post.Comments[0].Id != 3{
		t.Error("something is wrong with ID")
	}
}

func TestEncode(t *testing.T) {
	t.Skip("Skipping encoding for now")
}

/**
 * 模拟单元测试时间过长是否可以忽略的问题
 */
func TestLongRunningTest(t *testing.T)  {
	if testing.Short() {  //测试时查看命令是否在参数中携带-short
		t.Skip("Skipping long-running test in short mode")
	}

	time.Sleep(10 * time.Second)
}


func TestParallel_1(t *testing.T)  {
	t.Parallel()
	time.Sleep( 1 * time.Second)
}

func TestParallel_2(t *testing.T)  {
	t.Parallel()
	time.Sleep( 2 * time.Second)
}

func TestParallel_3(t *testing.T)  {
	t.Parallel()
	time.Sleep( 3 * time.Second)
}

func TestHandleGet(t *testing.T)  {

	mux := http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest)

	writer := httptest.NewRecorder()   //创建一个ResponseRecorder结构体，这个结构体用来检查保存的返回
	request, _ := http.NewRequest("GET", "/post/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}


	var post PostJson
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Error("Cannot retrieve JSON post")
	}


}