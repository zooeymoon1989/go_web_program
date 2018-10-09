package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type PostJson struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author AuthorJson `json:"author"`
	Comments []CommentJson `json:"comments"`
} 

type AuthorJson struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type CommentJson struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func main() {

	file , err := os.Open("src/chapter7/something.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	jsonData , err := ioutil.ReadAll(file)

	if err != nil {
		fmt.Println(err)
		return
	}

	var post PostJson

	json.Unmarshal(jsonData , &post)

	fmt.Println(post)


}
