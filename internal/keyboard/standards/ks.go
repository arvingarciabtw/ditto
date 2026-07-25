package standards

import "github.com/arvingarciabtw/ditto/internal/keyboard/base"

var HangeulLayout = map[string]string{
	"Q": "ㅂ", "W": "ㅈ", "E": "ㄷ", "R": "ㄱ", "T": "ㅅ",
	"Y": "ㅛ", "U": "ㅠ", "I": "ㅑ", "O": "ㅓ", "P": "ㅔ",

	"A": "ㅁ", "S": "ㄴ", "D": "ㅇ", "F": "ㄹ", "G": "ㅎ",
	"H": "ㅗ", "J": "ㅓ", "K": "ㅏ", "L": "ㅣ",

	"Z": "ㅋ", "X": "ㅌ", "C": "ㅊ", "V": "ㅍ", "B": "ㅠ",
	"N": "ㅜ", "M": "ㅡ",
}

var HangeulTensed = map[string]string{
	"Q": "ㅃ", "W": "ㅉ", "E": "ㄸ", "R": "ㄲ", "T": "ㅆ",
	"O": "ㅐ", "P": "ㅒ",
}

var SizesKS = map[int][][]base.Key{
	60:  size60KS,
	65:  size65KS,
	75:  size75KS,
	80:  size80KS,
	96:  size96KS,
	100: size100KS,
}

