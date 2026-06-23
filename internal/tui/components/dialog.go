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
	const contentW = 22
	center := lipgloss.NewStyle().Width(contentW).AlignHorizontal(lipgloss.Center)

	var b strings.Builder
	b.WriteString(center.Render("Are you sure you want to quit?"))
	b.WriteString("\n\n")

	accent := lipgloss.NewStyle().Foreground(d.AccentColor)
	var leftBtn, rightBtn string
	if d.Selected == 0 {
		leftBtn = accent.Render("> Quit")
		rightBtn = "  Cancel"
	} else {
		leftBtn = "  Quit"
		rightBtn = accent.Render("> Cancel")
	}

	b.WriteString(center.Render(leftBtn + "    " + rightBtn))

	return b.String()
}
