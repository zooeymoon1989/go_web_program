package main

import (
	"encoding/csv"
	"os"
	"strconv"
)


type PostCsv struct {
	Id     int
	Content string
	Author  string
}

func main() {

	csvFile , err := os.Create("posts.csv")
	if err != nil{
		panic(err)
	}

	defer csvFile.Close()
	//写入csv文件
	allPosts := []PostCsv{
		{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
		{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		{Id: 3, Content: "Hola Mundo!", Author: "Pedro"},
		{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}


	write := csv.NewWriter(csvFile)

	for _ , post := range allPosts {
		line := []string{strconv.Itoa(post.Id) , post.Content , post.Author}
		err := write.Write(line)
		if err != nil {
			panic(err)
		}
	}

	write.Flush()

	//读取csv文件

	file , err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record , err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []PostCsv

	for _ , item := range record {
		
	}


}