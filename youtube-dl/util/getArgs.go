package getArgs

import (
	"log"
	"os"
)

func GetArgs() string {
	l := len(os.Args)
	if l < 2 {
		return ""
	}
	if l == 2 {
		log.Printf("参数%v",os.Args[1])
		return os.Args[l-1]
	}
		return ""
}