var size60KS = [][]base.Key{
	{
		base.KeyGrave, base.Key1, base.Key2, base.Key3, base.Key4, base.Key5, base.Key6, base.Key7, base.Key8, base.Key9, base.Key0, base.KeyMinus, base.KeyEqual, base.KeyWon,
		{Label: "<-", Width: base.U1p5, Finger: base.KeyBackspace.Finger, EvCode: base.KeyBackspace.EvCode},
	},
	{
		{Label: base.KeyTab.Label, Width: base.U1p75, Finger: base.KeyTab.Finger, EvCode: base.KeyTab.EvCode},
		base.KeyQ, base.KeyW, base.KeyE, base.KeyR, base.KeyT, base.KeyY, base.KeyU, base.KeyI, base.KeyO, base.KeyP, base.KeyLeftBrace,
		{Label: base.KeyRightBrace.Label, Width: base.KeyRightBrace.Width, Finger: base.KeyRightBrace.Finger, EvCode: base.KeyRightBrace.EvCode, Leftless: true, BottomLeftlessAlt: true},
		{Label: base.KeyBlank.Label, Width: base.U2, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
	},
	{
		base.KeyCapsLock, base.KeyA, base.KeyS, base.KeyD, base.KeyF, base.KeyG, base.KeyH, base.KeyJ, base.KeyK, base.KeyL, base.KeySemicolon, base.KeyApostrophe,
		{Label: "Enter ↵", Width: base.U3, Finger: base.KeyEnter.Finger, EvCode: base.KeyEnter.EvCode, Gap: base.KeyEnter.Gap},
	},
	{
		{Label: "Shift", Width: base.U2p75, Finger: base.KeyLeftShiftISO.Finger, EvCode: base.KeyLeftShiftISO.EvCode},
		base.KeyZ, base.KeyX, base.KeyC, base.KeyV, base.KeyB, base.KeyN, base.KeyM, base.KeyComma, base.KeyDot, base.KeySlash,
		{Label: "Shift", Width: base.U3p75, Finger: base.KeyRightShiftISO.Finger, EvCode: base.KeyRightShiftISO.EvCode},
	},
	{
		base.KeyLeftCtrl, base.KeyLeftMeta, base.KeyLeftAlt, base.KeyHanja,
		{Label: base.KeySpace.Label, Width: base.U4, Finger: base.KeySpace.Finger, EvCode: base.KeySpace.EvCode},
		base.KeyHangeul, base.KeyRightAlt, base.KeyRightMeta, base.KeyFn, base.KeyRightCtrl,
	},
}

var size65KS = [][]base.Key{
	{
		base.KeyGrave, base.Key1, base.Key2, base.Key3, base.Key4, base.Key5, base.Key6, base.Key7, base.Key8, base.Key9, base.Key0, base.KeyMinus, base.KeyEqual, base.KeyWon,
		{Label: "<-", Width: base.U1p5, Finger: base.KeyBackspace.Finger, EvCode: base.KeyBackspace.EvCode},
		{Label: base.KeyLightsToggle.Label, Width: base.KeyLightsToggle.Width, Finger: base.Any, EvCode: base.KeyLightsToggle.EvCode},
	},
	{
		base.KeyTab, base.KeyQ, base.KeyW, base.KeyE, base.KeyR, base.KeyT, base.KeyY, base.KeyU, base.KeyI, base.KeyO, base.KeyP, base.KeyLeftBrace,
		{Label: base.KeyRightBrace.Label, Width: base.KeyRightBrace.Width, Finger: base.KeyRightBrace.Finger, EvCode: base.KeyRightBrace.EvCode, Leftless: true, BottomLeftlessAlt: true},
		{Label: base.KeyBlank.Label, Width: base.U2, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		{Label: base.KeyHome.Label, Width: base.KeyHome.Width, Finger: base.Any, EvCode: base.KeyHome.EvCode},
	},
	{
		base.KeyCapsLock, base.KeyA, base.KeyS, base.KeyD, base.KeyF, base.KeyG, base.KeyH, base.KeyJ, base.KeyK, base.KeyL, base.KeySemicolon, base.KeyApostrophe,
		{Label: "Enter ↵", Width: base.KeyEnter.Width, Finger: base.KeyEnter.Finger, EvCode: base.KeyEnter.EvCode},
		{Label: base.KeyPageUp.Label, Width: base.KeyPageUp.Width, Finger: base.Any, EvCode: base.KeyPageUp.EvCode},
	},
	{
		{Label: "Shift", Width: base.U2p75, Finger: base.KeyLeftShift.Finger, EvCode: base.KeyLeftShift.EvCode},
		base.KeyZ, base.KeyX, base.KeyC, base.KeyV, base.KeyB, base.KeyN, base.KeyM, base.KeyComma, base.KeyDot, base.KeySlash,
		{Label: "Shift", Width: base.U2p5, Finger: base.KeyRightShift.Finger, EvCode: base.KeyRightShift.EvCode},
		base.KeyUp,
		{Label: base.KeyPageDown.Label, Width: base.KeyPageDown.Width, Finger: base.Any, EvCode: base.KeyPageDown.EvCode},
	},
	{
		base.KeyLeftCtrl, base.KeyLeftMeta, base.KeyLeftAlt, base.KeyHanja,
		{Label: base.KeySpace.Label, Width: base.U2p75, Finger: base.KeySpace.Finger, EvCode: base.KeySpace.EvCode},
		base.KeyHangeul, base.KeyRightAlt, base.KeyRightMeta, base.KeyRightCtrl, base.KeyLeft, base.KeyDown, base.KeyRight,
	},
}

var size75KS = [][]base.Key{
	{
		base.KeyEsc, base.KeyF1, base.KeyF2, base.KeyF3, base.KeyF4, base.KeyF5, base.KeyF6, base.KeyF7, base.KeyF8, base.KeyF9, base.KeyF10, base.KeyF11, base.KeyF12,
		{Label: base.KeyPrintScreen.Label, Width: base.KeyPrintScreen.Width, Finger: base.Any, EvCode: base.KeyPrintScreen.EvCode},
		{Label: base.KeyDelete.Label, Width: base.KeyDelete.Width, Finger: base.Any, EvCode: base.KeyDelete.EvCode},
		{Label: base.KeyLightsToggle.Label, Width: base.KeyLightsToggle.Width, Finger: base.Any, EvCode: base.KeyLightsToggle.EvCode},
	},
	{
		base.KeyGrave, base.Key1, base.Key2, base.Key3, base.Key4, base.Key5, base.Key6, base.Key7, base.Key8, base.Key9, base.Key0, base.KeyMinus, base.KeyEqual, base.KeyWon,
		{Label: "<", Width: base.U1, Finger: base.KeyBackspace.Finger, EvCode: base.KeyBackspace.EvCode},
		{Label: base.KeyPageUp.Label, Width: base.KeyPageUp.Width, Finger: base.Any, EvCode: base.KeyPageUp.EvCode},
	},
	{
		base.KeyTab, base.KeyQ, base.KeyW, base.KeyE, base.KeyR, base.KeyT, base.KeyY, base.KeyU, base.KeyI, base.KeyO, base.KeyP, base.KeyLeftBrace,
		{Label: base.KeyRightBrace.Label, Width: base.KeyRightBrace.Width, Finger: base.KeyRightBrace.Finger, EvCode: base.KeyRightBrace.EvCode, Leftless: true, BottomLeftlessAlt: true},
		{Label: base.KeyBlank.Label, Width: base.U1p75, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		{Label: base.KeyPageDown.Label, Width: base.KeyPageDown.Width, Finger: base.Any, EvCode: base.KeyPageDown.EvCode},
	},
	{
		base.KeyCapsLock, base.KeyA, base.KeyS, base.KeyD, base.KeyF, base.KeyG, base.KeyH, base.KeyJ, base.KeyK, base.KeyL, base.KeySemicolon, base.KeyApostrophe,
		{Label: "Enter ↵", Width: base.U2p75, Finger: base.KeyEnter.Finger, EvCode: base.KeyEnter.EvCode},
		{Label: base.KeyHome.Label, Width: base.KeyHome.Width, Finger: base.Any, EvCode: base.KeyHome.EvCode},
	},
	{
		base.KeyLeftShift, base.KeyZ, base.KeyX, base.KeyC, base.KeyV, base.KeyB, base.KeyN, base.KeyM, base.KeyComma, base.KeyDot, base.KeySlash,
		{Label: "Shift", Width: base.U2p5, Finger: base.KeyRightShift.Finger, EvCode: base.KeyRightShift.EvCode},
		base.KeyUp,
		{Label: base.KeyEnd.Label, Width: base.KeyEnd.Width, Finger: base.Any, EvCode: base.KeyEnd.EvCode},
	},
	{
		base.KeyLeftCtrl, base.KeyLeftMeta,
		{Label: base.KeyLeftAlt.Label, Width: base.U1, Finger: base.KeyLeftAlt.Finger, EvCode: base.KeyLeftAlt.EvCode},
		base.KeyHanja,
		{Label: base.KeySpace.Label, Width: base.U3p75, Finger: base.KeySpace.Finger, EvCode: base.KeySpace.EvCode},
		base.KeyHangeul,
		{Label: base.KeyRightAlt.Label, Width: base.U1, Finger: base.KeyRightAlt.Finger, EvCode: base.KeyRightAlt.EvCode},
		base.KeyRightMeta, base.KeyRightCtrl, base.KeyLeft,
		base.KeyDown, base.KeyRight,
	},
}

var size80KS = [][]base.Key{
	{
		base.KeyEsc,
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode},
		base.KeyF1, base.KeyF2, base.KeyF3, base.KeyF4, base.KeyBlank, base.KeyF5, base.KeyF6, base.KeyF7, base.KeyF8, base.KeyBlank, base.KeyF9, base.KeyF10,
		base.KeyF11, base.KeyF12,
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		base.KeyPrintScreen, base.KeyScrollLock, base.KeyLightsToggle,
	},
	{
		base.KeyGrave, base.Key1, base.Key2, base.Key3, base.Key4, base.Key5, base.Key6, base.Key7, base.Key8, base.Key9, base.Key0, base.KeyMinus, base.KeyEqual, base.KeyWon,
		{Label: "<--", Width: base.U1p75, Finger: base.KeyBackspace.Finger, EvCode: base.KeyBackspace.EvCode},
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		base.KeyInsert, base.KeyHome, base.KeyPageUp,
	},
	{
		base.KeyTab, base.KeyQ, base.KeyW, base.KeyE, base.KeyR, base.KeyT, base.KeyY, base.KeyU, base.KeyI, base.KeyO, base.KeyP, base.KeyLeftBrace,
		{Label: base.KeyRightBrace.Label, Width: base.KeyRightBrace.Width, Finger: base.KeyRightBrace.Finger, EvCode: base.KeyRightBrace.EvCode, Leftless: true, BottomLeftlessAlt: true},
		{Label: base.KeyBlank.Label, Width: base.U2p5, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, GapAlt: true},
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Leftless: true, BottomLeftless: true},
		base.KeyDelete, base.KeyEnd, base.KeyPageDown,
	},
	{
		base.KeyCapsLock, base.KeyA, base.KeyS, base.KeyD, base.KeyF, base.KeyG, base.KeyH, base.KeyJ, base.KeyK, base.KeyL, base.KeySemicolon, base.KeyApostrophe,
		{Label: "Enter ↵", Width: base.U3p5, Finger: base.KeyEnter.Finger, EvCode: base.KeyEnter.EvCode},
		{Label: base.KeyBlank.Label, Width: base.U2, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Rightless: true, Leftless: false},
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: false, Rightless: true, Leftless: true, BottomLeftless: true},
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Rightless: false, Leftless: false},
	},
	{
		{Label: "Shift", Width: base.U3, Finger: base.KeyLeftShift.Finger, EvCode: base.KeyLeftShift.EvCode},
		base.KeyZ, base.KeyX, base.KeyC, base.KeyV, base.KeyB, base.KeyN, base.KeyM, base.KeyComma, base.KeyDot, base.KeySlash,
		{Label: "Shift", Width: base.U3p75, Finger: base.KeyRightShift.Finger, EvCode: base.KeyRightShift.EvCode},
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Rightless: true, Leftless: false},
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode},
		base.KeyUp,
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode},
	},
	{
		base.KeyLeftCtrl, base.KeyLeftMeta, base.KeyLeftAlt, base.KeyHanja,
		{Label: base.KeySpace.Label, Width: base.U4p5, Finger: base.KeySpace.Finger, EvCode: base.KeySpace.EvCode},
		base.KeyHangeul, base.KeyRightAlt, base.KeyRightMeta, base.KeyFn, base.KeyRightCtrl,
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode},
		base.KeyLeft, base.KeyDown, base.KeyRight,
	},
}

