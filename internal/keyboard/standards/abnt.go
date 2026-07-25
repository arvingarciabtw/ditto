package standards

import "github.com/arvingarciabtw/ditto/internal/keyboard/base"

var SizesABNT = map[int][][]base.Key{
	60:  size60ABNT,
	65:  size65ABNT,
	75:  size75ABNT,
	80:  size80ABNT,
	96:  size96ABNT,
	100: size100ABNT,
}

var size60ABNT = [][]base.Key{
	{
		base.KeyApostrophe, base.Key1, base.Key2, base.Key3, base.Key4, base.Key5, base.Key6, base.Key7, base.Key8, base.Key9, base.Key0, base.KeyMinus, base.KeyEqual,
		{Label: base.KeyBackspace.Label, Width: base.U2p75, Finger: base.KeyBackspace.Finger, EvCode: base.KeyBackspace.EvCode},
	},
	{
		base.KeyTab, base.KeyQ, base.KeyW, base.KeyE, base.KeyR, base.KeyT, base.KeyY, base.KeyU, base.KeyI, base.KeyO, base.KeyP, base.KeyAcute,
		{Label: base.KeyLeftBrace.Label, Width: base.KeyLeftBrace.Width, Finger: base.KeyLeftBrace.Finger, EvCode: base.KeyLeftBrace.EvCode, Leftless: true, BottomLeftlessAlt: true},
		{Label: "Ent↵", Width: base.U2, Finger: base.KeyEnterISO.Finger, EvCode: base.KeyEnterISO.EvCode, Gap: base.KeyEnterISO.Gap},
	},
	{
		base.KeyCapsLock, base.KeyA, base.KeyS, base.KeyD, base.KeyF, base.KeyG, base.KeyH, base.KeyJ, base.KeyK, base.KeyL, base.KeyCedilla, base.KeyTilde, base.KeyRightBrace,
		{Label: base.KeyEnterISOBlank.Label, Width: base.U1p75, Finger: base.KeyEnterISOBlank.Finger, EvCode: base.KeyEnterISOBlank.EvCode},
	},
	{
		{Label: "Shf", Width: base.U1p75, Finger: base.KeyLeftShiftISO.Finger, EvCode: base.KeyLeftShiftISO.EvCode},
		{Label: base.KeyBackslash.Label, Width: base.U1, Finger: base.KeyBackslash.Finger, EvCode: base.KeyBackslash.EvCode},
		base.KeyZ, base.KeyX, base.KeyC, base.KeyV, base.KeyB, base.KeyN, base.KeyM, base.KeyComma, base.KeyDot, base.KeySemicolon, base.KeySlash,
		{Label: base.KeyRightShiftISO.Label, Width: base.U2, Finger: base.KeyRightShiftISO.Finger, EvCode: base.KeyRightShiftISO.EvCode},
	},
	{
		base.KeyLeftCtrl, base.KeyLeftMeta, base.KeyLeftAlt,
		{Label: base.KeySpace.Label, Width: base.U7p5, Finger: base.KeySpace.Finger, EvCode: base.KeySpace.EvCode},
		base.KeyRightAlt, base.KeyRightMeta, base.KeyFn, base.KeyRightCtrl,
	},
}

