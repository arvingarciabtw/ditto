package base

var (
	// ALPHANUMERIC KEYS
	KeyA = Key{Label: "A", Width: U1, Finger: Pinky, EvCode: KEY_A}
	KeyB = Key{Label: "B", Width: U1, Finger: Index, EvCode: KEY_B}
	KeyC = Key{Label: "C", Width: U1, Finger: Middle, EvCode: KEY_C}
	KeyD = Key{Label: "D", Width: U1, Finger: Middle, EvCode: KEY_D}
	KeyE = Key{Label: "E", Width: U1, Finger: Middle, EvCode: KEY_E}
	KeyF = Key{Label: "F", Width: U1, Finger: Index, EvCode: KEY_F}
	KeyG = Key{Label: "G", Width: U1, Finger: Index, EvCode: KEY_G}
	KeyH = Key{Label: "H", Width: U1, Finger: Index, EvCode: KEY_H}
	KeyI = Key{Label: "I", Width: U1, Finger: Middle, EvCode: KEY_I}
	KeyJ = Key{Label: "J", Width: U1, Finger: Index, EvCode: KEY_J}
	KeyK = Key{Label: "K", Width: U1, Finger: Middle, EvCode: KEY_K}
	KeyL = Key{Label: "L", Width: U1, Finger: Ring, EvCode: KEY_L}
	KeyM = Key{Label: "M", Width: U1, Finger: Index, EvCode: KEY_M}
	KeyN = Key{Label: "N", Width: U1, Finger: Index, EvCode: KEY_N}
	KeyO = Key{Label: "O", Width: U1, Finger: Ring, EvCode: KEY_O}
	KeyP = Key{Label: "P", Width: U1, Finger: Pinky, EvCode: KEY_P}
	KeyQ = Key{Label: "Q", Width: U1, Finger: Pinky, EvCode: KEY_Q}
	KeyR = Key{Label: "R", Width: U1, Finger: Index, EvCode: KEY_R}
	KeyS = Key{Label: "S", Width: U1, Finger: Ring, EvCode: KEY_S}
	KeyT = Key{Label: "T", Width: U1, Finger: Index, EvCode: KEY_T}
	KeyU = Key{Label: "U", Width: U1, Finger: Index, EvCode: KEY_U}
	KeyV = Key{Label: "V", Width: U1, Finger: Index, EvCode: KEY_V}
	KeyW = Key{Label: "W", Width: U1, Finger: Ring, EvCode: KEY_W}
	KeyX = Key{Label: "X", Width: U1, Finger: Ring, EvCode: KEY_X}
	KeyY = Key{Label: "Y", Width: U1, Finger: Index, EvCode: KEY_Y}
	KeyZ = Key{Label: "Z", Width: U1, Finger: Pinky, EvCode: KEY_Z}
	Key0 = Key{Label: "0", Width: U1, Finger: Pinky, EvCode: KEY_0}
	Key1 = Key{Label: "1", Width: U1, Finger: Pinky, EvCode: KEY_1}
	Key2 = Key{Label: "2", Width: U1, Finger: Ring, EvCode: KEY_2}
	Key3 = Key{Label: "3", Width: U1, Finger: Middle, EvCode: KEY_3}
	Key4 = Key{Label: "4", Width: U1, Finger: Index, EvCode: KEY_4}
	Key5 = Key{Label: "5", Width: U1, Finger: Index, EvCode: KEY_5}
	Key6 = Key{Label: "6", Width: U1, Finger: Index, EvCode: KEY_6}
	Key7 = Key{Label: "7", Width: U1, Finger: Index, EvCode: KEY_7}
	Key8 = Key{Label: "8", Width: U1, Finger: Middle, EvCode: KEY_8}
	Key9 = Key{Label: "9", Width: U1, Finger: Ring, EvCode: KEY_9}

	// CONTROL KEYS
	KeyEsc        = Key{Label: "Esc", Width: U1, Finger: Pinky, EvCode: KEY_ESC}
	KeyLeftShift  = Key{Label: "Shift", Width: U2p5, Finger: Pinky, EvCode: KEY_LEFTSHIFT}
	KeyRightShift = Key{Label: "Shift", Width: U3, Finger: Pinky, EvCode: KEY_RIGHTSHIFT}
	KeyLeftCtrl   = Key{Label: "Ctrl", Width: U2, Finger: Pinky, EvCode: KEY_LEFTCTRL}
	KeyRightCtrl  = Key{Label: "Ctrl", Width: U2, Finger: Pinky, EvCode: KEY_RIGHTCTRL}
	KeyLeftMeta   = Key{Label: "⌘", Width: U1, Finger: Ring, EvCode: KEY_LEFTMETA}
	KeyRightMeta  = Key{Label: "⌘", Width: U1, Finger: Ring, EvCode: KEY_RIGHTMETA}
	KeyLeftAlt    = Key{Label: "Alt", Width: U1p75, Finger: Thumb, EvCode: KEY_LEFTALT}
	KeyRightAlt   = Key{Label: "Alt", Width: U1p75, Finger: Thumb, EvCode: KEY_RIGHTALT}
	KeyCapsLock   = Key{Label: "Caps", Width: U2, Finger: Pinky, EvCode: KEY_CAPSLOCK}
	KeyNumLock    = Key{Label: "Nlk", Width: U1, Finger: Index, EvCode: KEY_NUMLOCK}
	KeyScrollLock = Key{Label: "⚲", Width: U1, Finger: Middle, EvCode: KEY_SCROLLLOCK}

	// FUNCTION KEYS
	KeyFn  = Key{Label: "Fn", Width: U1, Finger: Ring, EvCode: KEY_FN}
	KeyF1  = Key{Label: "F1", Width: U1, Finger: Pinky, EvCode: KEY_F1}
	KeyF2  = Key{Label: "F2", Width: U1, Finger: Ring, EvCode: KEY_F2}
	KeyF3  = Key{Label: "F3", Width: U1, Finger: Middle, EvCode: KEY_F3}
	KeyF4  = Key{Label: "F4", Width: U1, Finger: Index, EvCode: KEY_F4}
	KeyF5  = Key{Label: "F5", Width: U1, Finger: Index, EvCode: KEY_F5}
	KeyF6  = Key{Label: "F6", Width: U1, Finger: Index, EvCode: KEY_F6}
	KeyF7  = Key{Label: "F7", Width: U1, Finger: Index, EvCode: KEY_F7}
	KeyF8  = Key{Label: "F8", Width: U1, Finger: Middle, EvCode: KEY_F8}
	KeyF9  = Key{Label: "F9", Width: U1, Finger: Ring, EvCode: KEY_F9}
	KeyF10 = Key{Label: "F10", Width: U1, Finger: Pinky, EvCode: KEY_F10}
	KeyF11 = Key{Label: "F11", Width: U1, Finger: Pinky, EvCode: KEY_F11}
	KeyF12 = Key{Label: "F12", Width: U1, Finger: Pinky, EvCode: KEY_F12}

	// NAVIGATION AND EDITING KEYS
	KeyUp        = Key{Label: "↑", Width: U1, Finger: Middle, EvCode: KEY_UP}
	KeyDown      = Key{Label: "↓", Width: U1, Finger: Middle, EvCode: KEY_DOWN}
	KeyLeft      = Key{Label: "←", Width: U1, Finger: Index, EvCode: KEY_LEFT}
	KeyRight     = Key{Label: "→", Width: U1, Finger: Ring, EvCode: KEY_RIGHT}
	KeyPageUp    = Key{Label: "⇡", Width: U1, Finger: Ring, EvCode: KEY_PAGEUP}
	KeyPageDown  = Key{Label: "⇣", Width: U1, Finger: Ring, EvCode: KEY_PAGEDOWN}
	KeyHome      = Key{Label: "⌂", Width: U1, Finger: Middle, EvCode: KEY_HOME}
	KeyEnd       = Key{Label: "⌿", Width: U1, Finger: Middle, EvCode: KEY_END}
	KeyInsert    = Key{Label: "Ins", Width: U1, Finger: Index, EvCode: KEY_INSERT}
	KeyDelete    = Key{Label: "Del", Width: U1, Finger: Index, EvCode: KEY_DELETE}
	KeyBackspace = Key{Label: "<--", Width: U2p5, Finger: Pinky, EvCode: KEY_BACKSPACE}

	// SPECIAL AND PUNCTUATION KEYS
	KeyTab        = Key{Label: "Tab", Width: U1p75, Finger: Pinky, EvCode: KEY_TAB}
	KeyEnter      = Key{Label: "Enter↵", Width: U3, Finger: Pinky, EvCode: KEY_ENTER}
	KeySpace      = Key{Label: "Space", Width: U5p75, Finger: Thumb, EvCode: KEY_SPACE}
	KeyGrave      = Key{Label: "`", Width: U1, Finger: Pinky, EvCode: KEY_GRAVE}
	KeyMinus      = Key{Label: "-", Width: U1, Finger: Pinky, EvCode: KEY_MINUS}
	KeyEqual      = Key{Label: "=", Width: U1, Finger: Pinky, EvCode: KEY_EQUAL}
	KeyLeftBrace  = Key{Label: "[", Width: U1, Finger: Pinky, EvCode: KEY_LEFTBRACE}
	KeyRightBrace = Key{Label: "]", Width: U1, Finger: Pinky, EvCode: KEY_RIGHTBRACE}
	KeyBackslash  = Key{Label: "\\", Width: U1p75, Finger: Pinky, EvCode: KEY_BACKSLASH}
	KeySemicolon  = Key{Label: ";", Width: U1, Finger: Pinky, EvCode: KEY_SEMICOLON}
	KeyApostrophe = Key{Label: "'", Width: U1, Finger: Pinky, EvCode: KEY_APOSTROPHE}
	KeyComma      = Key{Label: ",", Width: U1, Finger: Middle, EvCode: KEY_COMMA}
	KeyDot        = Key{Label: ".", Width: U1, Finger: Ring, EvCode: KEY_DOT}
	KeySlash      = Key{Label: "/", Width: U1, Finger: Pinky, EvCode: KEY_SLASH}
	KeyPound      = Key{Label: "#", Width: U1, Finger: Pinky, EvCode: KEY_NUMERIC_POUND}
	KeyAcute      = Key{Label: "´", Width: U1, Finger: Pinky, EvCode: KEY_RESERVED}

	// NUMPAD KEYS
	KeyPad0        = Key{Label: "0", Width: U2p5, Finger: Thumb, EvCode: KEY_KP0}
	KeyPad1        = Key{Label: "1", Width: U1, Finger: Index, EvCode: KEY_KP1}
	KeyPad2        = Key{Label: "2", Width: U1, Finger: Middle, EvCode: KEY_KP2}
	KeyPad3        = Key{Label: "3", Width: U1, Finger: Ring, EvCode: KEY_KP3}
	KeyPad4        = Key{Label: "4", Width: U1, Finger: Index, EvCode: KEY_KP4}
	KeyPad5        = Key{Label: "5", Width: U1, Finger: Middle, EvCode: KEY_KP5}
	KeyPad6        = Key{Label: "6", Width: U1, Finger: Ring, EvCode: KEY_KP6}
	KeyPad7        = Key{Label: "7", Width: U1, Finger: Index, EvCode: KEY_KP7}
	KeyPad8        = Key{Label: "8", Width: U1, Finger: Middle, EvCode: KEY_KP8}
	KeyPad9        = Key{Label: "9", Width: U1, Finger: Ring, EvCode: KEY_KP9}
	KeyPadPlus     = Key{Label: "", Width: U1, Finger: Pinky, Gap: true, DivLabel: "+", EvCode: KEY_KPPLUS}
	KeyPadMinus    = Key{Label: "-", Width: U1, Finger: Pinky, EvCode: KEY_KPMINUS}
	KeyPadAsterisk = Key{Label: "*", Width: U1, Finger: Ring, EvCode: KEY_KPASTERISK}
	KeyPadSlash    = Key{Label: "/", Width: U1, Finger: Middle, EvCode: KEY_KPSLASH}
	KeyPadDot      = Key{Label: ".", Width: U1, Finger: Ring, EvCode: KEY_KPDOT}
	KeyPadEnter    = Key{Label: "", Width: U1, Finger: Pinky, Gap: true, DivLabel: "↵", EvCode: KEY_KPENTER}

	// MISC
	KeyBlank         = Key{Label: "", Width: U0p75, Finger: Any, EvCode: KEY_RESERVED}
	KeyPrintScreen   = Key{Label: "Prt", Width: U1, Finger: Index, EvCode: KEY_SYSRQ}
	KeyLightsToggle  = Key{Label: "☼", Width: U1, Finger: Ring, EvCode: KEY_LIGHTS_TOGGLE}
	KeyEnterISO      = Key{Label: "Ent", Width: U1p75, Finger: Pinky, EvCode: KEY_ENTER, Gap: true}
	KeyEnterISOBlank = Key{Label: "", Width: U1p5, Finger: KeyEnterISO.Finger, EvCode: KeyEnterISO.EvCode}
	KeyLeftShiftISO  = Key{Label: "Shft", Width: U1p5, Finger: KeyLeftShift.Finger, EvCode: KeyLeftShift.EvCode}
	KeyRightShiftISO = Key{Label: "Shft", Width: U2, Finger: KeyRightShift.Finger, EvCode: KeyRightShift.EvCode}
	KeyCedilla       = Key{Label: "Ç", Width: U1, Finger: Pinky, EvCode: KEY_RESERVED}
	KeyTilde         = Key{Label: "~", Width: U1, Finger: Pinky, EvCode: KEY_RESERVED}
	KeyColon         = Key{Label: ":", Width: U1, Finger: Pinky, EvCode: KEY_102ND}

	// JIS-SPECIFIC KEYS
	KeyYen      = Key{Label: "¥", Width: U1, Finger: Pinky, EvCode: KEY_YEN}
	KeyMuhenkan = Key{Label: "無", Width: U1, Finger: Thumb, EvCode: KEY_MUHENKAN}
	KeyHenkan   = Key{Label: "変", Width: U1, Finger: Thumb, EvCode: KEY_HENKAN}
	KeyKana     = Key{Label: "仮", Width: U1, Finger: Thumb, EvCode: KEY_KATAKANAHIRAGANA}

	// KS-SPECIFIC KEYS
	KeyWon     = Key{Label: "₩", Width: U1, Finger: Pinky, EvCode: KEY_RESERVED}
	KeyHanja   = Key{Label: "한자", Width: U1p5, Finger: Thumb, EvCode: KEY_HANJA}
	KeyHangeul = Key{Label: "한영", Width: U1p5, Finger: Thumb, EvCode: KEY_HANGEUL}
)

