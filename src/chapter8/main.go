package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
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

func decode(filename string) (post PostJson , err error) {
	file , err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	decoder :=json.NewDecoder(file)
	err = decoder.Decode(&post)
	if err == io.EOF {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	return
}


func unmarshal(filename string) (post PostJson, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}
	json.Unmarshal(jsonData, &post)
	return
}

func handleRequest(w http.ResponseWriter , r *http.Request)  {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet( w , r)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter , r *http.Request) (err error) {
	id , err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	post , err := retrieve(id)
	if err != nil {
		return
	}

	output , err := json.MarshalIndent(&post , "" , "\t")

	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

