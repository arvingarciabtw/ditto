//go:build linux

package input

// StartInput opens Linux keyboards and starts one event listener per device.
func StartInput(send func(KeyEvent)) error {
	devs, err := Devices()
	if err != nil {
		return err
	}
	for _, dev := range devs {
		go ListenToKeyboard(send, dev)
	}
	return nil
}

func PrintStartError(err error) {
	PrintDeviceError(err)
}
