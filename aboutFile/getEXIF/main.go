package main

import (
	"fmt"
	"getEXIF/decodeGEO"
	"getEXIF/getGEO"
)

func main() {
	fp := "IMG4.jpg"
	location := getGEO.EXIF2GEO(fp)
	decodeGEO.Decode(location)
	area := decodeGEO.DeXML("result.xml")
	fmt.Printf("大致位置是:%s", area)
}
