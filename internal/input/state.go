package input

// keyStateTracker distinguishes an initial key press from repeated down events.
type keyStateTracker map[uint16]struct{}

// press records a held key and returns whether this down event is a repeat.
func (t keyStateTracker) press(code uint16) KeyState {
	if _, held := t[code]; held {
		return KeyStateRepeated
	}

	t[code] = struct{}{}
	return KeyStatePressed
}

// release forgets a held key and reports its release state.
func (t keyStateTracker) release(code uint16) KeyState {
	delete(t, code)
	return KeyStateReleased
}
