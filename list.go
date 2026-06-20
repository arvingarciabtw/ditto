package main

import (
	"image/color"
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

type listModel struct {
	items       []string
	selected    int
	title       string
	accentColor color.Color
}

type listAction int

const (
	listNone listAction = iota
	listConfirm
	listCancel
)

var layoutListItems = []string{
	"qwerty",
	"dvorak",
	"colemak",
	"colemak-dh",
	"workman",
	"azerty",
}

var layoutSizeItems = []string{
	"60%",
	"65%",
	"75%",
	"80%",
	"96%",
	"100%",
}

func (l listModel) Update(msg tea.KeyPressMsg) (listModel, listAction) {
	switch msg.String() {
	case "up", "k", "left":
		if l.selected > 0 {
			l.selected--
		}
	case "down", "j", "right":
		if l.selected < len(l.items)-1 {
			l.selected++
		}
	case "enter":
		return l, listConfirm
	case "q", "esc":
		return l, listCancel
	}
	return l, listNone
}

func (l listModel) View() string {
	var b strings.Builder

	accent := lipgloss.NewStyle().Foreground(l.accentColor)
	titleLine := accent.Render(l.title)
	b.WriteString(titleLine)
	b.WriteString("\n\n")

	maxWidth := lipgloss.Width(titleLine)

	start := l.windowStart()
	visible := l.visibleItems()

	var itemLines []string
	for i, item := range visible {
		var line string
		if start+i == l.selected {
			line = accent.Render("> " + item)
		} else {
			line = "  " + item
		}
		itemLines = append(itemLines, line)
		if w := lipgloss.Width(line); w > maxWidth {
			maxWidth = w
		}
	}

	for i, line := range itemLines {
		b.WriteString(line)
		if i < len(itemLines)-1 {
			b.WriteString("\n")
		}
	}

	b.WriteString("\n\n")

	help := statusBarStyle.Render("↵ / enter • q / quit")
	helpWidth := lipgloss.Width(help)
	padding := maxWidth - helpWidth
	if padding > 0 {
		b.WriteString(strings.Repeat(" ", padding))
	}
	b.WriteString(help)

	return b.String()
}

func (l listModel) windowStart() int {
	if len(l.items) <= 3 {
		return 0
	}
	return max(0, min(l.selected-1, len(l.items)-3))
}

func (l listModel) visibleItems() []string {
	start := l.windowStart()
	return l.items[start : start+3]
}
