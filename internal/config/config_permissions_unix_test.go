//go:build !windows

package config

import (
	"os"
	"testing"
)

/*
TestSaveConfig_writesPrivateFile verifies saved configs receive restrictive
Unix permissions.
*/
func TestSaveConfig_writesPrivateFile(t *testing.T) {
	path := tempConfigPath(t)
	if err := Save(Default()); err != nil {
		t.Fatal(err)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if perm := info.Mode().Perm(); perm != 0o600 {
		t.Errorf("saved config permissions = %o, want 600", perm)
	}
}

/*
TestSaveConfig_replacesExistingPermissions verifies an insecure existing Unix
mode is not retained when the config is replaced.
*/
func TestSaveConfig_replacesExistingPermissions(t *testing.T) {
	path := writeConfigFile(t, `{}`)
	if err := os.Chmod(path, 0o644); err != nil {
		t.Fatal(err)
	}

	if err := Save(Default()); err != nil {
		t.Fatal(err)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if perm := info.Mode().Perm(); perm != 0o600 {
		t.Errorf("saved config permissions = %o, want 600", perm)
	}
}
