//go:build windows || darwin

package input

import (
	tea "charm.land/bubbletea/v2"
	hook "github.com/robotn/gohook"
)

/*
ListenHook forwards supported IOHook keyboard events to Bubble Tea while
discarding values that have no confirmed evdev equivalent.
*/
func ListenHook(p *tea.Program) {
	evChan := hook.Start()
	defer hook.End()

	for ev := range evChan {
		switch ev.Kind {
		case hook.KeyDown, hook.KeyUp:
			code, ok := mapKey(ev.Keycode)
			if !ok {
				continue
			}
			p.Send(KeyMsg{
				Code: code,
				Down: ev.Kind == hook.KeyDown,
			})
		}
	}
}
