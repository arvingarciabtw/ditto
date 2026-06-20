package main

import (
	"image/color"
	"os"

	lipgloss "charm.land/lipgloss/v2"
)

var (
	layoutColor     color.Color
	sizeColor       color.Color
	quitColor       color.Color
	quitBorderColor color.Color
	overlayBase     lipgloss.Style
	statusBarStyle  lipgloss.Style
	fingerStyle     map[finger]lipgloss.Style
	fingerActive    map[finger]lipgloss.Style
)

var keyboardBorder = lipgloss.Border{
	Top:         "─",
	Bottom:      "─",
	Left:        "│",
	Right:       "│",
	TopLeft:     ",",
	TopRight:    ",",
	BottomLeft:  "'",
	BottomRight: "'",
}

var darkColors = map[finger]color.Color{
	pinky:  lipgloss.BrightMagenta,
	ring:   lipgloss.BrightBlue,
	middle: lipgloss.BrightGreen,
	index:  lipgloss.BrightYellow,
	thumb:  lipgloss.BrightCyan,
	any:    lipgloss.BrightRed,
}

var lightColors = map[finger]color.Color{
	pinky:  lipgloss.Magenta,
	ring:   lipgloss.Blue,
	middle: lipgloss.Green,
	index:  lipgloss.Yellow,
	thumb:  lipgloss.Cyan,
	any:    lipgloss.Red,
}

func init() {
	isDark := lipgloss.HasDarkBackground(os.Stdin, os.Stdout)

	colors := darkColors

	if !isDark {
		colors = lightColors
	}

	overlayBase = lipgloss.NewStyle().
		Border(keyboardBorder).
		Padding(1, 3)

	if isDark {
		layoutColor = lipgloss.BrightBlue
		sizeColor = lipgloss.BrightMagenta
		quitColor = lipgloss.BrightRed
		quitBorderColor = lipgloss.BrightBlack
	} else {
		layoutColor = lipgloss.Blue
		sizeColor = lipgloss.Magenta
		quitColor = lipgloss.Red
		quitBorderColor = lipgloss.Black
	}

	fingerStyle = make(map[finger]lipgloss.Style, len(colors))
	fingerActive = make(map[finger]lipgloss.Style, len(colors))

	for finger, c := range colors {
		base := lipgloss.NewStyle().Foreground(c)

		if isDark {
			fingerStyle[finger] = base.Faint(true)
		} else {
			fingerStyle[finger] = base
		}

		fingerActive[finger] = base.Bold(true).Italic(true)
	}

	if isDark {
		statusBarStyle = lipgloss.NewStyle().Foreground(lipgloss.BrightBlack)
	} else {
		statusBarStyle = lipgloss.NewStyle().Faint(true)
	}
}