var size65ABNT = [][]base.Key{
	{
		base.KeyApostrophe, base.Key1, base.Key2, base.Key3, base.Key4, base.Key5, base.Key6, base.Key7, base.Key8, base.Key9, base.Key0, base.KeyMinus, base.KeyEqual,
		{Label: base.KeyBackspace.Label, Width: base.U2p75, Finger: base.KeyBackspace.Finger, EvCode: base.KeyBackspace.EvCode},
		{Label: base.KeyLightsToggle.Label, Width: base.KeyLightsToggle.Width, Finger: base.Any, EvCode: base.KeyLightsToggle.EvCode},
	},
	{
		base.KeyTab, base.KeyQ, base.KeyW, base.KeyE, base.KeyR, base.KeyT, base.KeyY, base.KeyU, base.KeyI, base.KeyO, base.KeyP, base.KeyAcute,
		{Label: base.KeyLeftBrace.Label, Width: base.KeyLeftBrace.Width, Finger: base.KeyLeftBrace.Finger, EvCode: base.KeyLeftBrace.EvCode, Leftless: true, BottomLeftlessAlt: true},
		{Label: "Ent↵", Width: base.U2, Finger: base.KeyEnterISO.Finger, EvCode: base.KeyEnterISO.EvCode, Gap: base.KeyEnterISO.Gap},
		{Label: base.KeyHome.Label, Width: base.KeyHome.Width, Finger: base.Any, EvCode: base.KeyHome.EvCode},
	},
	{
		base.KeyCapsLock, base.KeyA, base.KeyS, base.KeyD, base.KeyF, base.KeyG, base.KeyH, base.KeyJ, base.KeyK, base.KeyL, base.KeyCedilla, base.KeyTilde, base.KeyRightBrace,
		{Label: base.KeyEnterISOBlank.Label, Width: base.U1p75, Finger: base.KeyEnterISOBlank.Finger, EvCode: base.KeyEnterISOBlank.EvCode},
		{Label: base.KeyPageUp.Label, Width: base.KeyPageUp.Width, Finger: base.Any, EvCode: base.KeyPageUp.EvCode},
	},
	{
		{Label: base.KeyLeftShiftISO.Label, Width: base.U1p5, Finger: base.KeyLeftShiftISO.Finger, EvCode: base.KeyLeftShiftISO.EvCode},
		{Label: base.KeyBackslash.Label, Width: base.U1, Finger: base.KeyBackslash.Finger, EvCode: base.KeyBackslash.EvCode},
		base.KeyZ, base.KeyX, base.KeyC, base.KeyV, base.KeyB, base.KeyN, base.KeyM, base.KeyComma, base.KeyDot, base.KeySemicolon, base.KeySlash,
		{Label: "Shf", Width: base.U1, Finger: base.KeyRightShiftISO.Finger, EvCode: base.KeyRightShiftISO.EvCode},
		base.KeyUp,
		{Label: base.KeyPageDown.Label, Width: base.KeyPageDown.Width, Finger: base.Any, EvCode: base.KeyPageDown.EvCode},
	},
	{
		base.KeyLeftCtrl, base.KeyLeftMeta, base.KeyLeftAlt,
		{Label: base.KeySpace.Label, Width: base.U6, Finger: base.KeySpace.Finger, EvCode: base.KeySpace.EvCode},
		base.KeyRightAlt, base.KeyFn, base.KeyRightCtrl, base.KeyLeft, base.KeyDown, base.KeyRight,
	},
}

