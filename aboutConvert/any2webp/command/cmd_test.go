package command

import (
	"sync"
	"testing"
)

func TestYtd(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)
	RunCmd("/Users/zen/Github/Widget/aboutConvert/any2webp/nier.jpg", &wg, 1)
	wg.Wait()

}
