package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

/*
tempConfigPath redirects config operations into a test-owned directory so
tests never inspect or modify the user's real settings.
*/
func tempConfigPath(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", dir)
	return filepath.Join(dir, DirName, "config.json")
}

/*
writeConfigFile creates raw config input for load tests that need precise JSON
or file-state control.
*/
func writeConfigFile(t *testing.T, content string) string {
	t.Helper()
	path := tempConfigPath(t)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
	return path
}

/*
TestLoadConfig_missingFile verifies that first-run startup receives defaults
without treating the absent config as an error.
*/
func TestLoadConfig_missingFile(t *testing.T) {
	tempConfigPath(t)

	got, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	if want := Default(); !reflect.DeepEqual(got, want) {
		t.Errorf("Load() = %+v, want %+v", got, want)
	}
}

/*
TestLoadConfig_invalidJSON verifies malformed settings are reported instead
of being silently replaced by defaults and later overwritten.
*/
func TestLoadConfig_invalidJSON(t *testing.T) {
	writeConfigFile(t, "not json")

	_, err := Load()
	if err == nil {
		t.Fatal("Load() error = nil, want invalid JSON error")
	}
	if !strings.Contains(err.Error(), "decode config") {
		t.Errorf("Load() error = %q, want decode context", err)
	}
}

/*
TestLoadConfig_partialConfigUsesDefaults verifies omitted required settings
inherit defaults while explicitly stored values remain unchanged.
*/
func TestLoadConfig_partialConfigUsesDefaults(t *testing.T) {
	writeConfigFile(t, `{"active_layout":"dvorak","active_size":65}`)

	got, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}
	want := Default()
	want.ActiveLayout = "dvorak"
	want.ActiveSize = 65
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Load() = %+v, want %+v", got, want)
	}
}

/*
TestConfigValidate verifies each persisted selection is rejected when it is
outside the domain supplied by the application.
*/
func TestConfigValidate(t *testing.T) {
	validLayouts := []string{"qwerty", "dvorak", "Custom"}
	validSizes := []int{60, 75, 100}
	validStandards := []string{"ansi", "iso"}

	tests := []struct {
		name string
		cfg  Config
		want string
	}{
		{name: "valid", cfg: Config{ActiveLayout: "Custom", ActiveSize: 75, ActiveStandard: "iso", ActiveVisual: VisualBoxDraw}},
		{name: "layout", cfg: Config{ActiveLayout: "unknown", ActiveSize: 75, ActiveStandard: "ansi", ActiveVisual: VisualASCII}, want: `layout "unknown"`},
		{name: "size", cfg: Config{ActiveLayout: "qwerty", ActiveSize: -1, ActiveStandard: "ansi", ActiveVisual: VisualASCII}, want: "size -1"},
		{name: "standard", cfg: Config{ActiveLayout: "qwerty", ActiveSize: 75, ActiveStandard: "other", ActiveVisual: VisualASCII}, want: `standard "other"`},
		{name: "visual", cfg: Config{ActiveLayout: "qwerty", ActiveSize: 75, ActiveStandard: "ansi", ActiveVisual: "other"}, want: `visual "other"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Validate(validLayouts, validSizes, validStandards)
			if tt.want == "" {
				if err != nil {
					t.Fatalf("Validate() error = %v", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tt.want) {
				t.Errorf("Validate() error = %v, want error containing %q", err, tt.want)
			}
		})
	}
}

/*
TestSaveConfig_writesFile verifies saved data and temporary-file cleanup after
replacing the config file.
*/
func TestSaveConfig_writesFile(t *testing.T) {
	path := tempConfigPath(t)
	want := Config{ActiveLayout: "colemak", ActiveSize: 80, ActiveStandard: "iso", ActiveVisual: VisualASCII}

	if err := Save(want); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read saved config: %v", err)
	}
	var got Config
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("decode saved config: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("saved config = %+v, want %+v", got, want)
	}
	if !strings.Contains(string(data), `"active_visual": "ascii"`) {
		t.Error("saved config does not contain active_visual field")
	}
	entries, err := os.ReadDir(filepath.Dir(path))
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 || entries[0].Name() != filepath.Base(path) {
		t.Errorf("config directory entries = %v, want only %q", entries, filepath.Base(path))
	}
}

/*
TestSaveConfig_failure verifies path setup errors reach the caller rather than
presenting an unsaved configuration as persisted.
*/
func TestSaveConfig_failure(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", filepath.Join(dir, "not-a-directory"))
	if err := os.WriteFile(filepath.Join(dir, "not-a-directory"), nil, 0o600); err != nil {
		t.Fatal(err)
	}

	if err := Save(Default()); err == nil {
		t.Fatal("Save() error = nil, want path error")
	}
}

/*
TestSaveLoad_roundTrip verifies every config field survives persistence,
including optional boolean values whose false state must remain distinguishable.
*/
func TestSaveLoad_roundTrip(t *testing.T) {
	tempConfigPath(t)
	showInfo := false
	want := Config{
		ActiveLayout:   "azerty",
		ActiveSize:     100,
		ActiveStandard: "iso",
		ActiveVisual:   VisualBoxDraw,
		Locked:         true,
		ShowAllInfo:    &showInfo,
	}

	if err := Save(want); err != nil {
		t.Fatal(err)
	}
	got, err := Load()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("round trip = %+v, want %+v", got, want)
	}
}