var size75ABNT = [][]base.Key{
	{
		base.KeyEsc, base.KeyF1, base.KeyF2, base.KeyF3, base.KeyF4, base.KeyF5, base.KeyF6, base.KeyF7, base.KeyF8, base.KeyF9, base.KeyF10, base.KeyF11, base.KeyF12,
		{Label: base.KeyPrintScreen.Label, Width: base.KeyPrintScreen.Width, Finger: base.Any, EvCode: base.KeyPrintScreen.EvCode},
		{Label: base.KeyDelete.Label, Width: base.KeyDelete.Width, Finger: base.Any, EvCode: base.KeyDelete.EvCode},
		{Label: base.KeyLightsToggle.Label, Width: base.KeyLightsToggle.Width, Finger: base.Any, EvCode: base.KeyLightsToggle.EvCode},
	},
	{
		base.KeyApostrophe, base.Key1, base.Key2, base.Key3, base.Key4, base.Key5, base.Key6, base.Key7, base.Key8, base.Key9, base.Key0, base.KeyMinus, base.KeyEqual,
		base.KeyBackspace,
		{Label: base.KeyPageUp.Label, Width: base.KeyPageUp.Width, Finger: base.Any, EvCode: base.KeyPageUp.EvCode},
	},
	{
		base.KeyTab, base.KeyQ, base.KeyW, base.KeyE, base.KeyR, base.KeyT, base.KeyY, base.KeyU, base.KeyI, base.KeyO, base.KeyP, base.KeyAcute,
		{Label: base.KeyLeftBrace.Label, Width: base.KeyLeftBrace.Width, Finger: base.KeyLeftBrace.Finger, EvCode: base.KeyLeftBrace.EvCode, Leftless: true, BottomLeftlessAlt: true},
		base.KeyEnterISO,
		{Label: base.KeyPageDown.Label, Width: base.KeyPageDown.Width, Finger: base.Any, EvCode: base.KeyPageDown.EvCode},
	},
	{
		base.KeyCapsLock, base.KeyA, base.KeyS, base.KeyD, base.KeyF, base.KeyG, base.KeyH, base.KeyJ, base.KeyK, base.KeyL, base.KeyCedilla, base.KeyTilde, base.KeyRightBrace,
		base.KeyEnterISOBlank,
		{Label: base.KeyHome.Label, Width: base.KeyHome.Width, Finger: base.Any, EvCode: base.KeyHome.EvCode},
	},
	{
		{Label: "Shf", Width: base.U1, Finger: base.KeyLeftShiftISO.Finger, EvCode: base.KeyLeftShiftISO.EvCode},
		{Label: base.KeyBackslash.Label, Width: base.U1, Finger: base.KeyBackslash.Finger, EvCode: base.KeyBackslash.EvCode},
		base.KeyZ, base.KeyX, base.KeyC, base.KeyV, base.KeyB, base.KeyN, base.KeyM, base.KeyComma, base.KeyDot, base.KeySemicolon, base.KeySlash,
		{Label: "Shf", Width: base.U1, Finger: base.KeyRightShiftISO.Finger, EvCode: base.KeyRightShiftISO.EvCode},
		base.KeyUp,
		{Label: base.KeyEnd.Label, Width: base.KeyEnd.Width, Finger: base.Any, EvCode: base.KeyEnd.EvCode},
	},
	{
		base.KeyLeftCtrl, base.KeyLeftMeta, base.KeyLeftAlt, base.KeySpace, base.KeyRightAlt, base.KeyFn, base.KeyRightCtrl, base.KeyLeft,
		base.KeyDown, base.KeyRight,
	},
}

