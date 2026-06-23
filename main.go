package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/go-vgo/robotgo"
)

const logPrefix string = "[Teams Active Keeper]:"

func main() {
	const timeFlag string = "time"
	const defaultTime int = 4

	durationMinutes := flag.Int(timeFlag, defaultTime, "Mouse movement period (Minutes)")
	flag.Parse()

	fmt.Printf("%s Teams Active Keeper CLI Started (Period: %d mins)\n", logPrefix, *durationMinutes)
	fmt.Printf("%s To close it, you can close the terminal or press %s. \n", logPrefix, getShortcut())

	ticker := time.NewTicker(time.Duration(*durationMinutes) * time.Minute)
	defer ticker.Stop()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	//* Mouse movement
	moveMouse()

	for {
		select {
		case <-ticker.C:
			moveMouse()
		case <-sigChan:
			fmt.Printf("\n%s The application is closing. Teams status will return to normal", logPrefix)
			os.Exit(0)
		}
	}
}

func getShortcut() string {
	shortcut := "Ctrl+C"

	if runtime.GOOS == "darwin" {
		shortcut = "Control+C"
	}

	return shortcut
}

func moveMouse() {
	x, y := robotgo.Location()

	robotgo.MoveSmooth(x+50, y+50, .05, .05)
	robotgo.MoveSmooth(x, y, .05, .05)

	currentTime := time.Now().Format("15:04:05")
	fmt.Printf("%s - %s: The mouse was moved. \n", logPrefix, currentTime)
}
