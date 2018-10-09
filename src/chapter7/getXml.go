package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type PostXml struct {
	XMLName xml.Name `xml:"post"`  //后面为结构标签，用来定义XML到结构体的映射
	Id string `xml:"id,attr"`
	Content string `xml:"content"`
	Author AuthorXml `xml:"author"`
	Comment []CommentXml `xml:"comments>comment"`
	Xml string `xml:",innerxml"`
}

type AuthorXml struct {
	Id string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type CommentXml struct {
	Id string `xml:"id,attr"`
	Content string `xml:"content"`
	Author string  `xml:"author"`
}


func main() {


	//简单使用unmarshal去解析xml文件
	//但是这里有一个问题，对于大型文件或者输入流，这种方法会影响效率


	//xmlFile , err := os.Open("src/chapter7/post.xml")
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//defer xmlFile.Close()
	//
	//xmlData , err := ioutil.ReadAll(xmlFile)
	//
	//if err != nil {
	//	fmt.Println("Error reading XML data:", err)
	//	return
	//}
	//
	//var post PostXml
	//xml.Unmarshal(xmlData , &post)
	//fmt.Println(post)



	//这里使用decoder来加速解析xml文件

	xmlFile , err := os.Open("src/chapter7/post.xml")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)

	for{
		t, err := decoder.Token()

		if err == io.EOF {//如果err指向文件的结尾，则说明没有多余的内容输入
			break
		}

		if err != nil{
			fmt.Println(err)
			return
		}

		switch se := t.(type) {
		case xml.StartElement:		//检查它们的起始元素
			if se.Name.Local == "comment" {
				var comment CommentXml
				decoder.DecodeElement(&comment , &se)
				fmt.Println(comment)
			}

		}

	}

}
