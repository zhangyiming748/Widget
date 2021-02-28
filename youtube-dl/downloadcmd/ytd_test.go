package downloadcmd

import (
	"log"
	"sync"
	"testing"
)

func TestYtd(t *testing.T) {
	var wg sync.WaitGroup
	link := ""
	wg.Add(1)
	RunCommand(link, &wg, 1)
	wg.Wait()

}

func TestSplit(t *testing.T) {
	link := ""
	ret := split(link)
	log.Println(ret)
}
