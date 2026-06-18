package main

import lipgloss "charm.land/lipgloss/v2"

var (
	listFrameStyle = lipgloss.NewStyle().Margin(1, 2)

	listStyle = lipgloss.NewStyle().
			BorderForeground(lipgloss.BrightBlack).
			Border(lipgloss.ThickBorder()).
			Padding(1, 2)

	fingerStyle = map[Finger]lipgloss.Style{
		FingerPinky:  lipgloss.NewStyle().Faint(true).Foreground(lipgloss.BrightMagenta),
		FingerRing:   lipgloss.NewStyle().Faint(true).Foreground(lipgloss.BrightRed),
		FingerMiddle: lipgloss.NewStyle().Faint(true).Foreground(lipgloss.BrightYellow),
		FingerIndex:  lipgloss.NewStyle().Faint(true).Foreground(lipgloss.BrightCyan),
		FingerThumb:  lipgloss.NewStyle().Faint(true).Foreground(lipgloss.BrightGreen),
	}

	fingerActive = map[Finger]lipgloss.Style{
		FingerPinky:  lipgloss.NewStyle().Bold(true).Italic(true).Foreground(lipgloss.BrightMagenta),
		FingerRing:   lipgloss.NewStyle().Bold(true).Italic(true).Foreground(lipgloss.BrightRed),
		FingerMiddle: lipgloss.NewStyle().Bold(true).Italic(true).Foreground(lipgloss.BrightYellow),
		FingerIndex:  lipgloss.NewStyle().Bold(true).Italic(true).Foreground(lipgloss.BrightCyan),
		FingerThumb:  lipgloss.NewStyle().Bold(true).Italic(true).Foreground(lipgloss.BrightGreen),
	}
)
