package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type PostGXml struct {
	XMLName xml.Name `xml:"post"`
	Id string `xml:"id,attr"`
	Content string `xml:"content"`
	Author AuthorGXml `xml:"author"`
}

type AuthorGXml struct {
	Id string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {



	post := PostGXml{
		Id: "123",
		Content:"foo",
		Author:AuthorGXml{
			Id:"456",
			Name:"bar",
		},
	}

	//这里是一个简答的把结构体转化为struct的方法

	//
	//output , err := xml.MarshalIndent(&post , "" , "\t")
	//
	//if err != nil{
	//	fmt.Println(err)
	//	return
	//}
	//
	//err = ioutil.WriteFile("src/chapter7/something.xml" , []byte(xml.Header + string(output)) , 0644)
	//
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}


	//使用xml的encoder来生成xml文件
	file , err := os.Create("src/chapter7/xmlGenerateFromEncoder.xml")
	//写入xml头部
	file.Write([]byte(xml.Header))

	if err != nil{
		fmt.Println(err)
		return
	}

	encoder  := xml.NewEncoder(file)
	encoder.Indent("" , "\t")

	err = encoder.Encode(&post)

	if err != nil{
		fmt.Println(err)
		return
	}


}