var EvCodeLabel = buildEvCodeLabelMap()

func buildEvCodeLabelMap() map[uint16]string {
	m := make(map[uint16]string, 70)
	keys := []Key{
		KeyA, KeyB, KeyC, KeyD, KeyE, KeyF, KeyG, KeyH, KeyI, KeyJ,
		KeyK, KeyL, KeyM, KeyN, KeyO, KeyP, KeyQ, KeyR, KeyS, KeyT,
		KeyU, KeyV, KeyW, KeyX, KeyY, KeyZ,
		Key0, Key1, Key2, Key3, Key4, Key5, Key6, Key7, Key8, Key9,
		KeyEsc,
		KeyLeftShiftISO, KeyLeftShift,
		KeyRightShiftISO, KeyRightShift,
		KeyLeftCtrl, KeyRightCtrl,
		KeyLeftMeta, KeyRightMeta,
		KeyLeftAlt, KeyRightAlt,
		KeyCapsLock, KeyNumLock, KeyScrollLock,
		KeyFn,
		KeyF1, KeyF2, KeyF3, KeyF4, KeyF5, KeyF6, KeyF7, KeyF8, KeyF9, KeyF10, KeyF11, KeyF12,
		KeyUp, KeyDown, KeyLeft, KeyRight,
		KeyPageUp, KeyPageDown, KeyHome, KeyEnd,
		KeyInsert, KeyDelete, KeyBackspace,
		KeyTab,
		KeyEnterISO, KeyEnter,
		KeySpace,
		KeyGrave, KeyMinus, KeyEqual,
		KeyLeftBrace, KeyRightBrace, KeyBackslash,
		KeySemicolon, KeyApostrophe,
		KeyComma, KeyDot, KeySlash,
		KeyPound, KeyAcute,
		KeyPad0, KeyPad1, KeyPad2, KeyPad3, KeyPad4,
		KeyPad5, KeyPad6, KeyPad7, KeyPad8, KeyPad9,
		KeyPadPlus, KeyPadMinus, KeyPadAsterisk, KeyPadSlash, KeyPadDot, KeyPadEnter,
		KeyPrintScreen, KeyLightsToggle,
		KeyYen, KeyMuhenkan, KeyHenkan, KeyKana,
		KeyHanja, KeyHangeul,
		KeyCedilla, KeyTilde, KeyColon, KeyBlank, KeyWon,
	}
	for _, k := range keys {
		if k.Label != "" {
			m[k.EvCode] = k.Label
		}
	}
	return m
}
