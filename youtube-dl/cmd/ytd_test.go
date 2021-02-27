package cmd

import (
	"sync"
	"testing"
)

func TestYtd(t *testing.T) {
	var wg sync.WaitGroup
	link := "https://www.xvideos.com/video21397485/leaked_full_video_shows_bigtitted_tranny_pumping_her_shecock"
	wg.Add(1)
	RunCommand(link, &wg, 1)
	wg.Wait()
}
