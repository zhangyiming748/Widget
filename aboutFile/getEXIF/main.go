package main

import (
	"errors"
	"fmt"
	"getEXIF/decodeGEO"
	"getEXIF/getGEO"
	"os"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			v := fmt.Sprint(err)
			if v == "runtime error: invalid memory address or nil pointer dereference" {
				getGEO.Debugf("也许当前图片没有GEO信息")
			}
		}
	}()
	if len(os.Args) < 2 {
		panic(errors.New("第二个参数需要指定图片路径"))
	}
	fp := os.Args[1]
	location := getGEO.EXIF2GEO(fp)
	decodeGEO.Decode(location)
	area := decodeGEO.DeXML("result.xml")
	fmt.Printf("大致位置是:%s\n", area)
}
