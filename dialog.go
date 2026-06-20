package main

import (
	"image/color"
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

type dialogModel struct {
	selected    int
	accentColor color.Color
}

type dialogAction int

const (
	dialogNone dialogAction = iota
	dialogConfirm
	dialogCancel
)

func (d dialogModel) Update(msg tea.KeyPressMsg) (dialogModel, dialogAction) {
	switch msg.String() {
	case "up", "k", "left":
		d.selected = 0
	case "down", "j", "right":
		d.selected = 1
	case "enter":
		if d.selected == 0 {
			return d, dialogConfirm
		}
		return d, dialogCancel
	case "q", "esc":
		return d, dialogCancel
	}
	return d, dialogNone
}

func (d dialogModel) View() string {
	const contentW = 22
	center := lipgloss.NewStyle().Width(contentW).AlignHorizontal(lipgloss.Center)

	var b strings.Builder
	b.WriteString(center.Render("Are you sure you want to quit?"))
	b.WriteString("\n\n")

	accent := lipgloss.NewStyle().Foreground(d.accentColor)
	var leftBtn, rightBtn string
	if d.selected == 0 {
		leftBtn = accent.Render("> Quit")
		rightBtn = "  Cancel"
	} else {
		leftBtn = "  Quit"
		rightBtn = accent.Render("> Cancel")
	}

	b.WriteString(center.Render(leftBtn + "    " + rightBtn))

	return b.String()
}