var size96KS = [][]base.Key{
	{
		base.KeyEsc, base.KeyF1, base.KeyF2, base.KeyF3, base.KeyF4, base.KeyF5, base.KeyF6, base.KeyF7, base.KeyF8, base.KeyF9, base.KeyF10, base.KeyF11, base.KeyF12,
		{Label: base.KeyDelete.Label, Width: base.KeyDelete.Width, Finger: base.Any, EvCode: base.KeyDelete.EvCode},
		{Label: base.KeyHome.Label, Width: base.KeyHome.Width, Finger: base.Any, EvCode: base.KeyHome.EvCode},
		{Label: base.KeyEnd.Label, Width: base.KeyEnd.Width, Finger: base.Any, EvCode: base.KeyEnd.EvCode},
		{Label: base.KeyPageUp.Label, Width: base.KeyPageUp.Width, Finger: base.Any, EvCode: base.KeyPageUp.EvCode},
		{Label: base.KeyPageDown.Label, Width: base.KeyPageDown.Width, Finger: base.Any, EvCode: base.KeyPageDown.EvCode},
		{Label: base.KeyLightsToggle.Label, Width: base.KeyLightsToggle.Width, Finger: base.Any, EvCode: base.KeyLightsToggle.EvCode},
	},
	{
		base.KeyGrave, base.Key1, base.Key2, base.Key3, base.Key4, base.Key5, base.Key6, base.Key7, base.Key8, base.Key9, base.Key0, base.KeyMinus, base.KeyEqual, base.KeyWon,
		{Label: "<", Width: base.U1, Finger: base.KeyBackspace.Finger, EvCode: base.KeyBackspace.EvCode},
		base.KeyNumLock, base.KeyPadSlash, base.KeyPadAsterisk, base.KeyPadMinus,
	},
	{
		base.KeyTab, base.KeyQ, base.KeyW, base.KeyE, base.KeyR, base.KeyT, base.KeyY, base.KeyU, base.KeyI, base.KeyO, base.KeyP, base.KeyLeftBrace,
		{Label: base.KeyRightBrace.Label, Width: base.KeyRightBrace.Width, Finger: base.KeyRightBrace.Finger, EvCode: base.KeyRightBrace.EvCode, Leftless: true, BottomLeftlessAlt: true},
		{Label: base.KeyBlank.Label, Width: base.U1p75, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		base.KeyPad7, base.KeyPad8, base.KeyPad9, base.KeyPadPlus,
	},
	{
		base.KeyCapsLock, base.KeyA, base.KeyS, base.KeyD, base.KeyF, base.KeyG, base.KeyH, base.KeyJ, base.KeyK, base.KeyL, base.KeySemicolon, base.KeyApostrophe,
		{Label: "Enter ↵", Width: base.U2p75, Finger: base.KeyEnter.Finger, EvCode: base.KeyEnter.EvCode},
		base.KeyPad4,
		base.KeyPad5, base.KeyPad6,
		{Label: base.KeyPadPlus.Label, Width: base.KeyPadPlus.Width, Finger: base.KeyPadPlus.Finger, EvCode: base.KeyPadPlus.EvCode},
	},
	{
		base.KeyLeftShift, base.KeyZ, base.KeyX, base.KeyC, base.KeyV, base.KeyB, base.KeyN, base.KeyM, base.KeyComma, base.KeyDot, base.KeySlash,
		{Label: "Shift", Width: base.U2p5, Finger: base.KeyRightShift.Finger, EvCode: base.KeyRightShift.EvCode},
		base.KeyUp, base.KeyPad1, base.KeyPad2, base.KeyPad3, base.KeyPadEnter,
	},
	{
		base.KeyLeftCtrl, base.KeyLeftMeta,
		{Label: base.KeyLeftAlt.Label, Width: base.U1, Finger: base.KeyLeftAlt.Finger, EvCode: base.KeyLeftAlt.EvCode},
		base.KeyHanja,
		{Label: base.KeySpace.Label, Width: base.U3p75, Finger: base.KeySpace.Finger, EvCode: base.KeySpace.EvCode},
		base.KeyHangeul,
		{Label: base.KeyRightAlt.Label, Width: base.U1, Finger: base.KeyRightAlt.Finger, EvCode: base.KeyRightAlt.EvCode},
		base.KeyRightMeta, base.KeyRightCtrl, base.KeyLeft,
		base.KeyDown, base.KeyRight,
		{Label: base.KeyPad0.Label, Width: base.U1, Finger: base.Middle, EvCode: base.KeyPad0.EvCode},
		base.KeyPadDot,
		{Label: base.KeyPadEnter.Label, Width: base.KeyPadEnter.Width, Finger: base.KeyPadEnter.Finger, EvCode: base.KeyPadEnter.EvCode},
	},
}

var size100KS = [][]base.Key{
	{
		base.KeyEsc,
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode},
		base.KeyF1, base.KeyF2, base.KeyF3, base.KeyF4, base.KeyBlank, base.KeyF5, base.KeyF6, base.KeyF7, base.KeyF8, base.KeyBlank, base.KeyF9, base.KeyF10,
		base.KeyF11, base.KeyF12,
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		base.KeyPrintScreen, base.KeyScrollLock, base.KeyLightsToggle,
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Rightless: true},
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Rightless: true, Leftless: true},
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Rightless: true, Leftless: true},
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Rightless: true, Leftless: true},
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode},
	},
	{
		base.KeyGrave, base.Key1, base.Key2, base.Key3, base.Key4, base.Key5, base.Key6, base.Key7, base.Key8, base.Key9, base.Key0, base.KeyMinus, base.KeyEqual, base.KeyWon,
		{Label: "<--", Width: base.U1p75, Finger: base.KeyBackspace.Finger, EvCode: base.KeyBackspace.EvCode},
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		base.KeyInsert, base.KeyHome, base.KeyPageUp,
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		base.KeyNumLock, base.KeyPadSlash, base.KeyPadAsterisk, base.KeyPadMinus,
	},
	{
		base.KeyTab, base.KeyQ, base.KeyW, base.KeyE, base.KeyR, base.KeyT, base.KeyY, base.KeyU, base.KeyI, base.KeyO, base.KeyP, base.KeyLeftBrace,
		{Label: base.KeyRightBrace.Label, Width: base.KeyRightBrace.Width, Finger: base.KeyRightBrace.Finger, EvCode: base.KeyRightBrace.EvCode, Leftless: true, BottomLeftlessAlt: true},
		{Label: base.KeyBlank.Label, Width: base.U2p5, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, GapAlt: true},
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Leftless: true, BottomLeftless: true},
		base.KeyDelete, base.KeyEnd,
		{Label: base.KeyPageDown.Label, Width: base.KeyPageDown.Width, Finger: base.KeyPageDown.Finger, EvCode: base.KeyPageDown.EvCode, Leftless: true, BottomLeftlessAlt: true},
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		base.KeyPad7, base.KeyPad8, base.KeyPad9, base.KeyPadPlus,
	},
	{
		base.KeyCapsLock, base.KeyA, base.KeyS, base.KeyD, base.KeyF, base.KeyG, base.KeyH, base.KeyJ, base.KeyK, base.KeyL, base.KeySemicolon, base.KeyApostrophe,
		{Label: "Enter ↵", Width: base.U3p5, Finger: base.KeyEnter.Finger, EvCode: base.KeyEnter.EvCode},
		{Label: base.KeyBlank.Label, Width: base.U2, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Rightless: true, Leftless: false},
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: false, Rightless: true, Leftless: true, BottomLeftless: true},
		{Label: base.KeyBlank.Label, Width: base.U2, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Rightless: false, Leftless: false},
		base.KeyPad4, base.KeyPad5, base.KeyPad6,
		{Label: base.KeyPadPlus.Label, Width: base.KeyPadPlus.Width, Finger: base.KeyPadPlus.Finger, EvCode: base.KeyPadPlus.EvCode},
	},
	{
		{Label: "Shift", Width: base.U3, Finger: base.KeyLeftShift.Finger, EvCode: base.KeyLeftShift.EvCode},
		base.KeyZ, base.KeyX, base.KeyC, base.KeyV, base.KeyB, base.KeyN, base.KeyM, base.KeyComma, base.KeyDot, base.KeySlash,
		{Label: "Shift", Width: base.U3p75, Finger: base.KeyRightShift.Finger, EvCode: base.KeyRightShift.EvCode},
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Rightless: true, Leftless: false},
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode},
		base.KeyUp,
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Rightless: true, Leftless: true, BottomLeftless: true},
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		base.KeyPad1, base.KeyPad2, base.KeyPad3, base.KeyPadEnter,
	},
	{
		base.KeyLeftCtrl, base.KeyLeftMeta, base.KeyLeftAlt, base.KeyHanja,
		{Label: base.KeySpace.Label, Width: base.U4p5, Finger: base.KeySpace.Finger, EvCode: base.KeySpace.EvCode},
		base.KeyHangeul, base.KeyRightAlt, base.KeyRightMeta, base.KeyFn, base.KeyRightCtrl,
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode},
		base.KeyLeft, base.KeyDown, base.KeyRight,
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		base.KeyPad0, base.KeyPadDot,
		{Label: base.KeyPadEnter.Label, Width: base.KeyPadEnter.Width, Finger: base.KeyPadEnter.Finger, EvCode: base.KeyPadEnter.EvCode},
	},
}
