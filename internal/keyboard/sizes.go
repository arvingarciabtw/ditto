package keyboard

type Finger int

const (
	Pinky Finger = iota
	Ring
	Middle
	Index
	Thumb
	Any
)

type key struct {
	label     string
	width     int
	finger    Finger
	gap       bool
	rightless bool
	leftless  bool
	divLabel  string
	evCode    uint16
}

var sizes = map[int][][]key{
	60:  size60ANSI,
	65:  size65ANSI,
	75:  size75ANSI,
	80:  size80ANSI,
	96:  size96ANSI,
	100: size100ANSI,
}

var sizesISO = map[int][][]key{
	60:  size60ISO,
	65:  size65ISO,
	75:  size75ISO,
	80:  size80ISO,
	96:  size96ISO,
	100: size100ISO,
}

var size60ANSI = [][]key{
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		{label: keyBackspace.label, width: u2_75, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
	},
	{
		{label: "Tab↹", width: u2, finger: keyTab.finger, evCode: keyTab.evCode},
		keyQ, keyW, keyE, keyR,
		keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace, keyBackslash,
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe,
		{label: "Enter", width: keyEnter.width, finger: keyEnter.finger, evCode: keyEnter.evCode},
	},
	{
		{label: keyLeftShift.label, width: u3_50, finger: keyLeftShift.finger, evCode: keyLeftShift.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash, keyRightShift,
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		{label: keySpace.label, width: u7_50, finger: keySpace.finger, evCode: keySpace.evCode},
		keyRightAlt, keyRightMeta, keyFn, keyRightCtrl,
	},
}

var size65ANSI = [][]key{
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		{label: keyBackspace.label, width: u2_75, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
		{label: keyLightsToggle.label, width: keyLightsToggle.width, finger: Any, evCode: keyLightsToggle.evCode},
	},
	{
		{label: "Tab↹", width: u2, finger: keyTab.finger, evCode: keyTab.evCode},
		keyQ, keyW, keyE, keyR,
		keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace, keyBackslash,
		{label: keyHome.label, width: keyHome.width, finger: Any, evCode: keyHome.evCode},
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe,
		{label: "Enter", width: keyEnter.width, finger: keyEnter.finger, evCode: keyEnter.evCode},
		{label: keyPageUp.label, width: keyPageUp.width, finger: Any, evCode: keyPageUp.evCode},
	},
	{
		{label: keyLeftShift.label, width: u2_75, finger: keyLeftShift.finger, evCode: keyLeftShift.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash,
		{label: keyRightShift.label, width: u2_50, finger: keyRightShift.finger, evCode: keyRightShift.evCode},
		keyUp,
		{label: keyPageDown.label, width: keyPageDown.width, finger: Any, evCode: keyPageDown.evCode},
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		{label: keySpace.label, width: u6, finger: keySpace.finger, evCode: keySpace.evCode},
		keyRightAlt, keyFn, keyRightCtrl, keyLeft, keyDown, keyRight,
	},
}

var size75ANSI = [][]key{
	{
		keyEsc, keyF1, keyF2, keyF3, keyF4, keyF5, keyF6, keyF7, keyF8, keyF9, keyF10, keyF11, keyF12,
		{label: keyPrintScreen.label, width: keyPrintScreen.width, finger: Any, evCode: keyPrintScreen.evCode},
		{label: keyDelete.label, width: keyDelete.width, finger: Any, evCode: keyDelete.evCode},
		{label: keyLightsToggle.label, width: keyLightsToggle.width, finger: Any, evCode: keyLightsToggle.evCode},
	},
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		keyBackspace,
		{label: keyPageUp.label, width: keyPageUp.width, finger: Any, evCode: keyPageUp.evCode},
	},
	{
		keyTab, keyQ, keyW, keyE, keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace,
		keyBackslash,
		{label: keyPageDown.label, width: keyPageDown.width, finger: Any, evCode: keyPageDown.evCode},
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe,
		{label: keyEnter.label, width: u2_75, finger: keyEnter.finger, evCode: keyEnter.evCode},
		{label: keyHome.label, width: keyHome.width, finger: Any, evCode: keyHome.evCode},
	},
	{
		keyLeftShift, keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash,
		{label: keyRightShift.label, width: u2_50, finger: keyRightShift.finger, evCode: keyRightShift.evCode},
		keyUp,
		{label: keyEnd.label, width: keyEnd.width, finger: Any, evCode: keyEnd.evCode},
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt, keySpace, keyRightAlt, keyFn, keyRightCtrl, keyLeft,
		keyDown, keyRight,
	},
}

var size80ANSI = [][]key{
	{
		keyEsc,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyF1, keyF2, keyF3, keyF4, keyBlank, keyF5, keyF6, keyF7, keyF8, keyBlank, keyF9, keyF10,
		keyF11, keyF12,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPrintScreen, keyScrollLock, keyLightsToggle,
	},
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		{label: keyBackspace.label, width: u3, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyInsert, keyHome, keyPageUp,
	},
	{
		{label: "Tab", width: u2_50, finger: keyTab.finger, evCode: keyTab.evCode},
		keyQ, keyW, keyE,
		keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace, keyBackslash,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyDelete, keyEnd, keyPageDown,
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe,
		{label: keyEnter.label, width: u3_50, finger: keyEnter.finger, evCode: keyEnter.evCode},
		{label: keyBlank.label, width: u2, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: false, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: false, leftless: false},
	},
	{
		{label: keyLeftShift.label, width: u3_75, finger: keyLeftShift.finger, evCode: keyLeftShift.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash, keyRightShift,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyUp,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		{label: keySpace.label, width: u7_75, finger: keySpace.finger, evCode: keySpace.evCode},
		keyRightAlt, keyRightMeta, keyFn, keyRightCtrl,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyLeft, keyDown, keyRight,
	},
}

var size96ANSI = [][]key{
	{
		keyEsc, keyF1, keyF2, keyF3, keyF4, keyF5, keyF6, keyF7, keyF8, keyF9, keyF10, keyF11, keyF12,
		{label: keyDelete.label, width: keyDelete.width, finger: Any, evCode: keyDelete.evCode},
		{label: keyHome.label, width: keyHome.width, finger: Any, evCode: keyHome.evCode},
		{label: keyEnd.label, width: keyEnd.width, finger: Any, evCode: keyEnd.evCode},
		{label: keyPageUp.label, width: keyPageUp.width, finger: Any, evCode: keyPageUp.evCode},
		{label: keyPageDown.label, width: keyPageDown.width, finger: Any, evCode: keyPageDown.evCode},
		{label: keyLightsToggle.label, width: keyLightsToggle.width, finger: Any, evCode: keyLightsToggle.evCode},
	},
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		keyBackspace, keyNumLock, keyPadSlash, keyPadAsterisk, keyPadMinus,
	},
	{
		keyTab, keyQ, keyW, keyE, keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace,
		keyBackslash, keyPad7, keyPad8, keyPad9, keyPadPlus,
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe,
		{label: keyEnter.label, width: u2_75, finger: keyEnter.finger, evCode: keyEnter.evCode},
		keyPad4,
		keyPad5, keyPad6,
		{label: keyPadPlus.label, width: keyPadPlus.width, finger: keyPadPlus.finger, evCode: keyPadPlus.evCode},
	},
	{
		keyLeftShift, keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash,
		{label: keyRightShift.label, width: u2_50, finger: keyRightShift.finger, evCode: keyRightShift.evCode},
		keyUp, keyPad1, keyPad2, keyPad3, keyPadEnter,
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt, keySpace, keyRightAlt, keyFn, keyRightCtrl, keyLeft,
		keyDown, keyRight,
		{label: keyPad0.label, width: u1, finger: Middle, evCode: keyPad0.evCode},
		keyPadDot,
		{label: keyPadEnter.label, width: keyPadEnter.width, finger: keyPadEnter.finger, evCode: keyPadEnter.evCode},
	},
}

var size100ANSI = [][]key{
	{
		keyEsc,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyF1, keyF2, keyF3, keyF4, keyBlank, keyF5, keyF6, keyF7, keyF8, keyBlank, keyF9, keyF10,
		keyF11, keyF12,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPrintScreen, keyScrollLock, keyLightsToggle,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
	},
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		{label: keyBackspace.label, width: u3, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyInsert, keyHome, keyPageUp,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyNumLock, keyPadSlash, keyPadAsterisk, keyPadMinus,
	},
	{
		{label: "Tab", width: u2_50, finger: keyTab.finger, evCode: keyTab.evCode},
		keyQ, keyW, keyE,
		keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace, keyBackslash,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyDelete, keyEnd, keyPageDown,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPad7, keyPad8, keyPad9, keyPadPlus,
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe,
		{label: keyEnter.label, width: u3_50, finger: keyEnter.finger, evCode: keyEnter.evCode},
		{label: keyBlank.label, width: u2, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: false, rightless: true, leftless: true},
		{label: keyBlank.label, width: u2, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: false, leftless: false},
		keyPad4, keyPad5, keyPad6,
		{label: keyPadPlus.label, width: keyPadPlus.width, finger: keyPadPlus.finger, evCode: keyPadPlus.evCode},
	},
	{
		{label: keyLeftShift.label, width: u3_75, finger: keyLeftShift.finger, evCode: keyLeftShift.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash, keyRightShift,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyUp,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPad1, keyPad2, keyPad3, keyPadEnter,
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		{label: keySpace.label, width: u7_75, finger: keySpace.finger, evCode: keySpace.evCode},
		keyRightAlt, keyRightMeta, keyFn, keyRightCtrl,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyLeft, keyDown, keyRight,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPad0, keyPadDot,
		{label: keyPadEnter.label, width: keyPadEnter.width, finger: keyPadEnter.finger, evCode: keyPadEnter.evCode},
	},
}

var size60ISO = [][]key{
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		{label: keyBackspace.label, width: u2_75, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
	},
	{
		keyTab, keyQ, keyW, keyE, keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace,
		{label: "Ent↵", width: u2, finger: keyEnterISO.finger, evCode: keyEnterISO.evCode, gap: keyEnterISO.gap},
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe,
		keyPound,
		{label: keyEnterISOBlank.label, width: u1_75, finger: keyEnterISOBlank.finger, evCode: keyEnterISOBlank.evCode},
	},
	{
		{label: keyLeftShiftISO.label, width: u2, finger: keyLeftShiftISO.finger, evCode: keyLeftShiftISO.evCode},
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash, keyRightShift,
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		{label: keySpace.label, width: u7_50, finger: keySpace.finger, evCode: keySpace.evCode},
		keyRightAlt, keyRightMeta, keyFn, keyRightCtrl,
	},
}

var size65ISO = [][]key{
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		{label: keyBackspace.label, width: u2_75, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
		{label: keyLightsToggle.label, width: keyLightsToggle.width, finger: Any, evCode: keyLightsToggle.evCode},
	},
	{
		keyTab, keyQ, keyW, keyE, keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace,
		{label: "Ent↵", width: u2, finger: keyEnterISO.finger, evCode: keyEnterISO.evCode, gap: keyEnterISO.gap},
		{label: keyHome.label, width: keyHome.width, finger: Any, evCode: keyHome.evCode},
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe,
		keyPound,
		{label: keyEnterISOBlank.label, width: u1_75, finger: keyEnterISOBlank.finger, evCode: keyEnterISOBlank.evCode},
		{label: keyPageUp.label, width: keyPageUp.width, finger: Any, evCode: keyPageUp.evCode},
	},
	{
		keyLeftShiftISO,
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash,
		{label: keyRightShift.label, width: u2_50, finger: keyRightShift.finger, evCode: keyRightShift.evCode},
		keyUp,
		{label: keyPageDown.label, width: keyPageDown.width, finger: Any, evCode: keyPageDown.evCode},
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		{label: keySpace.label, width: u6, finger: keySpace.finger, evCode: keySpace.evCode},
		keyRightAlt, keyFn, keyRightCtrl, keyLeft, keyDown, keyRight,
	},
}

var size75ISO = [][]key{
	{
		keyEsc, keyF1, keyF2, keyF3, keyF4, keyF5, keyF6, keyF7, keyF8, keyF9, keyF10, keyF11, keyF12,
		{label: keyPrintScreen.label, width: keyPrintScreen.width, finger: Any, evCode: keyPrintScreen.evCode},
		{label: keyDelete.label, width: keyDelete.width, finger: Any, evCode: keyDelete.evCode},
		{label: keyLightsToggle.label, width: keyLightsToggle.width, finger: Any, evCode: keyLightsToggle.evCode},
	},
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		keyBackspace,
		{label: keyPageUp.label, width: keyPageUp.width, finger: Any, evCode: keyPageUp.evCode},
	},
	{
		keyTab, keyQ, keyW, keyE, keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace,
		keyEnterISO,
		{label: keyPageDown.label, width: keyPageDown.width, finger: Any, evCode: keyPageDown.evCode},
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe,
		keyPound, keyEnterISOBlank,
		{label: keyHome.label, width: keyHome.width, finger: Any, evCode: keyHome.evCode},
	},
	{
		keyLeftShiftISO,
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash, keyRightShiftISO, keyUp,
		{label: keyEnd.label, width: keyEnd.width, finger: Any, evCode: keyEnd.evCode},
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt, keySpace, keyRightAlt, keyFn, keyRightCtrl, keyLeft,
		keyDown, keyRight,
	},
}

var size80ISO = [][]key{
	{
		keyEsc,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyF1, keyF2, keyF3, keyF4, keyBlank, keyF5, keyF6, keyF7, keyF8, keyBlank, keyF9, keyF10,
		keyF11, keyF12,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPrintScreen, keyScrollLock, keyLightsToggle,
	},
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		{label: keyBackspace.label, width: u3, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyInsert, keyHome, keyPageUp,
	},
	{
		{label: "Tab", width: u2_50, finger: keyTab.finger, evCode: keyTab.evCode},
		keyQ, keyW, keyE,
		keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace, keyEnterISO,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyDelete, keyEnd, keyPageDown,
	},
	{
		{label: keyCapsLock.label, width: u2_75, finger: keyCapsLock.finger, evCode: keyCapsLock.evCode},
		keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe, keyPound,
		keyEnterISOBlank,
		{label: keyBlank.label, width: u2, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: false, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: false, leftless: false},
	},
	{
		keyLeftShift,
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash, keyRightShift,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyUp,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		{label: keySpace.label, width: u7_75, finger: keySpace.finger, evCode: keySpace.evCode},
		keyRightAlt, keyRightMeta, keyFn, keyRightCtrl,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyLeft, keyDown, keyRight,
	},
}

var size96ISO = [][]key{
	{
		keyEsc, keyF1, keyF2, keyF3, keyF4, keyF5, keyF6, keyF7, keyF8, keyF9, keyF10, keyF11, keyF12,
		{label: keyDelete.label, width: keyDelete.width, finger: Any, evCode: keyDelete.evCode},
		{label: keyHome.label, width: keyHome.width, finger: Any, evCode: keyHome.evCode},
		{label: keyEnd.label, width: keyEnd.width, finger: Any, evCode: keyEnd.evCode},
		{label: keyPageUp.label, width: keyPageUp.width, finger: Any, evCode: keyPageUp.evCode},
		{label: keyPageDown.label, width: keyPageDown.width, finger: Any, evCode: keyPageDown.evCode},
		{label: keyLightsToggle.label, width: keyLightsToggle.width, finger: Any, evCode: keyLightsToggle.evCode},
	},
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		keyBackspace, keyNumLock, keyPadSlash, keyPadAsterisk, keyPadMinus,
	},
	{
		keyTab, keyQ, keyW, keyE, keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace,
		keyEnterISO, keyPad7, keyPad8, keyPad9, keyPadPlus,
	},
	{
		keyCapsLock, keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe,
		keyPound, keyEnterISOBlank, keyPad4, keyPad5, keyPad6,
		{label: keyPadPlus.label, width: keyPadPlus.width, finger: keyPadPlus.finger, evCode: keyPadPlus.evCode},
	},
	{
		keyLeftShiftISO,
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash, keyRightShiftISO, keyUp,
		keyPad1, keyPad2, keyPad3, keyPadEnter,
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt, keySpace, keyRightAlt, keyFn, keyRightCtrl, keyLeft,
		keyDown, keyRight,
		{label: keyPad0.label, width: u1, finger: Middle, evCode: keyPad0.evCode},
		keyPadDot,
		{label: keyPadEnter.label, width: keyPadEnter.width, finger: keyPadEnter.finger, evCode: keyPadEnter.evCode},
	},
}

