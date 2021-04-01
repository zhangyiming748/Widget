package readline

import (
	"log"
	"testing"
)

func TestReadlink(t *testing.T) {
	fp := "/Users/zen/Github/Widget/youtube-dl/links.txt"
	ln := Readlink(fp)
	log.Println(ln)
}
