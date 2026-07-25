package standards

import "github.com/arvingarciabtw/ditto/internal/keyboard/base"

var SizesISO = map[int][][]base.Key{
	60:  size60ISO,
	65:  size65ISO,
	75:  size75ISO,
	80:  size80ISO,
	96:  size96ISO,
	100: size100ISO,
}

var size60ISO = [][]base.Key{
	{
		base.KeyGrave, base.Key1, base.Key2, base.Key3, base.Key4, base.Key5, base.Key6, base.Key7, base.Key8, base.Key9, base.Key0, base.KeyMinus, base.KeyEqual,
		{Label: base.KeyBackspace.Label, Width: base.U2p75, Finger: base.KeyBackspace.Finger, EvCode: base.KeyBackspace.EvCode},
	},
	{
		base.KeyTab, base.KeyQ, base.KeyW, base.KeyE, base.KeyR, base.KeyT, base.KeyY, base.KeyU, base.KeyI, base.KeyO, base.KeyP, base.KeyLeftBrace,
		{Label: base.KeyRightBrace.Label, Width: base.KeyRightBrace.Width, Finger: base.KeyRightBrace.Finger, EvCode: base.KeyRightBrace.EvCode, Leftless: true, BottomLeftlessAlt: true},
		{Label: "Ent↵", Width: base.U2, Finger: base.KeyEnterISO.Finger, EvCode: base.KeyEnterISO.EvCode, Gap: base.KeyEnterISO.Gap, Leftless: true, BottomLeftless: true},
	},
	{
		base.KeyCapsLock, base.KeyA, base.KeyS, base.KeyD, base.KeyF, base.KeyG, base.KeyH, base.KeyJ, base.KeyK, base.KeyL, base.KeySemicolon, base.KeyApostrophe,
		base.KeyPound,
		{Label: base.KeyEnterISOBlank.Label, Width: base.U1p75, Finger: base.KeyEnterISOBlank.Finger, EvCode: base.KeyEnterISOBlank.EvCode},
	},
	{
		{Label: base.KeyLeftShiftISO.Label, Width: base.U2, Finger: base.KeyLeftShiftISO.Finger, EvCode: base.KeyLeftShiftISO.EvCode},
		{Label: base.KeyBackslash.Label, Width: base.U1, Finger: base.KeyBackslash.Finger, EvCode: base.KeyBackslash.EvCode},
		base.KeyZ, base.KeyX, base.KeyC, base.KeyV, base.KeyB, base.KeyN, base.KeyM, base.KeyComma, base.KeyDot, base.KeySlash, base.KeyRightShift,
	},
	{
		base.KeyLeftCtrl, base.KeyLeftMeta, base.KeyLeftAlt,
		{Label: base.KeySpace.Label, Width: base.U7p5, Finger: base.KeySpace.Finger, EvCode: base.KeySpace.EvCode},
		base.KeyRightAlt, base.KeyRightMeta, base.KeyFn, base.KeyRightCtrl,
	},
}

