package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/arvingarciabtw/ditto/internal/input"
)

/*
usage is kept with command error handling so help output and supported flags
remain synchronized at the process boundary.
*/
const usage = `Usage: ditto [options]

Options:
  -lock      lock settings toggles
  -unlock    unlock settings toggles
  -h, -help  show this help
`

/*
exit presents help or a general command error and terminates with the
appropriate status, keeping process-exit policy out of application logic.
*/
func exit(err error) {
	if !errors.Is(err, flag.ErrHelp) {
		exitError(err)
		return
	}
	if _, err := fmt.Fprint(os.Stdout, usage); err != nil {
		exitError(fmt.Errorf("write help: %w", err))
		return
	}
	os.Exit(0)
}

/*
exitError reports a command failure before terminating, checking the write
even though no reliable output channel remains if stderr itself fails.
*/
func exitError(err error) {
	if _, err := fmt.Fprintf(os.Stderr, "Error: %v\n", err); err != nil {
		os.Exit(1)
	}
	os.Exit(1)
}

/*
exitInput preserves the platform-specific startup guidance provided by the
input package before terminating after input initialization fails.
*/
func exitInput(err error) {
	input.PrintStartError(err)
	os.Exit(1)
}
