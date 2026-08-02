/*
Package input captures system-wide keyboard events and converts them into the
physical keycodes consumed by Ditto's TUI.
*/
package input

// KeyMsg reports whether one physical key was pressed or released.
type KeyMsg struct {
	Code uint16
	Down bool
}
