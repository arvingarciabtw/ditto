package tui

import (
	"image/color"
	"os"

	lipgloss "charm.land/lipgloss/v2"

	"github.com/arvingarciabtw/ditto/internal/keyboard"
)

var (
	LayoutColor     color.Color
	SizeColor       color.Color
	StandardColor   color.Color
	QuitColor       color.Color
	QuitBorderColor color.Color
	ModeColor color.Color
	OverlayBase     lipgloss.Style
	StatusBarStyle  lipgloss.Style
	FingerStyle     map[keyboard.Finger]lipgloss.Style
	FingerActive    map[keyboard.Finger]lipgloss.Style
	WarningStyle    lipgloss.Style
	WarningAccent   lipgloss.Style
)

var KeyboardBorder = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     ",",
	TopRight:    ",",
	BottomLeft:  "'",
	BottomRight: "'",
}

var KeyboardBorderBoxDraw = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     "╭",
	TopRight:    "╮",
	BottomLeft:  "╰",
	BottomRight: "╯",
}

var darkColors = map[keyboard.Finger]color.Color{
	keyboard.Pinky:  lipgloss.BrightMagenta,
	keyboard.Ring:   lipgloss.BrightBlue,
	keyboard.Middle: lipgloss.BrightGreen,
	keyboard.Index:  lipgloss.BrightYellow,
	keyboard.Thumb:  lipgloss.BrightCyan,
	keyboard.Any:    lipgloss.BrightRed,
}

var lightColors = map[keyboard.Finger]color.Color{
	keyboard.Pinky:  lipgloss.Magenta,
	keyboard.Ring:   lipgloss.Blue,
	keyboard.Middle: lipgloss.Green,
	keyboard.Index:  lipgloss.Yellow,
	keyboard.Thumb:  lipgloss.Cyan,
	keyboard.Any:    lipgloss.Red,
}

func init() {
	isDark := lipgloss.HasDarkBackground(os.Stdin, os.Stdout)

	colors := darkColors

	if !isDark {
		colors = lightColors
	}

	OverlayBase = lipgloss.NewStyle().
		Border(KeyboardBorder).
		Padding(1, 3)

	if isDark {
		LayoutColor = lipgloss.BrightBlue
		SizeColor = lipgloss.BrightMagenta
		StandardColor = lipgloss.BrightYellow
		QuitColor = lipgloss.BrightRed
		QuitBorderColor = lipgloss.BrightBlack
		ModeColor = lipgloss.BrightCyan
	} else {
		LayoutColor = lipgloss.Blue
		SizeColor = lipgloss.Magenta
		StandardColor = lipgloss.Yellow
		QuitColor = lipgloss.Red
		QuitBorderColor = lipgloss.Black
		ModeColor = lipgloss.Cyan
	}

	FingerStyle = make(map[keyboard.Finger]lipgloss.Style, len(colors))
	FingerActive = make(map[keyboard.Finger]lipgloss.Style, len(colors))

	for fng, c := range colors {
		base := lipgloss.NewStyle().Foreground(c)

		if isDark {
			FingerStyle[fng] = base.Faint(true)
		} else {
			FingerStyle[fng] = base
		}

		FingerActive[fng] = base.Bold(true).Italic(true)
	}

	if isDark {
		StatusBarStyle = lipgloss.NewStyle().Foreground(lipgloss.BrightBlack)
		WarningAccent = lipgloss.NewStyle().Foreground(lipgloss.BrightGreen)
	} else {
		StatusBarStyle = lipgloss.NewStyle().Faint(true)
		WarningAccent = lipgloss.NewStyle().Foreground(lipgloss.Green)
	}

	WarningStyle = lipgloss.NewStyle()
}
