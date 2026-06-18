package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	ActiveLayout string `json:"active_layout"`
	ActiveSize   int    `json:"active_size"`
}

const configDirName = "qwerty-keyboard"

func configPath() (string, error) {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(cfgDir, configDirName, "config.json"), nil
}

func loadConfig() Config {
	path, err := configPath()
	if err != nil {
		return Config{ActiveLayout: "qwerty", ActiveSize: 75}
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{ActiveLayout: "qwerty", ActiveSize: 75}
	}
	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{ActiveLayout: "qwerty", ActiveSize: 75}
	}
	if cfg.ActiveLayout == "" {
		cfg.ActiveLayout = "qwerty"
	}
	if cfg.ActiveSize == 0 {
		cfg.ActiveSize = 75
	}
	return cfg
}

func saveConfig(cfg Config) {
	path, err := configPath()
	if err != nil {
		return
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return
	}
	err = os.WriteFile(path, data, 0o600)
	if err != nil {
		return
	}
}
