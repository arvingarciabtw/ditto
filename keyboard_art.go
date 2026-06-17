package main

import (
	"strings"

	lipgloss "charm.land/lipgloss/v2"
)

type Finger int

const (
	FingerPinky Finger = iota
	FingerRing
	FingerMiddle
	FingerIndex
	FingerThumb
)

type KeyDef struct {
	Label     string
	Width     int
	Finger    Finger
	Gap       bool
	Rightless bool
	Leftless  bool
	DivLabel  string
}

var fingerStyle = map[Finger]lipgloss.Style{
	FingerPinky:  lipgloss.NewStyle().Foreground(lipgloss.BrightMagenta),
	FingerRing:   lipgloss.NewStyle().Foreground(lipgloss.BrightRed),
	FingerMiddle: lipgloss.NewStyle().Foreground(lipgloss.BrightYellow),
	FingerIndex:  lipgloss.NewStyle().Foreground(lipgloss.BrightCyan),
	FingerThumb:  lipgloss.NewStyle().Foreground(lipgloss.BrightGreen),
}

var keyboardSizes = map[int][][]KeyDef{}

var size60 = [][]KeyDef{
	// Row 0: Number row
	{
		{Label: "⎋", Width: 3, Finger: FingerPinky},
		{Label: "1", Width: 3, Finger: FingerPinky},
		{Label: "2", Width: 3, Finger: FingerRing},
		{Label: "3", Width: 3, Finger: FingerMiddle},
		{Label: "4", Width: 3, Finger: FingerIndex},
		{Label: "5", Width: 3, Finger: FingerIndex},
		{Label: "6", Width: 3, Finger: FingerIndex},
		{Label: "7", Width: 3, Finger: FingerIndex},
		{Label: "8", Width: 3, Finger: FingerMiddle},
		{Label: "9", Width: 3, Finger: FingerRing},
		{Label: "0", Width: 3, Finger: FingerPinky},
		{Label: "-", Width: 3, Finger: FingerPinky},
		{Label: "=", Width: 3, Finger: FingerPinky},
		{Label: "<--", Width: 8, Finger: FingerPinky},
	},
	// Row 1: Top alpha
	{
		{Label: "Tab⇄", Width: 6, Finger: FingerPinky},
		{Label: "Q", Width: 3, Finger: FingerPinky},
		{Label: "W", Width: 3, Finger: FingerRing},
		{Label: "E", Width: 3, Finger: FingerMiddle},
		{Label: "R", Width: 3, Finger: FingerIndex},
		{Label: "T", Width: 3, Finger: FingerIndex},
		{Label: "Y", Width: 3, Finger: FingerIndex},
		{Label: "U", Width: 3, Finger: FingerIndex},
		{Label: "I", Width: 3, Finger: FingerMiddle},
		{Label: "O", Width: 3, Finger: FingerRing},
		{Label: "P", Width: 3, Finger: FingerPinky},
		{Label: "[", Width: 3, Finger: FingerPinky},
		{Label: "]", Width: 3, Finger: FingerPinky},
		{Label: "\\", Width: 5, Finger: FingerPinky},
	},
	// Row 2: Home row
	{
		{Label: "Caps", Width: 6, Finger: FingerPinky},
		{Label: "A", Width: 3, Finger: FingerPinky},
		{Label: "S", Width: 3, Finger: FingerRing},
		{Label: "D", Width: 3, Finger: FingerMiddle},
		{Label: "F", Width: 3, Finger: FingerIndex},
		{Label: "G", Width: 3, Finger: FingerIndex},
		{Label: "H", Width: 3, Finger: FingerIndex},
		{Label: "J", Width: 3, Finger: FingerIndex},
		{Label: "K", Width: 3, Finger: FingerMiddle},
		{Label: "L", Width: 3, Finger: FingerRing},
		{Label: ";", Width: 3, Finger: FingerPinky},
		{Label: "'", Width: 3, Finger: FingerPinky},
		{Label: "Enter↵", Width: 9, Finger: FingerPinky},
	},
	// Row 3: Bottom alpha
	{
		{Label: "Shift", Width: 10, Finger: FingerPinky},
		{Label: "Z", Width: 3, Finger: FingerPinky},
		{Label: "X", Width: 3, Finger: FingerRing},
		{Label: "C", Width: 3, Finger: FingerMiddle},
		{Label: "V", Width: 3, Finger: FingerIndex},
		{Label: "B", Width: 3, Finger: FingerIndex},
		{Label: "N", Width: 3, Finger: FingerIndex},
		{Label: "M", Width: 3, Finger: FingerIndex},
		{Label: ",", Width: 3, Finger: FingerMiddle},
		{Label: ".", Width: 3, Finger: FingerRing},
		{Label: "/", Width: 3, Finger: FingerPinky},
		{Label: "Shift", Width: 9, Finger: FingerPinky},
	},
	// Row 4: Modifiers
	{
		{Label: "Ctrl", Width: 6, Finger: FingerPinky},
		{Label: "⌘", Width: 3, Finger: FingerThumb},
		{Label: "Alt", Width: 5, Finger: FingerThumb},
		{Label: "Space", Width: 22, Finger: FingerThumb},
		{Label: "Alt", Width: 5, Finger: FingerThumb},
		{Label: "⌘", Width: 3, Finger: FingerThumb},
		{Label: "Fn", Width: 3, Finger: FingerThumb},
		{Label: "Ctrl", Width: 6, Finger: FingerPinky},
	},
}

