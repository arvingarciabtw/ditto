package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func tempConfigDir(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", dir)
	return filepath.Join(dir, DirName, "config.json")
}

func TestLoadConfig_missingFile(t *testing.T) {
	dir := t.TempDir()
	t.Setenv("XDG_CONFIG_HOME", dir)

	cfg := LoadConfig()
	if cfg.ActiveLayout != "qwerty" {
		t.Errorf("expected default layout qwerty, got %q", cfg.ActiveLayout)
	}
	if cfg.ActiveSize != 75 {
		t.Errorf("expected default size 75, got %d", cfg.ActiveSize)
	}
}

func TestLoadConfig_invalidJSON(t *testing.T) {
	path := tempConfigDir(t)
	os.MkdirAll(filepath.Dir(path), 0o755)
	os.WriteFile(path, []byte("not json"), 0o600)

	cfg := LoadConfig()
	if cfg.ActiveLayout != "qwerty" {
		t.Errorf("expected default layout on invalid JSON, got %q", cfg.ActiveLayout)
	}
}

func TestLoadConfig_valid(t *testing.T) {
	path := tempConfigDir(t)
	os.MkdirAll(filepath.Dir(path), 0o755)
	os.WriteFile(path, []byte(`{"active_layout":"dvorak","active_size":65}`), 0o600)

	cfg := LoadConfig()
	if cfg.ActiveLayout != "dvorak" {
		t.Errorf("expected dvorak, got %q", cfg.ActiveLayout)
	}
	if cfg.ActiveSize != 65 {
		t.Errorf("expected size 65, got %d", cfg.ActiveSize)
	}
}

func TestSaveConfig_writesFile(t *testing.T) {
	path := tempConfigDir(t)

	SaveConfig(Config{ActiveLayout: "colemak", ActiveSize: 80})

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("expected config file to exist: %v", err)
	}
	got := string(data)
	if !strings.Contains(got, `"active_layout": "colemak"`) {
		t.Errorf("expected colemak in output, got %s", got)
	}
	if !strings.Contains(got, `"active_size": 80`) {
		t.Errorf("expected size 80 in output, got %s", got)
	}
}

func TestSaveLoad_roundTrip(t *testing.T) {
	tempConfigDir(t)

	original := Config{ActiveLayout: "azerty", ActiveSize: 100}
	SaveConfig(original)

	loaded := LoadConfig()
	if loaded != original {
		t.Errorf("round trip failed: got %+v, want %+v", loaded, original)
	}
}