var size65ISO = [][]base.Key{
	{
		base.KeyGrave, base.Key1, base.Key2, base.Key3, base.Key4, base.Key5, base.Key6, base.Key7, base.Key8, base.Key9, base.Key0, base.KeyMinus, base.KeyEqual,
		{Label: base.KeyBackspace.Label, Width: base.U2p75, Finger: base.KeyBackspace.Finger, EvCode: base.KeyBackspace.EvCode},
		{Label: base.KeyLightsToggle.Label, Width: base.KeyLightsToggle.Width, Finger: base.Any, EvCode: base.KeyLightsToggle.EvCode},
	},
	{
		base.KeyTab, base.KeyQ, base.KeyW, base.KeyE, base.KeyR, base.KeyT, base.KeyY, base.KeyU, base.KeyI, base.KeyO, base.KeyP, base.KeyLeftBrace,
		{Label: base.KeyRightBrace.Label, Width: base.KeyRightBrace.Width, Finger: base.KeyRightBrace.Finger, EvCode: base.KeyRightBrace.EvCode, Leftless: true, BottomLeftlessAlt: true},
		{Label: "Ent↵", Width: base.U2, Finger: base.KeyEnterISO.Finger, EvCode: base.KeyEnterISO.EvCode, Gap: base.KeyEnterISO.Gap},
		{Label: base.KeyHome.Label, Width: base.KeyHome.Width, Finger: base.Any, EvCode: base.KeyHome.EvCode},
	},
	{
		base.KeyCapsLock, base.KeyA, base.KeyS, base.KeyD, base.KeyF, base.KeyG, base.KeyH, base.KeyJ, base.KeyK, base.KeyL, base.KeySemicolon, base.KeyApostrophe,
		base.KeyPound,
		{Label: base.KeyEnterISOBlank.Label, Width: base.U1p75, Finger: base.KeyEnterISOBlank.Finger, EvCode: base.KeyEnterISOBlank.EvCode},
		{Label: base.KeyPageUp.Label, Width: base.KeyPageUp.Width, Finger: base.Any, EvCode: base.KeyPageUp.EvCode},
	},
	{
		base.KeyLeftShiftISO,
		{Label: base.KeyBackslash.Label, Width: base.U1, Finger: base.KeyBackslash.Finger, EvCode: base.KeyBackslash.EvCode},
		base.KeyZ, base.KeyX, base.KeyC, base.KeyV, base.KeyB, base.KeyN, base.KeyM, base.KeyComma, base.KeyDot, base.KeySlash,
		{Label: base.KeyRightShift.Label, Width: base.U2p5, Finger: base.KeyRightShift.Finger, EvCode: base.KeyRightShift.EvCode},
		base.KeyUp,
		{Label: base.KeyPageDown.Label, Width: base.KeyPageDown.Width, Finger: base.Any, EvCode: base.KeyPageDown.EvCode},
	},
	{
		base.KeyLeftCtrl, base.KeyLeftMeta, base.KeyLeftAlt,
		{Label: base.KeySpace.Label, Width: base.U6, Finger: base.KeySpace.Finger, EvCode: base.KeySpace.EvCode},
		base.KeyRightAlt, base.KeyFn, base.KeyRightCtrl, base.KeyLeft, base.KeyDown, base.KeyRight,
	},
}

var size75ISO = [][]base.Key{
	{
		base.KeyEsc, base.KeyF1, base.KeyF2, base.KeyF3, base.KeyF4, base.KeyF5, base.KeyF6, base.KeyF7, base.KeyF8, base.KeyF9, base.KeyF10, base.KeyF11, base.KeyF12,
		{Label: base.KeyPrintScreen.Label, Width: base.KeyPrintScreen.Width, Finger: base.Any, EvCode: base.KeyPrintScreen.EvCode},
		{Label: base.KeyDelete.Label, Width: base.KeyDelete.Width, Finger: base.Any, EvCode: base.KeyDelete.EvCode},
		{Label: base.KeyLightsToggle.Label, Width: base.KeyLightsToggle.Width, Finger: base.Any, EvCode: base.KeyLightsToggle.EvCode},
	},
	{
		base.KeyGrave, base.Key1, base.Key2, base.Key3, base.Key4, base.Key5, base.Key6, base.Key7, base.Key8, base.Key9, base.Key0, base.KeyMinus, base.KeyEqual,
		base.KeyBackspace,
		{Label: base.KeyPageUp.Label, Width: base.KeyPageUp.Width, Finger: base.Any, EvCode: base.KeyPageUp.EvCode},
	},
	{
		base.KeyTab, base.KeyQ, base.KeyW, base.KeyE, base.KeyR, base.KeyT, base.KeyY, base.KeyU, base.KeyI, base.KeyO, base.KeyP, base.KeyLeftBrace,
		{Label: base.KeyRightBrace.Label, Width: base.KeyRightBrace.Width, Finger: base.KeyRightBrace.Finger, EvCode: base.KeyRightBrace.EvCode, Leftless: true, BottomLeftlessAlt: true},
		base.KeyEnterISO,
		{Label: base.KeyPageDown.Label, Width: base.KeyPageDown.Width, Finger: base.Any, EvCode: base.KeyPageDown.EvCode},
	},
	{
		base.KeyCapsLock, base.KeyA, base.KeyS, base.KeyD, base.KeyF, base.KeyG, base.KeyH, base.KeyJ, base.KeyK, base.KeyL, base.KeySemicolon, base.KeyApostrophe,
		base.KeyPound, base.KeyEnterISOBlank,
		{Label: base.KeyHome.Label, Width: base.KeyHome.Width, Finger: base.Any, EvCode: base.KeyHome.EvCode},
	},
	{
		base.KeyLeftShiftISO,
		{Label: base.KeyBackslash.Label, Width: base.U1, Finger: base.KeyBackslash.Finger, EvCode: base.KeyBackslash.EvCode},
		base.KeyZ, base.KeyX, base.KeyC, base.KeyV, base.KeyB, base.KeyN, base.KeyM, base.KeyComma, base.KeyDot, base.KeySlash, base.KeyRightShiftISO, base.KeyUp,
		{Label: base.KeyEnd.Label, Width: base.KeyEnd.Width, Finger: base.Any, EvCode: base.KeyEnd.EvCode},
	},
	{
		base.KeyLeftCtrl, base.KeyLeftMeta, base.KeyLeftAlt, base.KeySpace, base.KeyRightAlt, base.KeyFn, base.KeyRightCtrl, base.KeyLeft,
		base.KeyDown, base.KeyRight,
	},
}