var size65 = [][]KeyDef{
	// Row 0: Number row + Del
	{
		{Label: "⎋", Width: 3, Finger: FingerPinky},
		{Label: "1", Width: 3, Finger: FingerPinky},
		{Label: "2", Width: 3, Finger: FingerRing},
		{Label: "3", Width: 3, Finger: FingerMiddle},
		{Label: "4", Width: 3, Finger: FingerIndex},
		{Label: "5", Width: 3, Finger: FingerIndex},
		{Label: "6", Width: 3, Finger: FingerIndex},
		{Label: "7", Width: 3, Finger: FingerIndex},
		{Label: "8", Width: 3, Finger: FingerMiddle},
		{Label: "9", Width: 3, Finger: FingerRing},
		{Label: "0", Width: 3, Finger: FingerPinky},
		{Label: "-", Width: 3, Finger: FingerPinky},
		{Label: "=", Width: 3, Finger: FingerPinky},
		{Label: "<--", Width: 8, Finger: FingerPinky},
		{Label: "☼", Width: 3, Finger: FingerPinky},
	},
	// Row 1: Top alpha + PgUp
	{
		{Label: "Tab⇄", Width: 6, Finger: FingerPinky},
		{Label: "Q", Width: 3, Finger: FingerPinky},
		{Label: "W", Width: 3, Finger: FingerRing},
		{Label: "E", Width: 3, Finger: FingerMiddle},
		{Label: "R", Width: 3, Finger: FingerIndex},
		{Label: "T", Width: 3, Finger: FingerIndex},
		{Label: "Y", Width: 3, Finger: FingerIndex},
		{Label: "U", Width: 3, Finger: FingerIndex},
		{Label: "I", Width: 3, Finger: FingerMiddle},
		{Label: "O", Width: 3, Finger: FingerRing},
		{Label: "P", Width: 3, Finger: FingerPinky},
		{Label: "[", Width: 3, Finger: FingerPinky},
		{Label: "]", Width: 3, Finger: FingerPinky},
		{Label: "\\", Width: 5, Finger: FingerPinky},
		{Label: "⌂", Width: 3, Finger: FingerPinky},
	},
	// Row 2: Home row + PgDn
	{
		{Label: "Caps", Width: 6, Finger: FingerPinky},
		{Label: "A", Width: 3, Finger: FingerPinky},
		{Label: "S", Width: 3, Finger: FingerRing},
		{Label: "D", Width: 3, Finger: FingerMiddle},
		{Label: "F", Width: 3, Finger: FingerIndex},
		{Label: "G", Width: 3, Finger: FingerIndex},
		{Label: "H", Width: 3, Finger: FingerIndex},
		{Label: "J", Width: 3, Finger: FingerIndex},
		{Label: "K", Width: 3, Finger: FingerMiddle},
		{Label: "L", Width: 3, Finger: FingerRing},
		{Label: ";", Width: 3, Finger: FingerPinky},
		{Label: "'", Width: 3, Finger: FingerPinky},
		{Label: "Enter↵", Width: 9, Finger: FingerPinky},
		{Label: "⇡", Width: 3, Finger: FingerPinky},
	},
	// Row 3: Bottom alpha + Home + End
	{
		{Label: "Shift", Width: 8, Finger: FingerPinky},
		{Label: "Z", Width: 3, Finger: FingerPinky},
		{Label: "X", Width: 3, Finger: FingerRing},
		{Label: "C", Width: 3, Finger: FingerMiddle},
		{Label: "V", Width: 3, Finger: FingerIndex},
		{Label: "B", Width: 3, Finger: FingerIndex},
		{Label: "N", Width: 3, Finger: FingerIndex},
		{Label: "M", Width: 3, Finger: FingerIndex},
		{Label: ",", Width: 3, Finger: FingerMiddle},
		{Label: ".", Width: 3, Finger: FingerRing},
		{Label: "/", Width: 3, Finger: FingerPinky},
		{Label: "Shift", Width: 7, Finger: FingerPinky},
		{Label: "↑", Width: 3, Finger: FingerPinky},
		{Label: "⇣", Width: 3, Finger: FingerPinky},
	},
	// Row 4: Modifiers + arrows
	{
		{Label: "Ctrl", Width: 6, Finger: FingerPinky},
		{Label: "⌘", Width: 3, Finger: FingerThumb},
		{Label: "Alt", Width: 3, Finger: FingerThumb},
		{Label: "Space", Width: 22, Finger: FingerThumb},
		{Label: "Alt", Width: 3, Finger: FingerThumb},
		{Label: "Fn", Width: 3, Finger: FingerThumb},
		{Label: "Ctrl", Width: 6, Finger: FingerPinky},
		{Label: "←", Width: 3, Finger: FingerPinky},
		{Label: "↓", Width: 3, Finger: FingerPinky},
		{Label: "→", Width: 3, Finger: FingerPinky},
	},
}

