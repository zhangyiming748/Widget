package readline

import (
	"bufio"
	"io"
	"os"
	. "youtube-dl/mylog"
)

func Readlink(fp string) []string {
	links := []string{}
	fi, err := os.Open(fp)
	if err != nil {
		Debug.Printf("Error: %s\n", err)
		return []string{}
	}
	defer fi.Close()
	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		//fmt.Println(string(a))
		links = append(links, string(a))
	}
	return links
}
