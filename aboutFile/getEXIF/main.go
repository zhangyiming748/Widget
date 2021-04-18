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
			fmt.Println(err)
		}
	}()
	if len(os.Args) < 2 {
		panic(errors.New("第二个参数需要指定图片路径"))
	} else {
		fmt.Println(os.Args[0], os.Args[1])
	}
	fp := os.Args[1]
	location := getGEO.EXIF2GEO(fp)
	decodeGEO.Decode(location)
	area := decodeGEO.DeXML("result.xml")
	fmt.Printf("大致位置是:%s\n", area)
}
