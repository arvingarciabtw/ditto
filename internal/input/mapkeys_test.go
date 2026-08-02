package input

import (
	"testing"

	"github.com/arvingarciabtw/ditto/internal/keyboard/base"
)

/*
TestMapKeyIdentity verifies IOHook values already matching evdev pass through,
including lock and function keys that the previous mapper corrupted.
*/
func TestMapKeyIdentity(t *testing.T) {
	tests := []struct {
		name string
		code uint16
	}{
		{name: "escape", code: base.KEY_ESC},
		{name: "number", code: base.KEY_1},
		{name: "Q", code: base.KEY_Q},
		{name: "A", code: base.KEY_A},
		{name: "left shift", code: base.KEY_LEFTSHIFT},
		{name: "F1", code: base.KEY_F1},
		{name: "F11", code: base.KEY_F11},
		{name: "F12", code: base.KEY_F12},
		{name: "num lock", code: base.KEY_NUMLOCK},
		{name: "scroll lock", code: base.KEY_SCROLLLOCK},
		{name: "keypad zero", code: base.KEY_KP0},
		{name: "keypad decimal", code: base.KEY_KPDOT},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := mapKey(tt.code)
			if !ok {
				t.Fatalf("mapKey(%#x) reported unsupported", tt.code)
			}
			if got != tt.code {
				t.Errorf("mapKey(%#x) = %#x, want identity mapping", tt.code, got)
			}
		})
	}
}

/*
TestMapKeyExtended verifies confirmed IOHook extended values map to the evdev
codes used by the keyboard model.
*/
func TestMapKeyExtended(t *testing.T) {
	tests := []struct {
		name string
		code uint16
		want uint16
	}{
		{name: "keypad enter", code: 0x0E1C, want: base.KEY_KPENTER},
		{name: "right control", code: 0x0E1D, want: base.KEY_RIGHTCTRL},
		{name: "keypad divide", code: 0x0E35, want: base.KEY_KPSLASH},
		{name: "print screen", code: 0x0E37, want: base.KEY_SYSRQ},
		{name: "right alt", code: 0x0E38, want: base.KEY_RIGHTALT},
		{name: "home", code: 0x0E47, want: base.KEY_HOME},
		{name: "page up", code: 0x0E49, want: base.KEY_PAGEUP},
		{name: "end", code: 0x0E4F, want: base.KEY_END},
		{name: "page down", code: 0x0E51, want: base.KEY_PAGEDOWN},
		{name: "insert", code: 0x0E52, want: base.KEY_INSERT},
		{name: "delete", code: 0x0E53, want: base.KEY_DELETE},
		{name: "left meta", code: 0x0E5B, want: base.KEY_LEFTMETA},
		{name: "right meta", code: 0x0E5C, want: base.KEY_RIGHTMETA},
		{name: "up", code: 0xE048, want: base.KEY_UP},
		{name: "left", code: 0xE04B, want: base.KEY_LEFT},
		{name: "right", code: 0xE04D, want: base.KEY_RIGHT},
		{name: "down", code: 0xE050, want: base.KEY_DOWN},
		{name: "Windows home", code: 0xEE47, want: base.KEY_HOME},
		{name: "Windows up", code: 0xEE48, want: base.KEY_UP},
		{name: "Windows page up", code: 0xEE49, want: base.KEY_PAGEUP},
		{name: "Windows left", code: 0xEE4B, want: base.KEY_LEFT},
		{name: "Windows right", code: 0xEE4D, want: base.KEY_RIGHT},
		{name: "Windows end", code: 0xEE4F, want: base.KEY_END},
		{name: "Windows down", code: 0xEE50, want: base.KEY_DOWN},
		{name: "Windows page down", code: 0xEE51, want: base.KEY_PAGEDOWN},
		{name: "Windows insert", code: 0xEE52, want: base.KEY_INSERT},
		{name: "Windows delete", code: 0xEE53, want: base.KEY_DELETE},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := mapKey(tt.code)
			if !ok {
				t.Fatalf("mapKey(%#x) reported unsupported", tt.code)
			}
			if got != tt.want {
				t.Errorf("mapKey(%#x) = %#x, want %#x", tt.code, got, tt.want)
			}
		})
	}
}

/*
TestMapKeyRejectsUnsupported verifies undefined and speculative values are not
misrepresented as unrelated keyboard keys.
*/
func TestMapKeyRejectsUnsupported(t *testing.T) {
	tests := []struct {
		name string
		code uint16
	}{
		{name: "undefined", code: 0},
		{name: "keypad equals", code: 0x0E0D},
		{name: "pause", code: 0x0E45},
		{name: "context menu", code: 0x0E5D},
		{name: "keypad comma", code: 0x007E},
		{name: "yen", code: 0x007D},
		{name: "media play", code: 0xE022},
		{name: "old keypad comma alias", code: 0xE01C},
		{name: "old keypad dot alias", code: 0xE039},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := mapKey(tt.code)
			if ok {
				t.Errorf("mapKey(%#x) = %#x, want unsupported", tt.code, got)
			}
		})
	}
}

/*
TestMapKeyMacRegression verifies normalized macOS keycodes are not translated
again as native Apple HID values.
*/
func TestMapKeyMacRegression(t *testing.T) {
	tests := []struct {
		name string
		code uint16
	}{
		{name: "number one", code: base.KEY_1},
		{name: "Q", code: base.KEY_Q},
		{name: "A", code: base.KEY_A},
		{name: "S", code: base.KEY_S},
		{name: "left control", code: base.KEY_LEFTCTRL},
		{name: "left alt", code: base.KEY_LEFTALT},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := mapKey(tt.code)
			if !ok || got != tt.code {
				t.Errorf("mapKey(%#x) = (%#x, %v), want (%#x, true)", tt.code, got, ok, tt.code)
			}
		})
	}
}
