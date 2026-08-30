//go:build linux

package input

import (
	"fmt"
	"image/color"
	"os"
	"os/exec"
	"strings"

	lipgloss "charm.land/lipgloss/v2"
	evdev "github.com/gvalkov/golang-evdev"
)

const maxEventDevices = 32

func Devices() ([]*evdev.InputDevice, error) {
	var result []*evdev.InputDevice
	sawPermissionError := false

	for e := range maxEventDevices {
		device, err := evdev.Open(fmt.Sprintf("/dev/input/event%d", e))
		if err != nil {
			if os.IsPermission(err) {
				sawPermissionError = true
			}
			continue
		}
		if isKeyboardDevice(e) {
			result = append(result, device)
		} else {
			device.File.Close()
		}
	}

	if len(result) == 0 {
		if sawPermissionError {
			return nil, fmt.Errorf("permission denied reading input devices")
		}
		return nil, fmt.Errorf("no keyboard device found")
	}

	return result, nil
}

func isKeyboardDevice(eventNum int) bool {
	if checkUeventProperty(eventNum) {
		return true
	}
	return checkUdevadm(eventNum)
}

// ListenToKeyboard converts Linux evdev key events and sends supported states.
func ListenToKeyboard(send func(KeyEvent), dev *evdev.InputDevice) {
	defer dev.File.Close()

	for {
		events, err := dev.Read()
		if err != nil {
			return
		}
		for _, ev := range events {
			event, ok := keyEventFromEvdev(ev)
			if !ok {
				continue
			}
			send(event)
		}
	}
}

// keyEventFromEvdev maps the three defined evdev key values to input states.
func keyEventFromEvdev(ev evdev.InputEvent) (KeyEvent, bool) {
	if ev.Type != evdev.EV_KEY {
		return KeyEvent{}, false
	}

	var state KeyState
	switch ev.Value {
	case 0:
		state = KeyStateReleased
	case 1:
		state = KeyStatePressed
	case 2:
		state = KeyStateRepeated
	default:
		return KeyEvent{}, false
	}

	return KeyEvent{Code: ev.Code, State: state}, true
}

func readUeventFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	return string(data), err
}

func checkUeventProperty(eventNum int) bool {
	data, err := readUeventFile(fmt.Sprintf("/sys/class/input/event%d/device/uevent", eventNum))
	if err == nil && strings.Contains(data, "ID_INPUT_KEYBOARD=1") {
		return true
	}
	data, err = readUeventFile(fmt.Sprintf("/sys/class/input/event%d/uevent", eventNum))
	if err == nil && strings.Contains(data, "ID_INPUT_KEYBOARD=1") {
		return true
	}
	return false
}

func checkUdevadm(eventNum int) bool {
	out, err := exec.Command("udevadm", "info", "--query=property", fmt.Sprintf("--name=/dev/input/event%d", eventNum)).Output()
	if err != nil {
		return false
	}
	return strings.Contains(string(out), "ID_INPUT_KEYBOARD=1")
}

func PrintDeviceError(err error) {
	isDark := lipgloss.HasDarkBackground(os.Stdin, os.Stderr)

	var (
		red    color.Color
		green  color.Color
		blue   color.Color
		yellow color.Color
	)
	if isDark {
		red = lipgloss.BrightRed
		green = lipgloss.BrightGreen
		blue = lipgloss.BrightBlue
		yellow = lipgloss.BrightYellow
	} else {
		red = lipgloss.Red
		green = lipgloss.Green
		blue = lipgloss.Blue
		yellow = lipgloss.Yellow
	}

	errorLabel := lipgloss.NewStyle().Bold(true).Foreground(red)
	errorMsg := lipgloss.NewStyle().Foreground(red)
	dim := lipgloss.NewStyle().Faint(true)
	fixLabel := lipgloss.NewStyle().Bold(true).Foreground(yellow)
	noteLabel := lipgloss.NewStyle().Bold(true).Foreground(blue)
	cmd := lipgloss.NewStyle().Foreground(green)

	exe, exeErr := os.Executable()
	if exeErr != nil {
		exe = "ditto"
	}

	fmt.Fprintf(
		os.Stderr, "%s %s\n\n",
		errorLabel.Render("Error:"),
		errorMsg.Render(err.Error()),
	)

	fmt.Fprintf(os.Stderr, "%s\n\n", dim.Render(
		"This app reads raw evdev keyboard events directly (rather than through\na display server) in order to work inside the TUI. That requires\nread access to /dev/input/event*, which isn't readable by normal\nusers by default.",
	))

	fmt.Fprintf(
		os.Stderr, "%s %s\n\n",
		fixLabel.Render("Fix:"),
		cmd.Render("sudo setcap cap_dac_read_search=ep "+exe),
	)

	fmt.Fprintf(os.Stderr, "%s\n\n", dim.Render(
		"This grants read access to just this binary. It doesn't run as\nroot, just bypasses one permission check.",
	))

	fmt.Fprintf(
		os.Stderr, "%s %s\n\n",
		fixLabel.Render("Revoke anytime with:"),
		cmd.Render("sudo setcap -r "+exe),
	)

	fmt.Fprintf(
		os.Stderr, "%s %s\n",
		noteLabel.Render("Note:"),
		dim.Render("re-run this after rebuilding/reinstalling the binary."),
	)
}
