package main

import lipgloss "charm.land/lipgloss/v2"

var (
	listFrameStyle = lipgloss.NewStyle().Margin(1, 2)
	listStyle      = lipgloss.NewStyle().
			BorderForeground(lipgloss.BrightBlack).
			Border(lipgloss.ThickBorder()).
			Padding(1, 2)
)
