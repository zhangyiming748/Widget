package timeNow

import (
	"log"
	"testing"
)

func TestDateNowFormatStr(t *testing.T) {
	tn := DateNowFormatStr()
	log.Println(tn)
}
