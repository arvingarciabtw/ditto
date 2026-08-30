package input

import "testing"

// TestKeyStateTracker verifies repeated downs are reset by a release.
func TestKeyStateTracker(t *testing.T) {
	tracker := make(keyStateTracker)

	if got := tracker.press(30); got != KeyStatePressed {
		t.Errorf("first press state = %v, want %v", got, KeyStatePressed)
	}
	if got := tracker.press(30); got != KeyStateRepeated {
		t.Errorf("second press state = %v, want %v", got, KeyStateRepeated)
	}
	if got := tracker.release(30); got != KeyStateReleased {
		t.Errorf("release state = %v, want %v", got, KeyStateReleased)
	}
	if got := tracker.press(30); got != KeyStatePressed {
		t.Errorf("press after release state = %v, want %v", got, KeyStatePressed)
	}
}

// TestKeyStateTrackerTracksCodesIndependently verifies one key does not affect another.
func TestKeyStateTrackerTracksCodesIndependently(t *testing.T) {
	tracker := make(keyStateTracker)
	tracker.press(30)

	if got := tracker.press(31); got != KeyStatePressed {
		t.Errorf("different key state = %v, want %v", got, KeyStatePressed)
	}
}
