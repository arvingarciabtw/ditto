package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	evdev "github.com/gvalkov/golang-evdev"
)

func main() {
	dev, err := device()
	if err != nil {
		printDeviceError(err)
		os.Exit(1)
	}

	p := tea.NewProgram(initModel())
	go listenToKeyboard(p, dev)
	_, err = p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}

func listenToKeyboard(p *tea.Program, dev *evdev.InputDevice) {
	defer dev.File.Close()

	for {
		events, err := dev.Read()
		if err != nil {
			return
		}
		for _, ev := range events {
			if ev.Type == evdev.EV_KEY {
				p.Send(globalKeyMsg{
					code: uint16(ev.Code),
					down: ev.Value != 0,
				})
			}
		}
	}
}

func printDeviceError(err error) {
	fmt.Fprintf(os.Stderr, `Error: %v

Adding the user to the input group grants read access to all input
devices (keyboards, mice, etc.). This is convenient but reduces
security — any process you run can log your input events.

It needs this access because it reads raw evdev keyboard events
directly (rather than through a display server or windowing system)
in order to work inside the TUI.

Options:
  1. Add user to the input group (convenient, less safe):
       sudo usermod -aG input $USER
     Then log out and back in.
  2. Run with sudo (safer, less convenient):
       sudo ./qwerty
`, err)
}
