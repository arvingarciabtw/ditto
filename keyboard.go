package main

import (
	"fmt"

	evdev "github.com/gvalkov/golang-evdev"
)

func getKeyboard() (*evdev.InputDevice, error) {
	maxEvents := 32
	for e := range maxEvents {
		device, err := openDevice(e)
		if err != nil {
			continue
		}
		if isKeyboardDevice(device) {
			return device, nil
		}
		_ = device.File.Close()
	}

	return nil, fmt.Errorf("no keyboard device found. Try: sudo usermod -aG input $USER. Alternatively, try running the program with sudo")
}

func openDevice(e int) (*evdev.InputDevice, error) {
	path := fmt.Sprintf("/dev/input/event%d", e)
	device, err := evdev.Open(path)
	if err != nil {
		return nil, err
	}
	return device, nil
}

func isKeyboardDevice(device *evdev.InputDevice) bool {
	for cpbType, cpbCodes := range device.Capabilities {
		if cpbType.Type != evdev.EV_KEY {
			continue
		}
		for _, cpbCode := range cpbCodes {
			if cpbCode.Code == evdev.KEY_A {
				return true
			}
		}
	}
	return false
}
