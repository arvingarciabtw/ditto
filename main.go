package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting program...")

	// 1. Detect the user's keyboard device
	device, err := getKeyboard()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v.", err)
		os.Exit(1)
	}
	defer func() {
		if err := device.File.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Error while closing resource %s: %v", device.File.Name(), err)
			os.Exit(1)
		}
	}()

	// TODO: remove this once program is finished
	fmt.Println("Device: ", device)
}
