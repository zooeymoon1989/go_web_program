package main

import (
	json2 "encoding/json"
	"fmt"
	"os"
)

type Post struct {
	Id       int	   `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	post :=Post{
		Id:1,
		Content:"foo bar",
		Author:Author{
			Id:2,
			Name:"SomeName2",
		},
		Comments:[]Comment{
			{
				Id:3,
				Content:"some 3",
				Author:"haha",
			},
			{
				Id:4,
				Content:"some 4",
				Author:"haha 4",
			},
		},
	}

	//output , err := json.MarshalIndent(&post , "", "\t")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//err = ioutil.WriteFile("src/chapter7/output.json" , output , 0644)
	//
	//if err != nil {
	//	fmt.Println("Error writing JSON to file:", err)
	//	return
	//}

	jsonFile , err := os.Create("src/chapter7/outputFromEncoder.json")
	if err != nil{
		fmt.Println(err)
		return
	}

	defer jsonFile.Close()

	encoder := json2.NewEncoder(jsonFile)
	encoder.SetIndent("","\t")

	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println(err)
	}


}