var size80ABNT = [][]base.Key{
	{
		base.KeyEsc,
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode},
		base.KeyF1, base.KeyF2, base.KeyF3, base.KeyF4, base.KeyBlank, base.KeyF5, base.KeyF6, base.KeyF7, base.KeyF8, base.KeyBlank, base.KeyF9, base.KeyF10,
		base.KeyF11, base.KeyF12,
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		base.KeyPrintScreen, base.KeyScrollLock, base.KeyLightsToggle,
	},
	{
		base.KeyApostrophe, base.Key1, base.Key2, base.Key3, base.Key4, base.Key5, base.Key6, base.Key7, base.Key8, base.Key9, base.Key0, base.KeyMinus, base.KeyEqual,
		{Label: base.KeyBackspace.Label, Width: base.U3, Finger: base.KeyBackspace.Finger, EvCode: base.KeyBackspace.EvCode},
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		base.KeyInsert, base.KeyHome, base.KeyPageUp,
	},
	{
		{Label: "Tab", Width: base.U2p5, Finger: base.KeyTab.Finger, EvCode: base.KeyTab.EvCode},
		base.KeyQ, base.KeyW, base.KeyE,
		base.KeyR, base.KeyT, base.KeyY, base.KeyU, base.KeyI, base.KeyO, base.KeyP, base.KeyAcute,
		{Label: base.KeyLeftBrace.Label, Width: base.KeyLeftBrace.Width, Finger: base.KeyLeftBrace.Finger, EvCode: base.KeyLeftBrace.EvCode, Leftless: true, BottomLeftlessAlt: true},
		{Label: base.KeyEnterISO.Label, Width: base.KeyEnterISO.Width, Finger: base.KeyEnterISO.Finger, EvCode: base.KeyEnterISO.EvCode, Gap: base.KeyEnterISO.Gap, GapAlt: true},
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Leftless: true, BottomLeftless: true},
		base.KeyDelete, base.KeyEnd, base.KeyPageDown,
	},
	{
		{Label: base.KeyCapsLock.Label, Width: base.U2p75, Finger: base.KeyCapsLock.Finger, EvCode: base.KeyCapsLock.EvCode},
		base.KeyA, base.KeyS, base.KeyD, base.KeyF, base.KeyG, base.KeyH, base.KeyJ, base.KeyK, base.KeyL, base.KeyCedilla, base.KeyTilde, base.KeyRightBrace,
		base.KeyEnterISOBlank,
		{Label: base.KeyBlank.Label, Width: base.U2, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Rightless: true, Leftless: false},
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: false, Rightless: true, Leftless: true, BottomLeftless: true},
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Rightless: false, Leftless: false},
	},
	{
		{Label: base.KeyLeftShiftISO.Label, Width: base.U2, Finger: base.KeyLeftShiftISO.Finger, EvCode: base.KeyLeftShiftISO.EvCode},
		{Label: base.KeyBackslash.Label, Width: base.U1, Finger: base.KeyBackslash.Finger, EvCode: base.KeyBackslash.EvCode},
		base.KeyZ, base.KeyX, base.KeyC, base.KeyV, base.KeyB, base.KeyN, base.KeyM, base.KeyComma, base.KeyDot, base.KeySemicolon, base.KeySlash,
		{Label: base.KeyRightShiftISO.Label, Width: base.U2, Finger: base.KeyRightShiftISO.Finger, EvCode: base.KeyRightShiftISO.EvCode},
		{Label: base.KeyBlank.Label, Width: base.U0p75, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Rightless: true, Leftless: false},
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

var size96ABNT = [][]base.Key{
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
		base.KeyApostrophe, base.Key1, base.Key2, base.Key3, base.Key4, base.Key5, base.Key6, base.Key7, base.Key8, base.Key9, base.Key0, base.KeyMinus, base.KeyEqual,
		base.KeyBackspace, base.KeyNumLock, base.KeyPadSlash, base.KeyPadAsterisk, base.KeyPadMinus,
	},
	{
		base.KeyTab, base.KeyQ, base.KeyW, base.KeyE, base.KeyR, base.KeyT, base.KeyY, base.KeyU, base.KeyI, base.KeyO, base.KeyP, base.KeyAcute,
		{Label: base.KeyLeftBrace.Label, Width: base.KeyLeftBrace.Width, Finger: base.KeyLeftBrace.Finger, EvCode: base.KeyLeftBrace.EvCode, Leftless: true, BottomLeftlessAlt: true},
		base.KeyEnterISO, base.KeyPad7, base.KeyPad8, base.KeyPad9, base.KeyPadPlus,
	},
	{
		base.KeyCapsLock, base.KeyA, base.KeyS, base.KeyD, base.KeyF, base.KeyG, base.KeyH, base.KeyJ, base.KeyK, base.KeyL, base.KeyCedilla, base.KeyTilde, base.KeyRightBrace,
		base.KeyEnterISOBlank, base.KeyPad4, base.KeyPad5, base.KeyPad6,
		{Label: base.KeyPadPlus.Label, Width: base.KeyPadPlus.Width, Finger: base.KeyPadPlus.Finger, EvCode: base.KeyPadPlus.EvCode},
	},
	{
		{Label: "Shf", Width: base.U1, Finger: base.KeyLeftShiftISO.Finger, EvCode: base.KeyLeftShiftISO.EvCode},
		{Label: base.KeyBackslash.Label, Width: base.U1, Finger: base.KeyBackslash.Finger, EvCode: base.KeyBackslash.EvCode},
		base.KeyZ, base.KeyX, base.KeyC, base.KeyV, base.KeyB, base.KeyN, base.KeyM, base.KeyComma, base.KeyDot, base.KeySemicolon, base.KeySlash,
		{Label: "Shf", Width: base.U1, Finger: base.KeyRightShiftISO.Finger, EvCode: base.KeyRightShiftISO.EvCode},
		base.KeyUp,
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

var size100ABNT = [][]base.Key{
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
		base.KeyApostrophe, base.Key1, base.Key2, base.Key3, base.Key4, base.Key5, base.Key6, base.Key7, base.Key8, base.Key9, base.Key0, base.KeyMinus, base.KeyEqual,
		{Label: base.KeyBackspace.Label, Width: base.U3, Finger: base.KeyBackspace.Finger, EvCode: base.KeyBackspace.EvCode},
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		base.KeyInsert, base.KeyHome, base.KeyPageUp,
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		base.KeyNumLock, base.KeyPadSlash, base.KeyPadAsterisk, base.KeyPadMinus,
	},
	{
		{Label: "Tab", Width: base.U2p5, Finger: base.KeyTab.Finger, EvCode: base.KeyTab.EvCode},
		base.KeyQ, base.KeyW, base.KeyE,
		base.KeyR, base.KeyT, base.KeyY, base.KeyU, base.KeyI, base.KeyO, base.KeyP, base.KeyAcute,
		{Label: base.KeyLeftBrace.Label, Width: base.KeyLeftBrace.Width, Finger: base.KeyLeftBrace.Finger, EvCode: base.KeyLeftBrace.EvCode, Leftless: true, BottomLeftlessAlt: true},
		{Label: base.KeyEnterISO.Label, Width: base.KeyEnterISO.Width, Finger: base.KeyEnterISO.Finger, EvCode: base.KeyEnterISO.EvCode, Gap: base.KeyEnterISO.Gap, GapAlt: true},
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Leftless: true, BottomLeftless: true},
		base.KeyDelete, base.KeyEnd,
		{Label: base.KeyPageDown.Label, Width: base.KeyPageDown.Width, Finger: base.KeyPageDown.Finger, EvCode: base.KeyPageDown.EvCode, Leftless: true, BottomLeftlessAlt: true},
		{Label: base.KeyBlank.Label, Width: base.KeyBlank.Width, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true},
		base.KeyPad7, base.KeyPad8, base.KeyPad9, base.KeyPadPlus,
	},
	{
		{Label: base.KeyCapsLock.Label, Width: base.U2p75, Finger: base.KeyCapsLock.Finger, EvCode: base.KeyCapsLock.EvCode},
		base.KeyA, base.KeyS, base.KeyD, base.KeyF, base.KeyG, base.KeyH, base.KeyJ, base.KeyK, base.KeyL, base.KeyCedilla, base.KeyTilde, base.KeyRightBrace,
		base.KeyEnterISOBlank,
		{Label: base.KeyBlank.Label, Width: base.U2, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Rightless: true, Leftless: false},
		{Label: base.KeyBlank.Label, Width: base.U1, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: false, Rightless: true, Leftless: true, BottomLeftless: true},
		{Label: base.KeyBlank.Label, Width: base.U2, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Rightless: false, Leftless: false},
		base.KeyPad4, base.KeyPad5, base.KeyPad6,
		{Label: base.KeyPadPlus.Label, Width: base.KeyPadPlus.Width, Finger: base.KeyPadPlus.Finger, EvCode: base.KeyPadPlus.EvCode},
	},
	{
		{Label: base.KeyLeftShiftISO.Label, Width: base.U2, Finger: base.KeyLeftShiftISO.Finger, EvCode: base.KeyLeftShiftISO.EvCode},
		{Label: base.KeyBackslash.Label, Width: base.U1, Finger: base.KeyBackslash.Finger, EvCode: base.KeyBackslash.EvCode},
		base.KeyZ, base.KeyX, base.KeyC, base.KeyV, base.KeyB, base.KeyN, base.KeyM, base.KeyComma, base.KeyDot, base.KeySemicolon, base.KeySlash,
		{Label: base.KeyRightShiftISO.Label, Width: base.U2, Finger: base.KeyRightShiftISO.Finger, EvCode: base.KeyRightShiftISO.EvCode},
		{Label: base.KeyBlank.Label, Width: base.U0p75, Finger: base.KeyBlank.Finger, EvCode: base.KeyBlank.EvCode, Gap: true, Rightless: true, Leftless: false},
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
