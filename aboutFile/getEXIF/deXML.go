package main

import (
	"github.com/beevik/etree"
)

func DecodeXML(xml string)string {
	// 初始化根节点
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(xml); err != nil {
		panic(err)
	}
	root := doc.SelectElement("response")
	res := root.FindElement("./regeocode/formatted_address").Text()
	//fmt.Println(res)
	return res
}
