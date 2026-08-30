//go:build linux

package input

import (
	"os"
	"path/filepath"
	"testing"

	evdev "github.com/gvalkov/golang-evdev"
)

// TestKeyEventFromEvdev verifies only defined Linux key states are forwarded.
func TestKeyEventFromEvdev(t *testing.T) {
	tests := []struct {
		name      string
		eventType uint16
		value     int32
		state     KeyState
		ok        bool
	}{
		{name: "release", eventType: evdev.EV_KEY, value: 0, state: KeyStateReleased, ok: true},
		{name: "press", eventType: evdev.EV_KEY, value: 1, state: KeyStatePressed, ok: true},
		{name: "repeat", eventType: evdev.EV_KEY, value: 2, state: KeyStateRepeated, ok: true},
		{name: "unknown value", eventType: evdev.EV_KEY, value: 3},
		{name: "non-key event", eventType: evdev.EV_SYN, value: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			event, ok := keyEventFromEvdev(evdev.InputEvent{
				Type:  tt.eventType,
				Code:  30,
				Value: tt.value,
			})
			if ok != tt.ok {
				t.Fatalf("keyEventFromEvdev() ok = %v, want %v", ok, tt.ok)
			}
			if !ok {
				return
			}
			if event.Code != 30 || event.State != tt.state {
				t.Errorf("keyEventFromEvdev() = %+v, want code 30 and state %v", event, tt.state)
			}
		})
	}
}

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
