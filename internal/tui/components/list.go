// Package components provides reusable TUI widgets: a keyboard-layout/size
// list picker, a dialog, and key binding definitions.
package components

import (
	"image/color"
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

type ListModel struct {
	Items       []string
	Selected    int
	Title       string
	AccentColor color.Color
}

type ListAction int

const (
	ListNone ListAction = iota
	ListConfirm
	ListCancel
)

func (l ListModel) Update(msg tea.KeyPressMsg) (ListModel, ListAction) {
	switch msg.String() {
	case "up", "k", "left":
		if l.Selected > 0 {
			l.Selected--
		}
	case "down", "j", "right":
		if l.Selected < len(l.Items)-1 {
			l.Selected++
		}
	case "enter":
		return l, ListConfirm
	case "q", "esc":
		return l, ListCancel
	}
	return l, ListNone
}

func (l ListModel) View(statusBarStyle lipgloss.Style) string {
	var b strings.Builder

	accent := lipgloss.NewStyle().Foreground(l.AccentColor)
	titleLine := accent.Render(l.Title)
	b.WriteString(titleLine)
	b.WriteString("\n\n")

	maxWidth := lipgloss.Width(titleLine)

	start := l.windowStart()
	visible := l.visibleItems()

	var itemLines []string
	for i, item := range visible {
		var line string
		if start+i == l.Selected {
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

func (l ListModel) windowStart() int {
	if len(l.Items) <= 3 {
		return 0
	}
	return max(0, min(l.Selected-1, len(l.Items)-3))
}

func (l ListModel) visibleItems() []string {
	start := l.windowStart()
	return l.Items[start : start+3]
}
