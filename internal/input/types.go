/*
Package input captures system-wide keyboard events and converts them into the
physical keycodes consumed by Ditto's TUI.
*/
package input

// KeyState describes one physical key transition reported by an input backend.
type KeyState uint8

const (
	// KeyStateUnknown represents an unsupported or uninitialized key transition.
	KeyStateUnknown KeyState = iota
	// KeyStateReleased reports that a physical key is no longer held.
	KeyStateReleased
	// KeyStatePressed reports the first down event for a physical key.
	KeyStatePressed
	// KeyStateRepeated reports an autorepeat event for a held physical key.
	KeyStateRepeated
)

// KeyEvent reports a physical keycode and its transition state.
type KeyEvent struct {
	Code  uint16
	State KeyState
}
