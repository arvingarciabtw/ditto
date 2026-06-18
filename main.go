package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	evdev "github.com/gvalkov/golang-evdev"
)

func main() {
	p := tea.NewProgram(getInitModel())
	go listenToKeyboard(p)
	_, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}
}

func listenToKeyboard(p *tea.Program) {
	dev, err := getKeyboard()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		p.Quit()
		return
	}
	defer dev.File.Close()

	for {
		events, err := dev.Read()
		if err != nil {
			return
		}
		for _, ev := range events {
			if ev.Type == evdev.EV_KEY {
				p.Send(GlobalKeyMsg{
					Code: uint16(ev.Code),
					Down: ev.Value != 0,
				})
			}
		}
	}
}
