/*
Package config loads, validates, and atomically persists Ditto's user
settings so startup code receives a complete configuration or a clear error.
*/
package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
)

/*
Config contains the persisted user settings needed to reconstruct the TUI's
initial state across program runs.
*/
type Config struct {
	ActiveLayout   string `json:"active_layout"`
	ActiveSize     int    `json:"active_size"`
	ActiveStandard string `json:"active_standard"`
	ActiveVisual   string `json:"active_visual"`
	Locked         bool   `json:"locked"`
	ShowAllInfo    *bool  `json:"show_all_info,omitempty"`
}

const (
	// VisualASCII uses plain ASCII characters for borders and keyboard lines.
	VisualASCII = "ascii"
	// VisualBoxDraw uses Unicode box-drawing characters for smoother borders.
	VisualBoxDraw = "box-draw"
)

/*
DirName identifies Ditto's directory beneath the operating system's user
configuration root so config paths are assembled consistently.
*/
const DirName = "ditto"

/*
configPath resolves the config file lazily so environment changes in tests
and platform-specific user config locations are respected.
*/
func configPath() (string, error) {
	cfgDir := os.Getenv("XDG_CONFIG_HOME")
	if cfgDir == "" {
		var err error
		cfgDir, err = os.UserConfigDir()
		if err != nil {
			return "", err
		}
	}
	return filepath.Join(cfgDir, DirName, "config.json"), nil
}

/*
Default returns a complete baseline for settings that must have valid values,
allowing partial config files to inherit safe defaults during decoding.
*/
func Default() Config {
	return Config{
		ActiveLayout:   "qwerty",
		ActiveSize:     75,
		ActiveStandard: "ansi",
		ActiveVisual:   VisualASCII,
	}
}

/*
Load reads and decodes the user's config, treating only a missing file as a
first run while surfacing other failures to prevent silent data loss.
*/
func Load() (Config, error) {
	path, err := configPath()
	if err != nil {
		return Config{}, fmt.Errorf("config path: %w", err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Default(), nil
		}
		return Config{}, fmt.Errorf("read config: %w", err)
	}
	cfg := Default()
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("decode config: %w", err)
	}
	return cfg, nil
}

/*
Validate checks persisted selections against supported values, accepting
keyboard choices from the caller to avoid depending on its registries.
*/
func (cfg Config) Validate(layouts []string, sizes []int, standards []string) error {
	if !slices.Contains(layouts, cfg.ActiveLayout) {
		return fmt.Errorf("invalid active layout %q", cfg.ActiveLayout)
	}
	if !slices.Contains(sizes, cfg.ActiveSize) {
		return fmt.Errorf("invalid active size %d", cfg.ActiveSize)
	}
	if !slices.Contains(standards, cfg.ActiveStandard) {
		return fmt.Errorf("invalid active standard %q", cfg.ActiveStandard)
	}
	if cfg.ActiveVisual != VisualASCII && cfg.ActiveVisual != VisualBoxDraw {
		return fmt.Errorf("invalid active visual %q", cfg.ActiveVisual)
	}
	return nil
}

/*
Save encodes and atomically replaces the user's config so interrupted writes
cannot leave a partially written settings file.
*/
func Save(cfg Config) error {
	path, err := configPath()
	if err != nil {
		return fmt.Errorf("config path: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("mkdir config dir: %w", err)
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal config: %w", err)
	}
	if err := writeFileAtomic(path, data); err != nil {
		return fmt.Errorf("write config: %w", err)
	}
	return nil
}

/*
writeFileAtomic writes private data to a sibling temporary file before rename,
preserving the previous file until the replacement is complete.
*/
func writeFileAtomic(path string, data []byte) (err error) {
	f, err := os.CreateTemp(filepath.Dir(path), ".config-*.tmp")
	if err != nil {
		return err
	}
	tempPath := f.Name()
	defer func() {
		removeErr := os.Remove(tempPath)
		if removeErr == nil || errors.Is(removeErr, os.ErrNotExist) {
			return
		}
		err = errors.Join(err, fmt.Errorf("remove temporary config: %w", removeErr))
	}()

	closeOnError := func(operationErr error) error {
		closeErr := f.Close()
		if closeErr == nil {
			return operationErr
		}
		return errors.Join(operationErr, fmt.Errorf("close temporary config: %w", closeErr))
	}

	if err := f.Chmod(0o600); err != nil {
		return closeOnError(err)
	}
	if _, err := f.Write(data); err != nil {
		return closeOnError(err)
	}
	if err := f.Sync(); err != nil {
		return closeOnError(err)
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(tempPath, path)
}