var size75 = [][]KeyDef{
	// Row 0: Esc + F1-F12 + Del
	{
		{Label: "⎋", Width: 3, Finger: FingerPinky},
		{Label: "F1", Width: 3, Finger: FingerPinky},
		{Label: "F2", Width: 3, Finger: FingerRing},
		{Label: "F3", Width: 3, Finger: FingerMiddle},
		{Label: "F4", Width: 3, Finger: FingerIndex},
		{Label: "F5", Width: 3, Finger: FingerIndex},
		{Label: "F6", Width: 3, Finger: FingerIndex},
		{Label: "F7", Width: 3, Finger: FingerIndex},
		{Label: "F8", Width: 3, Finger: FingerMiddle},
		{Label: "F9", Width: 3, Finger: FingerRing},
		{Label: "F10", Width: 3, Finger: FingerPinky},
		{Label: "F11", Width: 3, Finger: FingerPinky},
		{Label: "F12", Width: 3, Finger: FingerPinky},
		{Label: "⎚", Width: 3, Finger: FingerPinky},
		{Label: "Del", Width: 4, Finger: FingerPinky},
		{Label: "☼", Width: 3, Finger: FingerPinky},
	},
	// Row 1: Number row
	{
		{Label: "`", Width: 3, Finger: FingerPinky},
		{Label: "1", Width: 3, Finger: FingerPinky},
		{Label: "2", Width: 3, Finger: FingerRing},
		{Label: "3", Width: 3, Finger: FingerMiddle},
		{Label: "4", Width: 3, Finger: FingerIndex},
		{Label: "5", Width: 3, Finger: FingerIndex},
		{Label: "6", Width: 3, Finger: FingerIndex},
		{Label: "7", Width: 3, Finger: FingerIndex},
		{Label: "8", Width: 3, Finger: FingerMiddle},
		{Label: "9", Width: 3, Finger: FingerRing},
		{Label: "0", Width: 3, Finger: FingerPinky},
		{Label: "-", Width: 3, Finger: FingerPinky},
		{Label: "=", Width: 3, Finger: FingerPinky},
		{Label: "<--", Width: 8, Finger: FingerPinky},
		{Label: "⇡", Width: 3, Finger: FingerPinky},
	},
	// Row 2: Top alpha
	{
		{Label: "Tab⇄", Width: 6, Finger: FingerPinky},
		{Label: "Q", Width: 3, Finger: FingerPinky},
		{Label: "W", Width: 3, Finger: FingerRing},
		{Label: "E", Width: 3, Finger: FingerMiddle},
		{Label: "R", Width: 3, Finger: FingerIndex},
		{Label: "T", Width: 3, Finger: FingerIndex},
		{Label: "Y", Width: 3, Finger: FingerIndex},
		{Label: "U", Width: 3, Finger: FingerIndex},
		{Label: "I", Width: 3, Finger: FingerMiddle},
		{Label: "O", Width: 3, Finger: FingerRing},
		{Label: "P", Width: 3, Finger: FingerPinky},
		{Label: "[", Width: 3, Finger: FingerPinky},
		{Label: "]", Width: 3, Finger: FingerPinky},
		{Label: "\\", Width: 5, Finger: FingerPinky},
		{Label: "⇣", Width: 3, Finger: FingerPinky},
	},
	// Row 3: Home row
	{
		{Label: "Caps", Width: 6, Finger: FingerPinky},
		{Label: "A", Width: 3, Finger: FingerPinky},
		{Label: "S", Width: 3, Finger: FingerRing},
		{Label: "D", Width: 3, Finger: FingerMiddle},
		{Label: "F", Width: 3, Finger: FingerIndex},
		{Label: "G", Width: 3, Finger: FingerIndex},
		{Label: "H", Width: 3, Finger: FingerIndex},
		{Label: "J", Width: 3, Finger: FingerIndex},
		{Label: "K", Width: 3, Finger: FingerMiddle},
		{Label: "L", Width: 3, Finger: FingerRing},
		{Label: ";", Width: 3, Finger: FingerPinky},
		{Label: "'", Width: 3, Finger: FingerPinky},
		{Label: "Enter↵", Width: 9, Finger: FingerPinky},
		{Label: "⌂", Width: 3, Finger: FingerPinky},
	},
	// Row 4: Bottom alpha + up arrow
	{
		{Label: "Shift", Width: 8, Finger: FingerPinky},
		{Label: "Z", Width: 3, Finger: FingerPinky},
		{Label: "X", Width: 3, Finger: FingerRing},
		{Label: "C", Width: 3, Finger: FingerMiddle},
		{Label: "V", Width: 3, Finger: FingerIndex},
		{Label: "B", Width: 3, Finger: FingerIndex},
		{Label: "N", Width: 3, Finger: FingerIndex},
		{Label: "M", Width: 3, Finger: FingerIndex},
		{Label: ",", Width: 3, Finger: FingerMiddle},
		{Label: ".", Width: 3, Finger: FingerRing},
		{Label: "/", Width: 3, Finger: FingerPinky},
		{Label: "Shift", Width: 7, Finger: FingerPinky},
		{Label: "↑", Width: 3, Finger: FingerPinky},
		{Label: "⌿", Width: 3, Finger: FingerPinky},
	},
	// Row 5: Modifiers + arrows
	{
		{Label: "Ctrl", Width: 6, Finger: FingerPinky},
		{Label: "⌘", Width: 3, Finger: FingerThumb},
		{Label: "Alt", Width: 3, Finger: FingerThumb},
		{Label: "Space", Width: 22, Finger: FingerThumb},
		{Label: "Alt", Width: 3, Finger: FingerThumb},
		{Label: "Fn", Width: 3, Finger: FingerThumb},
		{Label: "Ctrl", Width: 6, Finger: FingerPinky},
		{Label: "←", Width: 3, Finger: FingerPinky},
		{Label: "↓", Width: 3, Finger: FingerPinky},
		{Label: "→", Width: 3, Finger: FingerPinky},
	},
}

