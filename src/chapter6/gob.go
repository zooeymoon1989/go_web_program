package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type PostGob struct {
	Id int
	Content string
	Author string
}

func storeGob (data interface{} , filename string)  {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}

	//file ,err := os.OpenFile(filename , os.O_WRONLY | os.O_APPEND , 0644)
	//if err != nil {
	//	panic(err)
	//}
	//
	//defer file.Close()
	//
	//len ,err := file.WriteString(buffer.String())
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println(len)

	err = ioutil.WriteFile(filename , buffer.Bytes() , 0600)
	if err != nil {
		panic(err)
	}
}

func loadGob(data interface{} , filename string)  {
	raw , err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	post := PostGob{Id:1 , Content:"hello world!" , Author: "liwenqiang"}
	post2 := PostGob{Id:2 , Content:"hello world!" , Author: "liwenqiang2"}
	storeGob(post , "post1")
	storeGob(post2 , "post1")
	var postRead PostGob
	loadGob(&postRead , "post1")
	fmt.Println(postRead)
}