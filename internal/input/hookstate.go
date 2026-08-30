package input

import "github.com/arvingarciabtw/ditto/internal/keyboard/base"

/*
hookKeyEvents normalizes one hook transition. Darwin reports Caps Lock as one
alternating down/up event per actuation and can report a held modifier as down
when its paired modifier is released.
*/
func hookKeyEvents(held keyStateTracker, code uint16, down, darwin bool) []KeyEvent {
	if darwin && code == base.KEY_CAPSLOCK {
		delete(held, code)
		return []KeyEvent{
			{Code: code, State: KeyStatePressed},
			{Code: code, State: KeyStateReleased},
		}
	}

	if darwin && down && isHookModifier(code) {
		if _, alreadyHeld := held[code]; alreadyHeld {
			return []KeyEvent{{Code: code, State: held.release(code)}}
		}
	}

	var state KeyState
	if down {
		state = held.press(code)
	} else {
		state = held.release(code)
	}
	return []KeyEvent{{Code: code, State: state}}
}

// isHookModifier reports keys that macOS emits through aggregate modifier flags.
func isHookModifier(code uint16) bool {
	switch code {
	case base.KEY_LEFTSHIFT, base.KEY_RIGHTSHIFT,
		base.KEY_LEFTCTRL, base.KEY_RIGHTCTRL,
		base.KEY_LEFTALT, base.KEY_RIGHTALT,
		base.KEY_LEFTMETA, base.KEY_RIGHTMETA:
		return true
	default:
		return false
	}
}
