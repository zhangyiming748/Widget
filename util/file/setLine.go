package file

import (
	"Widget/util/log"
	"os"
	"strings"
)

func WriteLines(fname string, s []string) {
	f, err := os.OpenFile(fname, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0776)
	if err != nil {
		log.Debug.Println(err)
	}
	defer f.Close()
	for _, v := range s {
		_, err := f.WriteString(strings.Join([]string{v, "\n"}, ""))
		if err != nil {
			return
		}
	}
}
