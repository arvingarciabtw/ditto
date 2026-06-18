package main

import (
	"fmt"
	"os"
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	ansi "github.com/charmbracelet/x/ansi"
	term "github.com/charmbracelet/x/term"
)

func (m Model) View() tea.View {
	s := buildBaseView(m)

	hasOpenList := m.showLayoutList || m.showSizeList
	if hasOpenList {
		termW, termH, _ := buildTerminalSize()

		pw := min(60, max(30, termW-4))
		ph := min(20, max(10, termH-4))
		x := (termW - pw) / 2
		y := (termH - ph) / 2

		var menu string
		listStyle.Width(pw).Height(ph)

		if m.showLayoutList {
			m.layoutList.SetSize(pw-6, ph-4)
			menu = listStyle.Render(m.layoutList.View())
		}
		if m.showSizeList {
			m.sizeList.SetSize(pw-6, ph-4)
			menu = listStyle.Render(m.sizeList.View())
		}

		s = buildOverlay(s, menu, x, y)
	}

	v := tea.NewView(s)
	v.AltScreen = true
	return v
}

func buildBaseView(m Model) string {
	termW, termH, err := buildTerminalSize()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v", err)
		os.Exit(1)
	}

	keyboard := buildKeyboard(m.activeSize, m.activeLayout, m.pressedKeys)

	if !m.showInfoBar {
		return lipgloss.Place(termW, termH, lipgloss.Center, lipgloss.Center, keyboard)
	}

	infoBar := buildInfoBar(m, termW)
	return lipgloss.Place(termW, termH-1, lipgloss.Center, lipgloss.Center, keyboard) + "\n" + infoBar
}

func buildInfoBar(m Model, terminalWidth int) string {
	activeSize := m.helpModel.Styles.FullKey.Render(fmt.Sprintf("%d%%", m.activeSize))
	activeLayout := m.helpModel.Styles.FullDesc.Render(" •︎", m.activeLayout)
	actives := lipgloss.JoinHorizontal(lipgloss.Bottom, activeSize, "", activeLayout)

	commands := m.helpModel.View(commands)

	spacerWidth := terminalWidth - lipgloss.Width(actives) - lipgloss.Width(commands)
	spacer := strings.Repeat(" ", max(0, spacerWidth))

	return lipgloss.JoinHorizontal(lipgloss.Bottom, actives, spacer, commands)
}

func buildOverlay(bg string, menu string, x, y int) string {
	bgLines := strings.Split(bg, "\n")
	popupLines := strings.Split(menu, "\n")
	for py, pl := range popupLines {
		by := y + py
		if by < 0 || by >= len(bgLines) {
			continue
		}
		bl := bgLines[by]
		bgW := ansi.StringWidth(bl)
		pw := ansi.StringWidth(pl)

		prefix := ansi.Cut(bl, 0, x)
		var suffix string
		if x+pw < bgW {
			suffix = ansi.Cut(bl, x+pw, bgW)
		}
		bgLines[by] = prefix + pl + suffix
	}
	return strings.Join(bgLines, "\n")
}

func buildTerminalSize() (int, int, error) {
	width, height, err := term.GetSize(uintptr(os.Stdout.Fd()))
	return width, height, err
}
