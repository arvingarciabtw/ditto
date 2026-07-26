package main

import (
	"flag"
	"fmt"
	"io"

	"github.com/arvingarciabtw/ditto/internal/config"
	"github.com/arvingarciabtw/ditto/internal/keyboard"
)

/*
options records command-line overrides separately from persisted settings so
omitted flags do not accidentally replace saved values.
*/
type options struct {
	lock   bool
	unlock bool
}

/*
loadConfig builds the effective configuration from disk and command-line
overrides, validating it before any input or terminal resources are started.
*/
func loadConfig(args []string) (config.Config, error) {
	opts, err := parseOptions(args)
	if err != nil {
		return config.Config{}, fmt.Errorf("parse flags: %w", err)
	}

	cfg, err := config.Load()
	if err != nil {
		return config.Config{}, fmt.Errorf("load config: %w", err)
	}

	if err := cfg.Validate(keyboard.LayoutListItems, keyboard.Sizes, keyboard.StandardListItems); err != nil {
		return config.Config{}, fmt.Errorf("validate config: %w", err)
	}

	cfg, changed := applyOptions(cfg, opts)
	if changed {
		if err := config.Save(cfg); err != nil {
			return config.Config{}, fmt.Errorf("save config: %w", err)
		}
	}

	return cfg, nil
}

/*
parseOptions parses an explicit argument slice with an isolated FlagSet so
callers and tests do not share process-global flag state.
*/
func parseOptions(args []string) (options, error) {
	var opts options

	flags := flag.NewFlagSet("ditto", flag.ContinueOnError)
	flags.SetOutput(io.Discard)
	flags.BoolVar(&opts.lock, "lock", false, "lock settings toggles")
	flags.BoolVar(&opts.unlock, "unlock", false, "unlock settings toggles")

	if err := flags.Parse(args); err != nil {
		return options{}, err
	}
	if flags.NArg() != 0 {
		return options{}, fmt.Errorf("unexpected argument %q", flags.Arg(0))
	}

	return opts, nil
}

/*
applyOptions applies explicit lock overrides and reports whether the result
must be persisted, allowing all CLI changes to be saved once.
*/
func applyOptions(cfg config.Config, opts options) (config.Config, bool) {
	if opts.lock {
		cfg.Locked = true
	}

	if opts.unlock {
		cfg.Locked = false
	}

	return cfg, opts.lock || opts.unlock
}