var size80ISO = [][]base.Key{
	{
		base.KeyEsc,
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode},
		base.KeyF1, base.KeyF2, base.KeyF3, base.KeyF4, base.KeyBlank, base.KeyF5, base.KeyF6, base.KeyF7, base.KeyF8, base.KeyBlank, base.KeyF9, base.KeyF10,
		base.KeyF11, base.KeyF12,
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		base.KeyPrintScreen, base.KeyScrollLock, base.KeyLightsToggle,
	},
	{
		base.KeyGrave, base.Key1, base.Key2, base.Key3, base.Key4, base.Key5, base.Key6, base.Key7, base.Key8, base.Key9, base.Key0, base.KeyMinus, base.KeyEqual,
		{Label: base.KeyBackspace.Label, Width: base.U3, Finger: base.KeyBackspace.Finger, EvCode: base.KeyBackspace.EvCode},
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		base.KeyInsert, base.KeyHome, base.KeyPageUp,
	},
	{
		{Label: "Tab", Width: base.U2p5, Finger: base.KeyTab.Finger, EvCode: base.KeyTab.EvCode},
		base.KeyQ, base.KeyW, base.KeyE,
		base.KeyR, base.KeyT, base.KeyY, base.KeyU, base.KeyI, base.KeyO, base.KeyP, base.KeyLeftBrace,
		{Label: base.KeyRightBrace.Label, Width: base.KeyRightBrace.Width, Finger: base.KeyRightBrace.Finger, EvCode: base.KeyRightBrace.EvCode, Leftless: true, BottomLeftlessAlt: true},
		{Label: base.KeyEnterISO.Label, Width: base.KeyEnterISO.Width, Finger: base.KeyEnterISO.Finger, EvCode: base.KeyEnterISO.EvCode, Gap: base.KeyEnterISO.Gap, GapAlt: true},
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Leftless: true, BottomLeftless: true},
		base.KeyDelete, base.KeyEnd, base.KeyPageDown,
	},
	{
		{Label: base.KeyCapsLock.Label, Width: base.U2p75, Finger: base.KeyCapsLock.Finger, EvCode: base.KeyCapsLock.EvCode},
		base.KeyA, base.KeyS, base.KeyD, base.KeyF, base.KeyG, base.KeyH, base.KeyJ, base.KeyK, base.KeyL, base.KeySemicolon, base.KeyApostrophe, base.KeyPound,
		base.KeyEnterISOBlank,
		{Label: base.KeyBlank.Label, Width: base.U2, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Rightless: true, Leftless: false},
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: false, Rightless: true, Leftless: true, BottomLeftless: true},
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Rightless: false, Leftless: false},
	},
	{
		base.KeyLeftShift,
		{Label: base.KeyBackslash.Label, Width: base.U1, Finger: base.KeyBackslash.Finger, EvCode: base.KeyBackslash.EvCode},
		base.KeyZ, base.KeyX, base.KeyC, base.KeyV, base.KeyB, base.KeyN, base.KeyM, base.KeyComma, base.KeyDot, base.KeySlash, base.KeyRightShift,
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Rightless: true, Leftless: false},
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode},
		base.KeyUp,
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode},
	},
	{
		base.KeyLeftCtrl, base.KeyLeftMeta, base.KeyLeftAlt,
		{Label: base.KeySpace.Label, Width: base.U7p75, Finger: base.KeySpace.Finger, EvCode: base.KeySpace.EvCode},
		base.KeyRightAlt, base.KeyRightMeta, base.KeyFn, base.KeyRightCtrl,
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode},
		base.KeyLeft, base.KeyDown, base.KeyRight,
	},
}

