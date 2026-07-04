package keyboard

import (
	"fmt"

	"github.com/arvingarciabtw/ditto/internal/keyboard/base"
)

var layouts = map[string]map[string]string{
	"qwerty": nil,

	"dvorak": {
		"-": "[", "=": "]",

		"Q": "'", "W": ",", "E": ".", "R": "P", "T": "Y",
		"Y": "F", "U": "G", "I": "C", "O": "R", "P": "L",
		"[": "/", "]": "=",

		"A": "A", "S": "O", "D": "E", "F": "U", "G": "I",
		"H": "D", "J": "H", "K": "T", "L": "N",
		";": "S", "'": "-",

		"Z": ";", "X": "Q", "C": "J", "V": "K",
		"B": "X", "N": "B", "M": "M",
		",": "W", ".": "V", "/": "Z",
	},

	"dvorak uk": {
		"-": "[", "=": "]",

		"Q": "'", "W": ",", "E": ".", "R": "P", "T": "Y",
		"Y": "F", "U": "G", "I": "C", "O": "R", "P": "L",
		"[": "/", "]": "=",

		"A": "A", "S": "O", "D": "E", "F": "U", "G": "I",
		"H": "D", "J": "H", "K": "T", "L": "N",
		";": "S", "'": "-",

		"Z": ";", "X": "Q", "C": "J", "V": "K",
		"B": "X", "N": "B", "M": "M",
		",": "W", ".": "V", "/": "Z",
	},

	"colemak": {
		"Caps": "<--",

		"Q": "Q", "W": "W", "E": "F", "R": "P", "T": "G",
		"Y": "J", "U": "L", "I": "U", "O": "Y", "P": ";",

		"A": "A", "S": "R", "D": "S", "F": "T", "G": "D",
		"H": "H", "J": "N", "K": "E", "L": "I",
		";": "O", "'": "'",

		"Z": "Z", "X": "X", "C": "C", "V": "V",
		"B": "B", "N": "K", "M": "M",
	},

	"colemak-dh": {
		"Caps": "<--",

		"Q": "Q", "W": "W", "E": "F", "R": "P", "T": "B",
		"Y": "J", "U": "L", "I": "U", "O": "Y", "P": ";",

		"A": "A", "S": "R", "D": "S", "F": "T", "G": "G",
		"H": "M", "J": "N", "K": "E", "L": "I",
		";": "O",

		"Z": "Z", "X": "X", "C": "C", "V": "D",
		"B": "V", "N": "K", "M": "H",
	},

	"workman": {
		"Q": "Q", "W": "D", "E": "R", "R": "W", "T": "B",
		"Y": "J", "U": "F", "I": "U", "O": "P", "P": ";",

		"A": "A", "S": "S", "D": "H", "F": "T", "G": "G",
		"H": "Y", "J": "N", "K": "E", "L": "O",
		";": "I",

		"Z": "Z", "X": "X", "C": "M", "V": "C",
		"B": "V", "N": "K", "M": "L",
	},

	"qwerty uk": nil,

	"azerty": {
		"`": "²",

		"1": "&", "2": "É", "3": "\"", "4": "'",
		"5": "(", "6": "-", "7": "È", "8": "_",
		"9": "Ç", "0": "À", "-": ")",

		"Q": "A", "W": "Z", "E": "E", "R": "R", "T": "T",
		"Y": "Y", "U": "U", "I": "I", "O": "O", "P": "P",
		"[": "^", "]": "$",

		"A": "Q", "S": "S", "D": "D", "F": "F", "G": "G",
		"H": "H", "J": "J", "K": "K", "L": "L",
		";": "M", "'": "Ù",

		"Z": "W", "X": "X", "C": "C", "V": "V",
		"B": "B", "N": "N", "M": ",",
		",": ";", ".": ":", "/": "!",
	},
}

var shiftMaps = map[string]map[string]string{
	"qwerty uk": base.UKShift,
	"dvorak uk": base.UKShift,
	"azerty":    base.AZERTYShift,
}

var altGrMaps = map[string]map[string]string{
	"qwerty uk": base.UKAltGr,
	"dvorak uk": base.UKAltGr,
}

var BuiltinLayoutNames = []string{
	"qwerty",
	"qwerty uk",
	"dvorak",
	"dvorak uk",
	"colemak",
	"colemak-dh",
	"workman",
	"azerty",
}

var Sizes = []int{60, 65, 75, 80, 96, 100}

var LayoutListItems []string

var LayoutSizeItems []string

func init() {
	loadCustomLayouts()

	LayoutSizeItems = make([]string, len(Sizes))
	for i, s := range Sizes {
		LayoutSizeItems[i] = fmt.Sprintf("%d%%", s)
	}

	LayoutListItems = make([]string, 0, len(BuiltinLayoutNames)+len(customLayoutNames))
	for _, name := range BuiltinLayoutNames {
		LayoutListItems = append(LayoutListItems, name)
	}
	for _, name := range customLayoutNames {
		LayoutListItems = append(LayoutListItems, name)
	}
}
