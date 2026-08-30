//go:build windows || darwin

package input

import (
	"runtime"

	hook "github.com/robotn/gohook"
)

/*
ListenHook forwards supported IOHook keyboard events while discarding values
that have no confirmed evdev equivalent.
*/
func ListenHook(send func(KeyEvent)) {
	evChan := hook.Start()
	defer hook.End()
	held := make(keyStateTracker)

	for ev := range evChan {
		if ev.Kind != hook.KeyDown && ev.Kind != hook.KeyUp {
			// KeyHold is IOHook's typed-character event, not physical autorepeat.
			continue
		}

		code, ok := mapKey(ev.Keycode)
		if !ok {
			continue
		}

		events := hookKeyEvents(held, code, ev.Kind == hook.KeyDown, runtime.GOOS == "darwin")
		for _, event := range events {
			send(event)
		}
	}
}