var size100ISO = [][]key{
	{
		keyEsc,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyF1, keyF2, keyF3, keyF4, keyBlank, keyF5, keyF6, keyF7, keyF8, keyBlank, keyF9, keyF10,
		keyF11, keyF12,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPrintScreen, keyScrollLock, keyLightsToggle,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
	},
	{
		keyGrave, key1, key2, key3, key4, key5, key6, key7, key8, key9, key0, keyMinus, keyEqual,
		{label: keyBackspace.label, width: u3, finger: keyBackspace.finger, evCode: keyBackspace.evCode},
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyInsert, keyHome, keyPageUp,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyNumLock, keyPadSlash, keyPadAsterisk, keyPadMinus,
	},
	{
		{label: "Tab", width: u2_50, finger: keyTab.finger, evCode: keyTab.evCode},
		keyQ, keyW, keyE,
		keyR, keyT, keyY, keyU, keyI, keyO, keyP, keyLeftBrace, keyRightBrace, keyEnterISO,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyDelete, keyEnd, keyPageDown,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPad7, keyPad8, keyPad9, keyPadPlus,
	},
	{
		{label: keyCapsLock.label, width: u2_75, finger: keyCapsLock.finger, evCode: keyCapsLock.evCode},
		keyA, keyS, keyD, keyF, keyG, keyH, keyJ, keyK, keyL, keySemicolon, keyApostrophe, keyPound,
		keyEnterISOBlank,
		{label: keyBlank.label, width: u2, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: false, rightless: true, leftless: true},
		{label: keyBlank.label, width: u2, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: false, leftless: false},
		keyPad4, keyPad5, keyPad6,
		{label: keyPadPlus.label, width: keyPadPlus.width, finger: keyPadPlus.finger, evCode: keyPadPlus.evCode},
	},
	{
		keyLeftShift,
		{label: keyBackslash.label, width: u1, finger: keyBackslash.finger, evCode: keyBackslash.evCode},
		keyZ, keyX, keyC, keyV, keyB, keyN, keyM, keyComma, keyDot, keySlash, keyRightShift,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true, rightless: true, leftless: false},
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyUp,
		{label: keyBlank.label, width: u1, finger: keyBlank.finger, evCode: keyBlank.evCode, rightless: true, leftless: true},
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPad1, keyPad2, keyPad3, keyPadEnter,
	},
	{
		keyLeftCtrl, keyLeftMeta, keyLeftAlt,
		{label: keySpace.label, width: u7_75, finger: keySpace.finger, evCode: keySpace.evCode},
		keyRightAlt, keyRightMeta, keyFn, keyRightCtrl,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode},
		keyLeft, keyDown, keyRight,
		{label: keyBlank.label, width: keyBlank.width, finger: keyBlank.finger, evCode: keyBlank.evCode, gap: true},
		keyPad0, keyPadDot,
		{label: keyPadEnter.label, width: keyPadEnter.width, finger: keyPadEnter.finger, evCode: keyPadEnter.evCode},
	},
}
