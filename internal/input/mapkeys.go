package input

import "github.com/arvingarciabtw/ditto/internal/keyboard/base"

/*
mapKey translates normalized IOHook virtual keycodes into evdev codes used by
the keyboard model. It rejects values whose equivalent is unknown rather than
guessing platform-specific meanings.
*/
func mapKey(code uint16) (uint16, bool) {
	switch code {
	case 0: // VC_UNDEFINED
		return 0, false
	case 0x0E1C: // VC_KP_ENTER
		return base.KEY_KPENTER, true
	case 0x0E1D: // VC_CONTROL_R
		return base.KEY_RIGHTCTRL, true
	case 0x0E35: // VC_KP_DIVIDE
		return base.KEY_KPSLASH, true
	case 0x0E37: // VC_PRINTSCREEN
		return base.KEY_SYSRQ, true
	case 0x0E38: // VC_ALT_R
		return base.KEY_RIGHTALT, true
	case 0x0E47, 0xEE47: // VC_HOME, Windows extended VC_HOME
		return base.KEY_HOME, true
	case 0x0E49, 0xEE49: // VC_PAGE_UP, Windows extended VC_PAGE_UP
		return base.KEY_PAGEUP, true
	case 0x0E4F, 0xEE4F: // VC_END, Windows extended VC_END
		return base.KEY_END, true
	case 0x0E51, 0xEE51: // VC_PAGE_DOWN, Windows extended VC_PAGE_DOWN
		return base.KEY_PAGEDOWN, true
	case 0x0E52, 0xEE52: // VC_INSERT, Windows extended VC_INSERT
		return base.KEY_INSERT, true
	case 0x0E53, 0xEE53: // VC_DELETE, Windows extended VC_DELETE
		return base.KEY_DELETE, true
	case 0x0E5B: // VC_META_L
		return base.KEY_LEFTMETA, true
	case 0x0E5C: // VC_META_R
		return base.KEY_RIGHTMETA, true
	case 0xE048, 0xEE48: // VC_UP, Windows extended VC_UP
		return base.KEY_UP, true
	case 0xE04B, 0xEE4B: // VC_LEFT, Windows extended VC_LEFT
		return base.KEY_LEFT, true
	case 0xE04D, 0xEE4D: // VC_RIGHT, Windows extended VC_RIGHT
		return base.KEY_RIGHT, true
	case 0xE050, 0xEE50: // VC_DOWN, Windows extended VC_DOWN
		return base.KEY_DOWN, true
	}

	if code >= base.KEY_ESC && code <= base.KEY_KPDOT {
		return code, true
	}
	if code == base.KEY_F11 || code == base.KEY_F12 {
		return code, true
	}
	return 0, false
}
