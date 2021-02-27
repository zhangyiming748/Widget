package readline

import (
	"log"
	"testing"
)

func TestReadlink(t *testing.T) {
	fp:="/Users/zen/Github/Tools/youtube-dl/links.txt"
	ln:=Readlink(fp)
	log.Println(ln)
}