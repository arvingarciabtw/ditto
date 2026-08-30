//go:build windows || darwin

package input

import (
	"fmt"
	"os"
)

// StartInput starts system-wide hook capture without blocking the TUI startup.
func StartInput(send func(KeyEvent)) error {
	go ListenHook(send)
	return nil
}

// PrintStartError presents a hook startup error to the user.
func PrintStartError(err error) {
	if _, writeErr := fmt.Fprintf(os.Stderr, "Error: %v\n", err); writeErr != nil {
		return
	}
}
