package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"net/http/httptest"
	"os"
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

var mux *http.ServeMux
var writer *httptest.ResponseRecorder

func TestMain(m *testing.M)  {  //这个函数是用来解决一些服用代码
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp()  {
	var err error
	db , err :=sql.Open("mysql" , "root:@/gwp")
	//db , err :=sql.Open("mysql" , "root:@/gwp")
	if err != nil {
		panic(err)
	}

	mux = http.NewServeMux()
	mux.HandleFunc("/post/", handleRequest(&Post{Db:db}))
	writer = httptest.NewRecorder()
}


func TestHandleGet(t *testing.T)  {

	request, _ := http.NewRequest("GET", "/post/1", nil)  //创建请求到你想要测试的handler
	mux.ServeHTTP(writer, request)  //发送请求到tested handler

	//json := strings.NewReader(`{"content":"Updated post","author":"Sau Sheong"}`)
	//request , _ := http.NewRequest("PUT", "/post/1", json)
	//mux.ServeHTTP(writer , request)

	if writer.Code != 200 {
		fmt.Println(writer.Body)
		t.Errorf("Response code is %v", writer.Code)

	}


	var post PostJson
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Error("id is not correct")
	}


}