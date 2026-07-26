package main

import (
	"errors"
	"flag"
	"testing"

	"github.com/arvingarciabtw/ditto/internal/config"
)

/*
TestParseOptions verifies that supported flags are isolated from positional
and unknown arguments, protecting the command's accepted input surface.
*/
func TestParseOptions(t *testing.T) {
	tests := []struct {
		name       string
		args       []string
		wantLock   bool
		wantUnlock bool
		wantErr    bool
	}{
		{name: "none"},
		{name: "lock", args: []string{"--lock"}, wantLock: true},
		{name: "unlock", args: []string{"--unlock"}, wantUnlock: true},
		{name: "both", args: []string{"--lock", "--unlock"}, wantLock: true, wantUnlock: true},
		{name: "help", args: []string{"--help"}, wantErr: true},
		{name: "unknown", args: []string{"--unknown"}, wantErr: true},
		{name: "positional", args: []string{"extra"}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseOptions(tt.args)
			if (err != nil) != tt.wantErr {
				t.Fatalf("parseOptions() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got.lock != tt.wantLock || got.unlock != tt.wantUnlock {
				t.Errorf("parseOptions() = %+v, want lock=%v unlock=%v", got, tt.wantLock, tt.wantUnlock)
			}
		})
	}
}

/*
TestParseOptions_help verifies that the standard help signal survives parsing
so the process layer can print usage and exit successfully.
*/
func TestParseOptions_help(t *testing.T) {
	_, err := parseOptions([]string{"-h"})
	if !errors.Is(err, flag.ErrHelp) {
		t.Errorf("parseOptions(-h) error = %v, want flag.ErrHelp", err)
	}
}

/*
TestApplyOptions verifies lock override precedence and persistence signaling
so command-line changes are applied predictably and saved only when requested.
*/
func TestApplyOptions(t *testing.T) {
	tests := []struct {
		name        string
		locked      bool
		options     options
		wantLocked  bool
		wantChanged bool
	}{
		{name: "none", locked: true, wantLocked: true},
		{name: "lock", options: options{lock: true}, wantLocked: true, wantChanged: true},
		{name: "unlock", locked: true, options: options{unlock: true}, wantChanged: true},
		{name: "both prefer unlock", options: options{lock: true, unlock: true}, wantChanged: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.Default()
			cfg.Locked = tt.locked
			got, changed := applyOptions(cfg, tt.options)
			if got.Locked != tt.wantLocked || changed != tt.wantChanged {
				t.Errorf("applyOptions() locked=%v changed=%v, want locked=%v changed=%v", got.Locked, changed, tt.wantLocked, tt.wantChanged)
			}
		})
	}
}
