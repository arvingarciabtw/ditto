package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

func main() {
	devs, err := devices()
	if err != nil {
		printDeviceError(err)
		os.Exit(1)
	}

	p := tea.NewProgram(initModel())
	for _, dev := range devs {
		go listenToKeyboard(p, dev)
	}
	_, err = p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}
