package main

import lipgloss "charm.land/lipgloss/v2"

var (
	listFrameStyle = lipgloss.NewStyle().Margin(1, 2)
	listStyle      = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(1, 2)
)
