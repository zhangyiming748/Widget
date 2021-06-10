package format

import (
	"testing"
)

func TestFormatSymbol(t *testing.T) {
	src := "find.md"
	FormatSymbol(src)
}
func TestByte(t *testing.T) {
	str := "“"
	b := []byte(str)
	t.Log(str)
	t.Log(b)
}