var size80 = [][]KeyDef{
	// Row 0: Esc + F1-F12 + PrSc + ScLk + Pse
	{
		{Label: "⎋", Width: 3, Finger: FingerPinky},
		{Label: "   ", Width: 3, Finger: FingerPinky},
		{Label: "F1", Width: 3, Finger: FingerPinky},
		{Label: "F2", Width: 3, Finger: FingerRing},
		{Label: "F3", Width: 3, Finger: FingerMiddle},
		{Label: "F4", Width: 3, Finger: FingerIndex},
		{Label: "  ", Width: 2, Finger: FingerIndex},
		{Label: "F5", Width: 3, Finger: FingerIndex},
		{Label: "F6", Width: 3, Finger: FingerIndex},
		{Label: "F7", Width: 3, Finger: FingerIndex},
		{Label: "F8", Width: 3, Finger: FingerMiddle},
		{Label: "  ", Width: 2, Finger: FingerIndex},
		{Label: "F9", Width: 3, Finger: FingerRing},
		{Label: "F10", Width: 3, Finger: FingerPinky},
		{Label: "F11", Width: 3, Finger: FingerPinky},
		{Label: "F12", Width: 3, Finger: FingerPinky},
		{Label: "  ", Width: 2, Finger: FingerIndex, Gap: true},
		{Label: "⎚", Width: 3, Finger: FingerPinky},
		{Label: "⚲", Width: 3, Finger: FingerPinky},
		{Label: "☼", Width: 3, Finger: FingerPinky},
	},
	// Row 1: Number row + Ins + Hme + PgUp
	{
		{Label: "`", Width: 3, Finger: FingerPinky},
		{Label: "1", Width: 3, Finger: FingerPinky},
		{Label: "2", Width: 3, Finger: FingerRing},
		{Label: "3", Width: 3, Finger: FingerMiddle},
		{Label: "4", Width: 3, Finger: FingerIndex},
		{Label: "5", Width: 3, Finger: FingerIndex},
		{Label: "6", Width: 3, Finger: FingerIndex},
		{Label: "7", Width: 3, Finger: FingerIndex},
		{Label: "8", Width: 3, Finger: FingerMiddle},
		{Label: "9", Width: 3, Finger: FingerRing},
		{Label: "0", Width: 3, Finger: FingerPinky},
		{Label: "-", Width: 3, Finger: FingerPinky},
		{Label: "=", Width: 3, Finger: FingerPinky},
		{Label: "<--", Width: 9, Finger: FingerPinky},
		{Label: "  ", Width: 2, Finger: FingerPinky, Gap: true},
		{Label: "Ins", Width: 3, Finger: FingerPinky},
		{Label: "⌂", Width: 3, Finger: FingerPinky},
		{Label: "⇡", Width: 3, Finger: FingerPinky},
	},
	// Row 2: Top alpha + Del + End + PgDn
	{
		{Label: "Tab⇄", Width: 7, Finger: FingerPinky},
		{Label: "Q", Width: 3, Finger: FingerPinky},
		{Label: "W", Width: 3, Finger: FingerRing},
		{Label: "E", Width: 3, Finger: FingerMiddle},
		{Label: "R", Width: 3, Finger: FingerIndex},
		{Label: "T", Width: 3, Finger: FingerIndex},
		{Label: "Y", Width: 3, Finger: FingerIndex},
		{Label: "U", Width: 3, Finger: FingerIndex},
		{Label: "I", Width: 3, Finger: FingerMiddle},
		{Label: "O", Width: 3, Finger: FingerRing},
		{Label: "P", Width: 3, Finger: FingerPinky},
		{Label: "[", Width: 3, Finger: FingerPinky},
		{Label: "]", Width: 3, Finger: FingerPinky},
		{Label: "\\", Width: 5, Finger: FingerPinky},
		{Label: "  ", Width: 2, Finger: FingerPinky, Gap: true},
		{Label: "Del", Width: 3, Finger: FingerPinky},
		{Label: "⌿", Width: 3, Finger: FingerPinky},
		{Label: "⇣", Width: 3, Finger: FingerPinky},
	},
	// Row 3: Home row
	{
		{Label: "Caps", Width: 6, Finger: FingerPinky},
		{Label: "A", Width: 3, Finger: FingerPinky},
		{Label: "S", Width: 3, Finger: FingerRing},
		{Label: "D", Width: 3, Finger: FingerMiddle},
		{Label: "F", Width: 3, Finger: FingerIndex},
		{Label: "G", Width: 3, Finger: FingerIndex},
		{Label: "H", Width: 3, Finger: FingerIndex},
		{Label: "J", Width: 3, Finger: FingerIndex},
		{Label: "K", Width: 3, Finger: FingerMiddle},
		{Label: "L", Width: 3, Finger: FingerRing},
		{Label: ";", Width: 3, Finger: FingerPinky},
		{Label: "'", Width: 3, Finger: FingerPinky},
		{Label: "Enter↵", Width: 10, Finger: FingerPinky},
		{Label: "      ", Width: 6, Finger: FingerPinky, Gap: true, Rightless: true, Leftless: false},
		{Label: "   ", Width: 3, Finger: FingerPinky, Gap: false, Rightless: true, Leftless: true},
		{Label: "   ", Width: 3, Finger: FingerPinky, Gap: true, Rightless: false, Leftless: false},
	},
	// Row 4: Bottom alpha + ↑
	{
		{Label: "Shift", Width: 11, Finger: FingerPinky},
		{Label: "Z", Width: 3, Finger: FingerPinky},
		{Label: "X", Width: 3, Finger: FingerRing},
		{Label: "C", Width: 3, Finger: FingerMiddle},
		{Label: "V", Width: 3, Finger: FingerIndex},
		{Label: "B", Width: 3, Finger: FingerIndex},
		{Label: "N", Width: 3, Finger: FingerIndex},
		{Label: "M", Width: 3, Finger: FingerIndex},
		{Label: ",", Width: 3, Finger: FingerMiddle},
		{Label: ".", Width: 3, Finger: FingerRing},
		{Label: "/", Width: 3, Finger: FingerPinky},
		{Label: "Shift", Width: 9, Finger: FingerPinky},
		{Label: "  ", Width: 2, Finger: FingerPinky, Gap: true, Rightless: true, Leftless: false},
		{Label: "   ", Width: 3, Finger: FingerPinky},
		{Label: "↑", Width: 3, Finger: FingerPinky},
		{Label: "   ", Width: 3, Finger: FingerPinky},
	},
	// Row 5: Modifiers + arrows
	{
		{Label: "Ctrl", Width: 6, Finger: FingerPinky},
		{Label: "⌘", Width: 3, Finger: FingerThumb},
		{Label: "Alt", Width: 5, Finger: FingerThumb},
		{Label: "Space", Width: 23, Finger: FingerThumb},
		{Label: "Alt", Width: 5, Finger: FingerThumb},
		{Label: "⌘", Width: 3, Finger: FingerThumb},
		{Label: "Fn", Width: 3, Finger: FingerThumb},
		{Label: "Ctrl", Width: 6, Finger: FingerPinky},
		{Label: "  ", Width: 2, Finger: FingerPinky},
		{Label: "←", Width: 3, Finger: FingerPinky},
		{Label: "↓", Width: 3, Finger: FingerPinky},
		{Label: "→", Width: 3, Finger: FingerPinky},
	},
}

