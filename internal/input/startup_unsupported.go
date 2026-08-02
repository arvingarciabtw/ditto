//go:build !linux && !windows && !darwin

package input

import (
	"fmt"
	"os"
	"runtime"

	tea "charm.land/bubbletea/v2"
)

// StartInput reports that Ditto has no input backend for the current platform.
func StartInput(_ *tea.Program) error {
	return fmt.Errorf("keyboard input is unsupported on %s", runtime.GOOS)
}

// PrintStartError presents an unsupported-platform startup error to the user.
func PrintStartError(err error) {
	if _, writeErr := fmt.Fprintf(os.Stderr, "Error: %v\n", err); writeErr != nil {
		return
	}
}
