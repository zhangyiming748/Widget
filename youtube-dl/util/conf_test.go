package util

import (
	"log"
	"testing"
)

func TestGetVal(t *testing.T) {
	os := GetVal("os", "macOS")
	log.Println(os)
}