var size96 = [][]KeyDef{
	// Row 0: Esc + F1-F12 + PrSc
	{
		{Label: "⎋", Width: 3, Finger: FingerPinky},
		{Label: "F1", Width: 3, Finger: FingerPinky},
		{Label: "F2", Width: 3, Finger: FingerRing},
		{Label: "F3", Width: 3, Finger: FingerMiddle},
		{Label: "F4", Width: 3, Finger: FingerIndex},
		{Label: "F5", Width: 3, Finger: FingerIndex},
		{Label: "F6", Width: 3, Finger: FingerIndex},
		{Label: "F7", Width: 3, Finger: FingerIndex},
		{Label: "F8", Width: 3, Finger: FingerMiddle},
		{Label: "F9", Width: 3, Finger: FingerRing},
		{Label: "F10", Width: 3, Finger: FingerPinky},
		{Label: "F11", Width: 3, Finger: FingerPinky},
		{Label: "F12", Width: 3, Finger: FingerPinky},
		{Label: "Del", Width: 3, Finger: FingerPinky},
		{Label: "⌂", Width: 3, Finger: FingerPinky},
		{Label: "⌿", Width: 3, Finger: FingerPinky},
		{Label: "⇡", Width: 3, Finger: FingerPinky},
		{Label: "⇣", Width: 3, Finger: FingerPinky},
		{Label: "☼", Width: 3, Finger: FingerPinky},
	},
	// Row 1: Number row + Ins + Home + NmLk + / + * + -
	{
		{Label: "`", Width: 3, Finger: FingerPinky},
		{Label: "1", Width: 3, Finger: FingerPinky},
		{Label: "2", Width: 3, Finger: FingerRing},
		{Label: "3", Width: 3, Finger: FingerMiddle},
		{Label: "4", Width: 3, Finger: FingerIndex},
		{Label: "5", Width: 3, Finger: FingerIndex},
		{Label: "6", Width: 3, Finger: FingerIndex},
		{Label: "7", Width: 3, Finger: FingerIndex},
		{Label: "8", Width: 3, Finger: FingerMiddle},
		{Label: "9", Width: 3, Finger: FingerRing},
		{Label: "0", Width: 3, Finger: FingerPinky},
		{Label: "-", Width: 3, Finger: FingerPinky},
		{Label: "=", Width: 3, Finger: FingerPinky},
		{Label: "<--", Width: 7, Finger: FingerPinky},
		{Label: "Nlk", Width: 3, Finger: FingerPinky},
		{Label: "/", Width: 3, Finger: FingerPinky},
		{Label: "*", Width: 3, Finger: FingerPinky},
		{Label: "-", Width: 3, Finger: FingerPinky},
	},
	// Row 2: Top alpha + Del + End + 7 + 8 + 9 + +
	{
		{Label: "Tab ⇄", Width: 7, Finger: FingerPinky},
		{Label: "Q", Width: 3, Finger: FingerPinky},
		{Label: "W", Width: 3, Finger: FingerRing},
		{Label: "E", Width: 3, Finger: FingerMiddle},
		{Label: "R", Width: 3, Finger: FingerIndex},
		{Label: "T", Width: 3, Finger: FingerIndex},
		{Label: "Y", Width: 3, Finger: FingerIndex},
		{Label: "U", Width: 3, Finger: FingerIndex},
		{Label: "I", Width: 3, Finger: FingerMiddle},
		{Label: "O", Width: 3, Finger: FingerRing},
		{Label: "P", Width: 3, Finger: FingerPinky},
		{Label: "[", Width: 3, Finger: FingerPinky},
		{Label: "]", Width: 3, Finger: FingerPinky},
		{Label: "\\", Width: 3, Finger: FingerPinky},
		{Label: "7", Width: 3, Finger: FingerPinky},
		{Label: "8", Width: 3, Finger: FingerPinky},
		{Label: "9", Width: 3, Finger: FingerPinky},
		{Label: "   ", Width: 3, Finger: FingerPinky, Gap: true, DivLabel: "+"},
	},
	// Row 3: Home row + 4 + 5 + 6 (blank for + continuation)
	{
		{Label: "Caps", Width: 6, Finger: FingerPinky},
		{Label: "A", Width: 3, Finger: FingerPinky},
		{Label: "S", Width: 3, Finger: FingerRing},
		{Label: "D", Width: 3, Finger: FingerMiddle},
		{Label: "F", Width: 3, Finger: FingerIndex},
		{Label: "G", Width: 3, Finger: FingerIndex},
		{Label: "H", Width: 3, Finger: FingerIndex},
		{Label: "J", Width: 3, Finger: FingerIndex},
		{Label: "K", Width: 3, Finger: FingerMiddle},
		{Label: "L", Width: 3, Finger: FingerRing},
		{Label: ";", Width: 3, Finger: FingerPinky},
		{Label: "'", Width: 3, Finger: FingerPinky},
		{Label: "Enter↵", Width: 8, Finger: FingerPinky},
		{Label: "4", Width: 3, Finger: FingerPinky},
		{Label: "5", Width: 3, Finger: FingerPinky},
		{Label: "6", Width: 3, Finger: FingerPinky},
		{Label: "   ", Width: 3, Finger: FingerPinky},
	},
	// Row 4: Bottom alpha + ↑ + 1 + 2 + 3 (blank for Enter continuation)
	{
		{Label: "Shift", Width: 7, Finger: FingerPinky},
		{Label: "Z", Width: 3, Finger: FingerPinky},
		{Label: "X", Width: 3, Finger: FingerRing},
		{Label: "C", Width: 3, Finger: FingerMiddle},
		{Label: "V", Width: 3, Finger: FingerIndex},
		{Label: "B", Width: 3, Finger: FingerIndex},
		{Label: "N", Width: 3, Finger: FingerIndex},
		{Label: "M", Width: 3, Finger: FingerIndex},
		{Label: ",", Width: 3, Finger: FingerMiddle},
		{Label: ".", Width: 3, Finger: FingerRing},
		{Label: "/", Width: 3, Finger: FingerPinky},
		{Label: "Shift", Width: 7, Finger: FingerPinky},
		{Label: "↑", Width: 3, Finger: FingerPinky},
		{Label: "1", Width: 3, Finger: FingerPinky},
		{Label: "2", Width: 3, Finger: FingerPinky},
		{Label: "3", Width: 3, Finger: FingerPinky},
		{Label: "   ", Width: 3, Finger: FingerPinky, Gap: true, DivLabel: "↵"},
	},
	// Row 5: Modifiers + arrows + numpad bottom
	{
		{Label: "Ctrl", Width: 6, Finger: FingerPinky},
		{Label: "⌘", Width: 3, Finger: FingerThumb},
		{Label: "Alt", Width: 5, Finger: FingerThumb},
		{Label: "Space", Width: 17, Finger: FingerThumb},
		{Label: "Alt", Width: 5, Finger: FingerThumb},
		{Label: "Fn", Width: 3, Finger: FingerThumb},
		{Label: "Ctrl", Width: 6, Finger: FingerPinky},
		{Label: "←", Width: 3, Finger: FingerPinky},
		{Label: "↓", Width: 3, Finger: FingerPinky},
		{Label: "→", Width: 3, Finger: FingerPinky},
		{Label: "0", Width: 3, Finger: FingerPinky},
		{Label: ".", Width: 3, Finger: FingerPinky},
		{Label: "   ", Width: 3, Finger: FingerPinky},
	},
}

