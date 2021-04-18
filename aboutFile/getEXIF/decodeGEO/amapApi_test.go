package decodeGEO

import (
	"getEXIF"
	"testing"
)

func TestDecode(t *testing.T) {
	location:="116.36800384499999,39.91048431388889"
	Decode(location)
}
func TestDecodeXML(t *testing.T) {
	xmlFile:="result.xml"
	main.DecodeXML(xmlFile)
}