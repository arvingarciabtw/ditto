package input

import (
	"testing"

	"github.com/arvingarciabtw/ditto/internal/keyboard/base"
)

// TestHookKeyEvents verifies ordinary hook downs become presses and repeats.
func TestHookKeyEvents(t *testing.T) {
	held := make(keyStateTracker)

	events := hookKeyEvents(held, base.KEY_A, true, false)
	if len(events) != 1 || events[0].State != KeyStatePressed {
		t.Fatalf("first down events = %+v, want one press", events)
	}
	events = hookKeyEvents(held, base.KEY_A, true, false)
	if len(events) != 1 || events[0].State != KeyStateRepeated {
		t.Fatalf("second down events = %+v, want one repeat", events)
	}
	events = hookKeyEvents(held, base.KEY_A, false, false)
	if len(events) != 1 || events[0].State != KeyStateReleased {
		t.Fatalf("up events = %+v, want one release", events)
	}
}

// TestHookKeyEventsDarwinCapsLock verifies every toggle is a complete actuation.
func TestHookKeyEventsDarwinCapsLock(t *testing.T) {
	held := make(keyStateTracker)

	for _, down := range []bool{true, false} {
		events := hookKeyEvents(held, base.KEY_CAPSLOCK, down, true)
		if len(events) != 2 ||
			events[0].State != KeyStatePressed ||
			events[1].State != KeyStateReleased {
			t.Errorf("Caps Lock events = %+v, want press then release", events)
		}
	}
}

/*
TestHookKeyEventsDarwinPairedModifiers verifies an aggregate modifier flag does
not leave the released side held while its pair remains down.
*/
func TestHookKeyEventsDarwinPairedModifiers(t *testing.T) {
	held := make(keyStateTracker)

	leftDown := hookKeyEvents(held, base.KEY_LEFTSHIFT, true, true)
	rightDown := hookKeyEvents(held, base.KEY_RIGHTSHIFT, true, true)
	leftReleased := hookKeyEvents(held, base.KEY_LEFTSHIFT, true, true)
	rightReleased := hookKeyEvents(held, base.KEY_RIGHTSHIFT, false, true)

	states := []KeyState{
		leftDown[0].State,
		rightDown[0].State,
		leftReleased[0].State,
		rightReleased[0].State,
	}
	want := []KeyState{
		KeyStatePressed,
		KeyStatePressed,
		KeyStateReleased,
		KeyStateReleased,
	}
	for i := range want {
		if states[i] != want[i] {
			t.Errorf("state %d = %v, want %v", i, states[i], want[i])
		}
	}
}

// TestHookKeyEventsDarwinOrdinaryRepeat verifies Darwin still repeats normal keys.
func TestHookKeyEventsDarwinOrdinaryRepeat(t *testing.T) {
	held := make(keyStateTracker)
	hookKeyEvents(held, base.KEY_A, true, true)

	events := hookKeyEvents(held, base.KEY_A, true, true)
	if len(events) != 1 || events[0].State != KeyStateRepeated {
		t.Errorf("ordinary second down events = %+v, want one repeat", events)
	}
}