var size100 = [][]KeyDef{
	// Row 0: Esc + F1-F12 + PrSc + ScLk + Pse
	{
		{Label: "⎋", Width: 3, Finger: FingerPinky},
		{Label: "   ", Width: 3, Finger: FingerPinky},
		{Label: "F1", Width: 3, Finger: FingerPinky},
		{Label: "F2", Width: 3, Finger: FingerRing},
		{Label: "F3", Width: 3, Finger: FingerMiddle},
		{Label: "F4", Width: 3, Finger: FingerIndex},
		{Label: "  ", Width: 2, Finger: FingerIndex},
		{Label: "F5", Width: 3, Finger: FingerIndex},
		{Label: "F6", Width: 3, Finger: FingerIndex},
		{Label: "F7", Width: 3, Finger: FingerIndex},
		{Label: "F8", Width: 3, Finger: FingerMiddle},
		{Label: "  ", Width: 2, Finger: FingerIndex},
		{Label: "F9", Width: 3, Finger: FingerRing},
		{Label: "F10", Width: 3, Finger: FingerPinky},
		{Label: "F11", Width: 3, Finger: FingerPinky},
		{Label: "F12", Width: 3, Finger: FingerPinky},
		{Label: "  ", Width: 2, Finger: FingerIndex, Gap: true},
		{Label: "⎚", Width: 3, Finger: FingerPinky},
		{Label: "⚲", Width: 3, Finger: FingerPinky},
		{Label: "☼", Width: 3, Finger: FingerPinky},
		{Label: " ", Width: 2, Finger: FingerIndex, Gap: true, Rightless: true},
		{Label: " ", Width: 3, Finger: FingerIndex, Rightless: true, Leftless: true},
		{Label: " ", Width: 3, Finger: FingerIndex, Rightless: true, Leftless: true},
		{Label: " ", Width: 3, Finger: FingerIndex, Rightless: true, Leftless: true},
		{Label: " ", Width: 3, Finger: FingerIndex},
	},
	// Row 1: Number row + Ins + Hme + PgUp + NmLk + / + * + -
	{
		{Label: "`", Width: 3, Finger: FingerPinky},
		{Label: "1", Width: 3, Finger: FingerPinky},
		{Label: "2", Width: 3, Finger: FingerRing},
		{Label: "3", Width: 3, Finger: FingerMiddle},
		{Label: "4", Width: 3, Finger: FingerIndex},
		{Label: "5", Width: 3, Finger: FingerIndex},
		{Label: "6", Width: 3, Finger: FingerIndex},
		{Label: "7", Width: 3, Finger: FingerIndex},
		{Label: "8", Width: 3, Finger: FingerMiddle},
		{Label: "9", Width: 3, Finger: FingerRing},
		{Label: "0", Width: 3, Finger: FingerPinky},
		{Label: "-", Width: 3, Finger: FingerPinky},
		{Label: "=", Width: 3, Finger: FingerPinky},
		{Label: "<--", Width: 9, Finger: FingerPinky},
		{Label: "  ", Width: 2, Finger: FingerPinky, Gap: true},
		{Label: "Ins", Width: 3, Finger: FingerPinky},
		{Label: "⌂", Width: 3, Finger: FingerPinky},
		{Label: "⇡", Width: 3, Finger: FingerPinky},
		{Label: "  ", Width: 2, Finger: FingerPinky, Gap: true},
		{Label: "Nlk", Width: 3, Finger: FingerPinky},
		{Label: "/", Width: 3, Finger: FingerPinky},
		{Label: "*", Width: 3, Finger: FingerPinky},
		{Label: "-", Width: 3, Finger: FingerPinky},
	},
	// Row 2: Top alpha + Del + End + PgDn + 7 + 8 + 9 + +
	{
		{Label: "Tab ⇄", Width: 7, Finger: FingerPinky},
		{Label: "Q", Width: 3, Finger: FingerPinky},
		{Label: "W", Width: 3, Finger: FingerRing},
		{Label: "E", Width: 3, Finger: FingerMiddle},
		{Label: "R", Width: 3, Finger: FingerIndex},
		{Label: "T", Width: 3, Finger: FingerIndex},
		{Label: "Y", Width: 3, Finger: FingerIndex},
		{Label: "U", Width: 3, Finger: FingerIndex},
		{Label: "I", Width: 3, Finger: FingerMiddle},
		{Label: "O", Width: 3, Finger: FingerRing},
		{Label: "P", Width: 3, Finger: FingerPinky},
		{Label: "[", Width: 3, Finger: FingerPinky},
		{Label: "]", Width: 3, Finger: FingerPinky},
		{Label: "\\", Width: 5, Finger: FingerPinky},
		{Label: "  ", Width: 2, Finger: FingerPinky, Gap: true},
		{Label: "Del", Width: 3, Finger: FingerPinky},
		{Label: "⌿", Width: 3, Finger: FingerPinky},
		{Label: "⇣", Width: 3, Finger: FingerPinky},
		{Label: "  ", Width: 2, Finger: FingerPinky, Gap: true},
		{Label: "7", Width: 3, Finger: FingerPinky},
		{Label: "8", Width: 3, Finger: FingerPinky},
		{Label: "9", Width: 3, Finger: FingerPinky},
		{Label: "   ", Width: 3, Finger: FingerPinky, Gap: true, DivLabel: "+"},
	},
	// Row 3: Home row + gap + 4 + 5 + 6 (blank for + continuation)
	{
		{Label: "Caps", Width: 6, Finger: FingerPinky},
		{Label: "A", Width: 3, Finger: FingerPinky},
		{Label: "S", Width: 3, Finger: FingerRing},
		{Label: "D", Width: 3, Finger: FingerMiddle},
		{Label: "F", Width: 3, Finger: FingerIndex},
		{Label: "G", Width: 3, Finger: FingerIndex},
		{Label: "H", Width: 3, Finger: FingerIndex},
		{Label: "J", Width: 3, Finger: FingerIndex},
		{Label: "K", Width: 3, Finger: FingerMiddle},
		{Label: "L", Width: 3, Finger: FingerRing},
		{Label: ";", Width: 3, Finger: FingerPinky},
		{Label: "'", Width: 3, Finger: FingerPinky},
		{Label: "Enter↵", Width: 10, Finger: FingerPinky},
		{Label: " ", Width: 6, Finger: FingerPinky, Gap: true, Rightless: true, Leftless: false},
		{Label: " ", Width: 3, Finger: FingerPinky, Gap: false, Rightless: true, Leftless: true},
		{Label: " ", Width: 6, Finger: FingerPinky, Gap: true, Rightless: false, Leftless: false},
		{Label: "4", Width: 3, Finger: FingerPinky},
		{Label: "5", Width: 3, Finger: FingerPinky},
		{Label: "6", Width: 3, Finger: FingerPinky},
		{Label: "   ", Width: 3, Finger: FingerPinky},
	},
	// Row 4: Bottom alpha + ↑ + gap + 1 + 2 + 3 + (Enter blank)
	{
		{Label: "Shift", Width: 11, Finger: FingerPinky},
		{Label: "Z", Width: 3, Finger: FingerPinky},
		{Label: "X", Width: 3, Finger: FingerRing},
		{Label: "C", Width: 3, Finger: FingerMiddle},
		{Label: "V", Width: 3, Finger: FingerIndex},
		{Label: "B", Width: 3, Finger: FingerIndex},
		{Label: "N", Width: 3, Finger: FingerIndex},
		{Label: "M", Width: 3, Finger: FingerIndex},
		{Label: ",", Width: 3, Finger: FingerMiddle},
		{Label: ".", Width: 3, Finger: FingerRing},
		{Label: "/", Width: 3, Finger: FingerPinky},
		{Label: "Shift", Width: 9, Finger: FingerPinky},
		{Label: " ", Width: 2, Finger: FingerPinky, Gap: true, Rightless: true, Leftless: false},
		{Label: " ", Width: 3, Finger: FingerPinky},
		{Label: "↑", Width: 3, Finger: FingerPinky},
		{Label: " ", Width: 3, Finger: FingerPinky, Rightless: true, Leftless: true},
		{Label: " ", Width: 2, Finger: FingerPinky, Gap: true},
		{Label: "1", Width: 3, Finger: FingerPinky},
		{Label: "2", Width: 3, Finger: FingerPinky},
		{Label: "3", Width: 3, Finger: FingerPinky},
		{Label: " ", Width: 3, Finger: FingerPinky, Gap: true, DivLabel: "↵"},
	},
	// Row 5: Modifiers + arrows + gap + 0 + . + (Enter blank)
	{
		{Label: "Ctrl", Width: 6, Finger: FingerPinky},
		{Label: "⌘", Width: 3, Finger: FingerThumb},
		{Label: "Alt", Width: 5, Finger: FingerThumb},
		{Label: "Space", Width: 23, Finger: FingerThumb},
		{Label: "Alt", Width: 5, Finger: FingerThumb},
		{Label: "⌘", Width: 3, Finger: FingerThumb},
		{Label: "Fn", Width: 3, Finger: FingerThumb},
		{Label: "Ctrl", Width: 6, Finger: FingerPinky},
		{Label: "  ", Width: 2, Finger: FingerPinky},
		{Label: "←", Width: 3, Finger: FingerPinky},
		{Label: "↓", Width: 3, Finger: FingerPinky},
		{Label: "→", Width: 3, Finger: FingerPinky},
		{Label: "  ", Width: 2, Finger: FingerPinky, Gap: true},
		{Label: "0", Width: 7, Finger: FingerPinky},
		{Label: ".", Width: 3, Finger: FingerPinky},
		{Label: "   ", Width: 3, Finger: FingerPinky},
	},
}

