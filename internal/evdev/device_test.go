package evdev

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadUeventFile_happyPath(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "uevent")
	if err := os.WriteFile(path, []byte("ID_INPUT_KEYBOARD=1\n"), 0644); err != nil {
		t.Fatal(err)
	}
	data, err := readUeventFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if data != "ID_INPUT_KEYBOARD=1\n" {
		t.Errorf("expected content, got %q", data)
	}
}

func TestReadUeventFile_missing(t *testing.T) {
	_, err := readUeventFile("/nonexistent/uevent")
	if err == nil {
		t.Error("expected error for missing file")
	}
}

func TestReadUeventFile_empty(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "uevent")
	if err := os.WriteFile(path, nil, 0644); err != nil {
		t.Fatal(err)
	}
	data, err := readUeventFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if data != "" {
		t.Errorf("expected empty string, got %q", data)
	}
}

func TestCheckInputGroup_notInGroup(t *testing.T) {
	err := CheckInputGroup()
	if err != nil {
		t.Logf("not in input group (expected in most test environments): %v", err)
	}
}
