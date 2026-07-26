/*
Package main assembles configuration, keyboard input, and the terminal UI
into the Ditto executable.
*/
package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"

	"github.com/arvingarciabtw/ditto/internal/input"
	"github.com/arvingarciabtw/ditto/internal/tui"
)

/*
main wires together configuration, input capture, and the TUI so the program's
startup sequence remains visible in one place.
*/
func main() {
	cfg, err := loadConfig(os.Args[1:])
	if err != nil {
		exit(err)
	}

	program := tea.NewProgram(tui.InitModel(cfg))
	if err := input.StartInput(program); err != nil {
		exitInput(err)
	}

	if _, err := program.Run(); err != nil {
		exit(fmt.Errorf("run program: %w", err))
	}
}
