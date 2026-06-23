package keyboard

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
	"qwerty":     usShift,
	"dvorak":     usShift,
	"colemak":    usShift,
	"colemak-dh": usShift,
	"workman":    usShift,
	"azerty":     azertyShift,
}

var usShift = map[string]string{
	"`": "~", "1": "!", "2": "@", "3": "#", "4": "$",
	"5": "%", "6": "^", "7": "&", "8": "*", "9": "(", "0": ")",
	"-": "_", "=": "+", "[": "{", "]": "}", "\\": "|",
	";": ":", "'": "\"", ",": "<", ".": ">", "/": "?",
}

var azertyShift = map[string]string{
	"²": "~", "&": "1", "É": "2", "\"": "3", "'": "4",
	"(": "5", "-": "6", "È": "7", "_": "8",
	"Ç": "9", "À": "0", ")": "°",
	"^": "¨", "$": "£",
}

var BuiltinLayoutNames = []string{
	"qwerty",
	"dvorak",
	"colemak",
	"colemak-dh",
	"workman",
	"azerty",
}

var LayoutListItems []string

var LayoutSizeItems = []string{
	"60%",
	"65%",
	"75%",
	"80%",
	"96%",
	"100%",
}

func init() {
	seen := make(map[string]bool, len(BuiltinLayoutNames))
	LayoutListItems = make([]string, 0, len(layouts))
	for _, name := range BuiltinLayoutNames {
		LayoutListItems = append(LayoutListItems, name)
		seen[name] = true
	}
	for name := range layouts {
		if !seen[name] {
			LayoutListItems = append(LayoutListItems, name)
			seen[name] = true
		}
	}
}
