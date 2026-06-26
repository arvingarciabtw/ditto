package keyboard

import evdev "github.com/gvalkov/golang-evdev"

/*
* These are constants for keycap units. 1 unit is equivalent to a width of 3.
* Essentially, square keys are 1 unit of measurement. It scales by 3 to keep
* the ASCII proportional. Since this is the scaling, decimal variants aren't
* exact. I rounded them instead.
*
* For more info on units:
* https://www.minimaldesksetups.com/mechanical-keyboard-keycap-guide/
 */
const (
	u1    = 3
	u1_50 = 4
	u1_75 = 5
	u2    = 6
	u2_50 = 7
	u2_75 = 8
	u3    = 9
	u3_50 = 10
	u3_75 = 11
	u5_75 = 17
	u6    = 18
	u7_50 = 22
	u7_75 = 23
)

var (
	// ALPHANUMERIC KEYS
	keyA = key{label: "A", width: u1, finger: Pinky, evCode: evdev.KEY_A}
	keyB = key{label: "B", width: u1, finger: Index, evCode: evdev.KEY_B}
	keyC = key{label: "C", width: u1, finger: Middle, evCode: evdev.KEY_C}
	keyD = key{label: "D", width: u1, finger: Middle, evCode: evdev.KEY_D}
	keyE = key{label: "E", width: u1, finger: Middle, evCode: evdev.KEY_E}
	keyF = key{label: "F", width: u1, finger: Index, evCode: evdev.KEY_F}
	keyG = key{label: "G", width: u1, finger: Index, evCode: evdev.KEY_G}
	keyH = key{label: "H", width: u1, finger: Index, evCode: evdev.KEY_H}
	keyI = key{label: "I", width: u1, finger: Middle, evCode: evdev.KEY_I}
	keyJ = key{label: "J", width: u1, finger: Index, evCode: evdev.KEY_J}
	keyK = key{label: "K", width: u1, finger: Middle, evCode: evdev.KEY_K}
	keyL = key{label: "L", width: u1, finger: Ring, evCode: evdev.KEY_L}
	keyM = key{label: "M", width: u1, finger: Index, evCode: evdev.KEY_M}
	keyN = key{label: "N", width: u1, finger: Index, evCode: evdev.KEY_N}
	keyO = key{label: "O", width: u1, finger: Ring, evCode: evdev.KEY_O}
	keyP = key{label: "P", width: u1, finger: Pinky, evCode: evdev.KEY_P}
	keyQ = key{label: "Q", width: u1, finger: Pinky, evCode: evdev.KEY_Q}
	keyR = key{label: "R", width: u1, finger: Index, evCode: evdev.KEY_R}
	keyS = key{label: "S", width: u1, finger: Ring, evCode: evdev.KEY_S}
	keyT = key{label: "T", width: u1, finger: Index, evCode: evdev.KEY_T}
	keyU = key{label: "U", width: u1, finger: Index, evCode: evdev.KEY_U}
	keyV = key{label: "V", width: u1, finger: Index, evCode: evdev.KEY_V}
	keyW = key{label: "W", width: u1, finger: Ring, evCode: evdev.KEY_W}
	keyX = key{label: "X", width: u1, finger: Ring, evCode: evdev.KEY_X}
	keyY = key{label: "Y", width: u1, finger: Index, evCode: evdev.KEY_Y}
	keyZ = key{label: "Z", width: u1, finger: Pinky, evCode: evdev.KEY_Z}
	key0 = key{label: "0", width: u1, finger: Pinky, evCode: evdev.KEY_0}
	key1 = key{label: "1", width: u1, finger: Pinky, evCode: evdev.KEY_1}
	key2 = key{label: "2", width: u1, finger: Ring, evCode: evdev.KEY_2}
	key3 = key{label: "3", width: u1, finger: Middle, evCode: evdev.KEY_3}
	key4 = key{label: "4", width: u1, finger: Index, evCode: evdev.KEY_4}
	key5 = key{label: "5", width: u1, finger: Index, evCode: evdev.KEY_5}
	key6 = key{label: "6", width: u1, finger: Index, evCode: evdev.KEY_6}
	key7 = key{label: "7", width: u1, finger: Index, evCode: evdev.KEY_7}
	key8 = key{label: "8", width: u1, finger: Middle, evCode: evdev.KEY_8}
	key9 = key{label: "9", width: u1, finger: Ring, evCode: evdev.KEY_9}

	// CONTROL KEYS
	keyEsc        = key{label: "Esc", width: u1, finger: Pinky, evCode: evdev.KEY_ESC}
	keyLeftShift  = key{label: "Shift", width: u2_50, finger: Pinky, evCode: evdev.KEY_LEFTSHIFT}
	keyRightShift = key{label: "Shift", width: u3, finger: Pinky, evCode: evdev.KEY_RIGHTSHIFT}
	keyLeftCtrl   = key{label: "Ctrl", width: u2, finger: Pinky, evCode: evdev.KEY_LEFTCTRL}
	keyRightCtrl  = key{label: "Ctrl", width: u2, finger: Pinky, evCode: evdev.KEY_RIGHTCTRL}
	keyLeftMeta   = key{label: "⌘", width: u1, finger: Ring, evCode: evdev.KEY_LEFTMETA}
	keyRightMeta  = key{label: "⌘", width: u1, finger: Ring, evCode: evdev.KEY_RIGHTMETA}
	keyLeftAlt    = key{label: "Alt", width: u1_75, finger: Thumb, evCode: evdev.KEY_LEFTALT}
	keyRightAlt   = key{label: "Alt", width: u1_75, finger: Thumb, evCode: evdev.KEY_RIGHTALT}
	keyCapsLock   = key{label: "Caps", width: u2, finger: Pinky, evCode: evdev.KEY_CAPSLOCK}
	keyNumLock    = key{label: "Nlk", width: u1, finger: Index, evCode: evdev.KEY_NUMLOCK}
	keyScrollLock = key{label: "⚲", width: u1, finger: Middle, evCode: evdev.KEY_SCROLLLOCK}

	// FUNCTION KEYS
	keyFn  = key{label: "Fn", width: u1, finger: Ring, evCode: evdev.KEY_FN}
	keyF1  = key{label: "F1", width: u1, finger: Pinky, evCode: evdev.KEY_F1}
	keyF2  = key{label: "F2", width: u1, finger: Ring, evCode: evdev.KEY_F2}
	keyF3  = key{label: "F3", width: u1, finger: Middle, evCode: evdev.KEY_F3}
	keyF4  = key{label: "F4", width: u1, finger: Index, evCode: evdev.KEY_F4}
	keyF5  = key{label: "F5", width: u1, finger: Index, evCode: evdev.KEY_F5}
	keyF6  = key{label: "F6", width: u1, finger: Index, evCode: evdev.KEY_F6}
	keyF7  = key{label: "F7", width: u1, finger: Index, evCode: evdev.KEY_F7}
	keyF8  = key{label: "F8", width: u1, finger: Middle, evCode: evdev.KEY_F8}
	keyF9  = key{label: "F9", width: u1, finger: Ring, evCode: evdev.KEY_F9}
	keyF10 = key{label: "F10", width: u1, finger: Pinky, evCode: evdev.KEY_F10}
	keyF11 = key{label: "F11", width: u1, finger: Pinky, evCode: evdev.KEY_F11}
	keyF12 = key{label: "F12", width: u1, finger: Pinky, evCode: evdev.KEY_F12}

	// NAVIGATION AND EDITING KEYS
	keyUp        = key{label: "↑", width: u1, finger: Middle, evCode: evdev.KEY_UP}
	keyDown      = key{label: "↓", width: u1, finger: Middle, evCode: evdev.KEY_DOWN}
	keyLeft      = key{label: "←", width: u1, finger: Index, evCode: evdev.KEY_LEFT}
	keyRight     = key{label: "→", width: u1, finger: Ring, evCode: evdev.KEY_RIGHT}
	keyPageUp    = key{label: "⇡", width: u1, finger: Ring, evCode: evdev.KEY_PAGEUP}
	keyPageDown  = key{label: "⇣", width: u1, finger: Ring, evCode: evdev.KEY_PAGEDOWN}
	keyHome      = key{label: "⌂", width: u1, finger: Middle, evCode: evdev.KEY_HOME}
	keyEnd       = key{label: "⌿", width: u1, finger: Middle, evCode: evdev.KEY_END}
	keyInsert    = key{label: "Ins", width: u1, finger: Index, evCode: evdev.KEY_INSERT}
	keyDelete    = key{label: "Del", width: u1, finger: Index, evCode: evdev.KEY_DELETE}
	keyBackspace = key{label: "<--", width: u2_50, finger: Pinky, evCode: evdev.KEY_BACKSPACE}

	// PUNCTUATION AND SPECIAL KEYS
	keyTab        = key{label: "Tab", width: u1_75, finger: Pinky, evCode: evdev.KEY_TAB}
	keyEnter      = key{label: "Enter↵", width: u3, finger: Pinky, evCode: evdev.KEY_ENTER}
	keySpace      = key{label: "Space", width: u5_75, finger: Thumb, evCode: evdev.KEY_SPACE}
	keyGrave      = key{label: "`", width: u1, finger: Pinky, evCode: evdev.KEY_GRAVE}
	keyMinus      = key{label: "-", width: u1, finger: Pinky, evCode: evdev.KEY_MINUS}
	keyEqual      = key{label: "=", width: u1, finger: Pinky, evCode: evdev.KEY_EQUAL}
	keyLeftBrace  = key{label: "[", width: u1, finger: Pinky, evCode: evdev.KEY_LEFTBRACE}
	keyRightBrace = key{label: "]", width: u1, finger: Pinky, evCode: evdev.KEY_RIGHTBRACE}
	keyBackslash  = key{label: "\\", width: u1_75, finger: Pinky, evCode: evdev.KEY_BACKSLASH}
	keySemicolon  = key{label: ";", width: u1, finger: Pinky, evCode: evdev.KEY_SEMICOLON}
	keyApostrophe = key{label: "'", width: u1, finger: Pinky, evCode: evdev.KEY_APOSTROPHE}
	keyComma      = key{label: ",", width: u1, finger: Middle, evCode: evdev.KEY_COMMA}
	keyDot        = key{label: ".", width: u1, finger: Ring, evCode: evdev.KEY_DOT}
	keySlash      = key{label: "/", width: u1, finger: Pinky, evCode: evdev.KEY_SLASH}
	keyPound      = key{label: "#", width: u1, finger: Pinky, evCode: evdev.KEY_NUMERIC_POUND}

	// SYSTEM COMMAND KEYS
	keyPrintScreen = key{label: "Prt", width: u1, finger: Index, evCode: evdev.KEY_SYSRQ}

	// HARDWARE KEYS
	keyLightsToggle = key{label: "☼", width: u1, finger: Ring, evCode: evdev.KEY_LIGHTS_TOGGLE}

	// NUMERIC KEYPAD
	keyPad0        = key{label: "0", width: u2_50, finger: Thumb, evCode: evdev.KEY_KP0}
	keyPad1        = key{label: "1", width: u1, finger: Index, evCode: evdev.KEY_KP1}
	keyPad2        = key{label: "2", width: u1, finger: Middle, evCode: evdev.KEY_KP2}
	keyPad3        = key{label: "3", width: u1, finger: Ring, evCode: evdev.KEY_KP3}
	keyPad4        = key{label: "4", width: u1, finger: Index, evCode: evdev.KEY_KP4}
	keyPad5        = key{label: "5", width: u1, finger: Middle, evCode: evdev.KEY_KP5}
	keyPad6        = key{label: "6", width: u1, finger: Ring, evCode: evdev.KEY_KP6}
	keyPad7        = key{label: "7", width: u1, finger: Index, evCode: evdev.KEY_KP7}
	keyPad8        = key{label: "8", width: u1, finger: Middle, evCode: evdev.KEY_KP8}
	keyPad9        = key{label: "9", width: u1, finger: Ring, evCode: evdev.KEY_KP9}
	keyPadPlus     = key{label: "", width: u1, finger: Pinky, gap: true, divLabel: "+", evCode: evdev.KEY_KPPLUS}
	keyPadMinus    = key{label: "-", width: u1, finger: Pinky, evCode: evdev.KEY_KPMINUS}
	keyPadAsterisk = key{label: "*", width: u1, finger: Ring, evCode: evdev.KEY_KPASTERISK}
	keyPadSlash    = key{label: "/", width: u1, finger: Middle, evCode: evdev.KEY_KPSLASH}
	keyPadDot      = key{label: ".", width: u1, finger: Ring, evCode: evdev.KEY_KPDOT}
	keyPadEnter    = key{label: "", width: u1, finger: Pinky, gap: true, divLabel: "↵", evCode: evdev.KEY_KPENTER}

	// MISC
	keyBlank         = key{label: "", width: 2, finger: Any, evCode: evdev.KEY_RESERVED}
	keyEnterISO      = key{label: "Ent", width: u1_75, finger: Pinky, evCode: evdev.KEY_ENTER, gap: true}
	keyEnterISOBlank = key{label: "", width: u1_50, finger: keyEnterISO.finger, evCode: keyEnterISO.evCode}
	keyLeftShiftISO  = key{label: "Shft", width: u1_50, finger: keyLeftShift.finger, evCode: keyLeftShift.evCode}
	keyRightShiftISO = key{label: "Shft", width: u2, finger: keyRightShift.finger, evCode: keyRightShift.evCode}
)
