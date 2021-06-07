package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
)

func main() {
	filename := "static.json"
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	json := string(data)
	value := gjson.Get(json, "retData.list搞")
	println(value.String())

}
