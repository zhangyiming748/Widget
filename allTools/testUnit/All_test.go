package testUnit

import (
	"fmt"
	"github.com/go-cmd/cmd"
	"testing"
	"time"
)

func TestGo_cmd(t *testing.T) {
	// Start a long-running process, capture stdout and stderr
	findCmd := cmd.NewCmd("find", "/", "--name", "needle")
	statusChan := findCmd.Start() // non-blocking

	ticker := time.NewTicker(2 * time.Second)

	// Print last line of stdout every 2s
	go func() {
		for range ticker.C {
			status := findCmd.Status()
			n := len(status.Stdout)
			fmt.Println(status.Stdout[n-1])
		}
	}()

	// Stop command after 1 hour
	go func() {
		<-time.After(1 * time.Hour)
		findCmd.Stop()
	}()

	// Check if command is done
	select {
	case finalStatus := <-statusChan:
		// done
		t.Logf("finish%v", finalStatus)
	default:
		// no, still running
	}

	// Block waiting for command to exit, be stopped, or be killed
	finalStatus := <-statusChan
	t.Logf("finish%v", finalStatus)

}
func TestPrint(t *testing.T) {
	printInLine()
}
func printInLine() {
	for i := 0; i < 100; i++ {
		fmt.Printf("这是第%d次输出", i)
		time.Sleep(10 * time.Millisecond)
		fmt.Println("\r\033[k")
	}
}
