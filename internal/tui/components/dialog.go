package components

import (
	"image/color"
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

type DialogModel struct {
	Selected    int
	AccentColor color.Color
	Prompt      string
	LeftLabel   string
	RightLabel  string
}

type DialogAction int

const (
	DialogNone DialogAction = iota
	DialogConfirm
	DialogCancel
)

func (d DialogModel) Update(msg tea.KeyPressMsg) (DialogModel, DialogAction) {
	switch msg.String() {
	case "up", "k", "left":
		d.Selected = 0
	case "down", "j", "right":
		d.Selected = 1
	case "enter":
		if d.Selected == 0 {
			return d, DialogConfirm
		}
		return d, DialogCancel
	case "q", "esc":
		return d, DialogCancel
	}
	return d, DialogNone
}

func (d DialogModel) View() string {
	prompt := d.Prompt
	if prompt == "" {
		prompt = "Are you sure you want to quit?"
	}
	leftLabel := d.LeftLabel
	if leftLabel == "" {
		leftLabel = "Quit"
	}
	rightLabel := d.RightLabel
	if rightLabel == "" {
		rightLabel = "Cancel"
	}

	const contentW = 22
	center := lipgloss.NewStyle().Width(contentW).AlignHorizontal(lipgloss.Center)

	var b strings.Builder
	b.WriteString(center.Render(prompt))
	b.WriteString("\n\n")

	accent := lipgloss.NewStyle().Foreground(d.AccentColor)
	var leftBtn, rightBtn string
	if d.Selected == 0 {
		leftBtn = accent.Render("> " + leftLabel)
		rightBtn = "  " + rightLabel
	} else {
		leftBtn = "  " + leftLabel
		rightBtn = accent.Render("> " + rightLabel)
	}

	b.WriteString(center.Render(leftBtn + "    " + rightBtn))

	return b.String()
}