func init() {
	keyboardSizes[60] = size60
	keyboardSizes[65] = size65
	keyboardSizes[75] = size75
	keyboardSizes[80] = size80
	keyboardSizes[96] = size96
	keyboardSizes[100] = size100
}

func renderKeyboardArt(size int) string {
	rows, ok := keyboardSizes[size]
	if !ok {
		return ""
	}
	var lines []string
	for i, row := range rows {
		if i == 0 {
			lines = append(lines, renderTopLine(row))
		}
		lines = append(lines, renderMidLine(row))
		if i < len(rows)-1 {
			lines = append(lines, renderDivLine(row))
		} else {
			lines = append(lines, renderBotLine(row))
		}
	}
	return strings.Join(lines, "\n")
}

func renderTopLine(keys []KeyDef) string {
	var b strings.Builder
	b.WriteByte(',')
	for _, k := range keys {
		b.WriteString(strings.Repeat("-", k.Width))
		b.WriteByte(',')
	}
	return b.String()
}

func renderMidLine(keys []KeyDef) string {
	var b strings.Builder
	b.WriteByte('|')
	for _, k := range keys {
		label := k.Label
		if k.DivLabel != "" {
			label = ""
		}
		b.WriteString(fingerStyle[k.Finger].Render(centerLabel(label, k.Width)))
		if k.Rightless {
			b.WriteByte(' ')
			continue
		}
		b.WriteByte('|')
	}
	return b.String()
}

func renderDivLine(keys []KeyDef) string {
	var b strings.Builder
	b.WriteByte('|')
	for _, k := range keys {
		if k.Gap {
			if k.DivLabel != "" {
				b.WriteString(fingerStyle[k.Finger].Render(centerLabel(k.DivLabel, k.Width)))
			} else {
				b.WriteString(strings.Repeat(" ", k.Width))
			}
			if k.Rightless {
				b.WriteByte(',')
			} else {
				b.WriteByte('\'')
			}
			continue
		}
		b.WriteString(strings.Repeat("-", k.Width))
		if k.Leftless {
			b.WriteByte(',')
		} else {
			b.WriteByte('\'')
		}
	}
	return b.String()
}

func renderBotLine(keys []KeyDef) string {
	var b strings.Builder
	b.WriteByte('\'')
	for _, k := range keys {
		b.WriteString(strings.Repeat("-", k.Width))
		if k.Leftless {
			b.WriteByte(',')
		} else {
			b.WriteByte('\'')
		}
	}
	return b.String()
}