var size96ISO = [][]base.Key{
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
		base.KeyGrave, base.Key1, base.Key2, base.Key3, base.Key4, base.Key5, base.Key6, base.Key7, base.Key8, base.Key9, base.Key0, base.KeyMinus, base.KeyEqual,
		base.KeyBackspace, base.KeyNumLock, base.KeyPadSlash, base.KeyPadAsterisk, base.KeyPadMinus,
	},
	{
		base.KeyTab, base.KeyQ, base.KeyW, base.KeyE, base.KeyR, base.KeyT, base.KeyY, base.KeyU, base.KeyI, base.KeyO, base.KeyP, base.KeyLeftBrace,
		{Label: base.KeyRightBrace.Label, Width: base.KeyRightBrace.Width, Finger: base.KeyRightBrace.Finger, EvCode: base.KeyRightBrace.EvCode, Leftless: true, BottomLeftlessAlt: true},
		base.KeyEnterISO, base.KeyPad7, base.KeyPad8, base.KeyPad9, base.KeyPadPlus,
	},
	{
		base.KeyCapsLock, base.KeyA, base.KeyS, base.KeyD, base.KeyF, base.KeyG, base.KeyH, base.KeyJ, base.KeyK, base.KeyL, base.KeySemicolon, base.KeyApostrophe,
		base.KeyPound, base.KeyEnterISOBlank, base.KeyPad4, base.KeyPad5, base.KeyPad6,
		{Label: base.KeyPadPlus.Label, Width: base.KeyPadPlus.Width, Finger: base.KeyPadPlus.Finger, EvCode: base.KeyPadPlus.EvCode},
	},
	{
		base.KeyLeftShiftISO,
		{Label: base.KeyBackslash.Label, Width: base.U1, Finger: base.KeyBackslash.Finger, EvCode: base.KeyBackslash.EvCode},
		base.KeyZ, base.KeyX, base.KeyC, base.KeyV, base.KeyB, base.KeyN, base.KeyM, base.KeyComma, base.KeyDot, base.KeySlash, base.KeyRightShiftISO, base.KeyUp,
		base.KeyPad1, base.KeyPad2, base.KeyPad3, base.KeyPadEnter,
	},
	{
		base.KeyLeftCtrl, base.KeyLeftMeta, base.KeyLeftAlt, base.KeySpace, base.KeyRightAlt, base.KeyFn, base.KeyRightCtrl, base.KeyLeft,
		base.KeyDown, base.KeyRight,
		{Label: base.KeyPad0.Label, Width: base.U1, Finger: base.Middle, EvCode: base.KeyPad0.EvCode},
		base.KeyPadDot,
		{Label: base.KeyPadEnter.Label, Width: base.KeyPadEnter.Width, Finger: base.KeyPadEnter.Finger, EvCode: base.KeyPadEnter.EvCode},
	},
}

