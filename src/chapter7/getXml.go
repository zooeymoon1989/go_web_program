package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type PostXml struct {
	XMLName xml.Name `xml:"post"`  //后面为结构标签，用来定义XML到结构体的映射
	Id string `xml:"id,attr"`
	Content string `xml:"content"`
	Author string `xml:"author"`
	Xml string `xml:",innerxml"`

}

type AuthorXml struct {
	Id string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {

	xmlFile , err := os.Open("src/chapter7/post.xml")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer xmlFile.Close()

	xmlData , err := ioutil.ReadAll(xmlFile)

	if err != nil {
		fmt.Println("Error reading XML data:", err)
		return
	}

	var post PostXml
	xml.Unmarshal(xmlData , &post)
	fmt.Println(post)
}