var size100ISO = [][]base.Key{
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
		base.KeyGrave, base.Key1, base.Key2, base.Key3, base.Key4, base.Key5, base.Key6, base.Key7, base.Key8, base.Key9, base.Key0, base.KeyMinus, base.KeyEqual,
		{Label: base.KeyBackspace.Label, Width: base.U3, Finger: base.KeyBackspace.Finger, EvCode: base.KeyBackspace.EvCode},
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		base.KeyInsert, base.KeyHome, base.KeyPageUp,
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		base.KeyNumLock, base.KeyPadSlash, base.KeyPadAsterisk, base.KeyPadMinus,
	},
	{
		{Label: "Tab", Width: base.U2p5, Finger: base.KeyTab.Finger, EvCode: base.KeyTab.EvCode},
		base.KeyQ, base.KeyW, base.KeyE,
		base.KeyR, base.KeyT, base.KeyY, base.KeyU, base.KeyI, base.KeyO, base.KeyP, base.KeyLeftBrace,
		{Label: base.KeyRightBrace.Label, Width: base.KeyRightBrace.Width, Finger: base.KeyRightBrace.Finger, EvCode: base.KeyRightBrace.EvCode, Leftless: true, BottomLeftlessAlt: true},
		{Label: base.KeyEnterISO.Label, Width: base.KeyEnterISO.Width, Finger: base.KeyEnterISO.Finger, EvCode: base.KeyEnterISO.EvCode, Gap: base.KeyEnterISO.Gap, GapAlt: true},
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Leftless: true, BottomLeftless: true},
		base.KeyDelete, base.KeyEnd,
		{Label: base.KeyPageDown.Label, Width: base.KeyPageDown.Width, Finger: base.KeyPageDown.Finger, EvCode: base.KeyPageDown.EvCode, Leftless: true, BottomLeftlessAlt: true},
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		base.KeyPad7, base.KeyPad8, base.KeyPad9, base.KeyPadPlus,
	},
	{
		{Label: base.KeyCapsLock.Label, Width: base.U2p75, Finger: base.KeyCapsLock.Finger, EvCode: base.KeyCapsLock.EvCode},
		base.KeyA, base.KeyS, base.KeyD, base.KeyF, base.KeyG, base.KeyH, base.KeyJ, base.KeyK, base.KeyL, base.KeySemicolon, base.KeyApostrophe, base.KeyPound,
		base.KeyEnterISOBlank,
		{Label: base.KeyBlank.Label, Width: base.U2, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Rightless: true, Leftless: false},
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: false, Rightless: true, Leftless: true, BottomLeftless: true},
		{Label: base.KeyBlank.Label, Width: base.U2, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Rightless: false, Leftless: false},
		base.KeyPad4, base.KeyPad5, base.KeyPad6,
		{Label: base.KeyPadPlus.Label, Width: base.KeyPadPlus.Width, Finger: base.KeyPadPlus.Finger, EvCode: base.KeyPadPlus.EvCode},
	},
	{
		base.KeyLeftShift,
		{Label: base.KeyBackslash.Label, Width: base.U1, Finger: base.KeyBackslash.Finger, EvCode: base.KeyBackslash.EvCode},
		base.KeyZ, base.KeyX, base.KeyC, base.KeyV, base.KeyB, base.KeyN, base.KeyM, base.KeyComma, base.KeyDot, base.KeySlash, base.KeyRightShift,
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Rightless: true, Leftless: false},
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode},
		base.KeyUp,
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Rightless: true, Leftless: true, BottomLeftless: true},
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		base.KeyPad1, base.KeyPad2, base.KeyPad3, base.KeyPadEnter,
	},
	{
		base.KeyLeftCtrl, base.KeyLeftMeta, base.KeyLeftAlt,
		{Label: base.KeySpace.Label, Width: base.U7p75, Finger: base.KeySpace.Finger, EvCode: base.KeySpace.EvCode},
		base.KeyRightAlt, base.KeyRightMeta, base.KeyFn, base.KeyRightCtrl,
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode},
		base.KeyLeft, base.KeyDown, base.KeyRight,
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		base.KeyPad0, base.KeyPadDot,
		{Label: base.KeyPadEnter.Label, Width: base.KeyPadEnter.Width, Finger: base.KeyPadEnter.Finger, EvCode: base.KeyPadEnter.EvCode},
	},
}
