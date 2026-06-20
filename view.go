package main

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	ansi "github.com/charmbracelet/x/ansi"
)

func (m Model) View() tea.View {
	const overlayWidth = 30

	s := base(m)

	hasOverlay := m.showLayoutList || m.showSizeList || m.showQuitDialog

	if hasOverlay {
		tw, th := m.terminalWidth, m.terminalHeight
		var ov string
		var h int

		switch {
		case m.showLayoutList:
			h = 10
			ov = overlayBase.BorderForeground(layoutColor).Width(overlayWidth).Height(h).Render(m.layoutList.View())
		case m.showSizeList:
			h = 10
			ov = overlayBase.BorderForeground(sizeColor).Width(overlayWidth).Height(h).Render(m.sizeList.View())
		case m.showQuitDialog:
			h = 8
			ov = overlayBase.BorderForeground(quitBorderColor).Width(overlayWidth).Height(h).Render(m.quitDialog.View())
		}

		x := (tw - overlayWidth) / 2
		y := (th - h) / 2

		s = overlay(s, ov, x, y)
	}

	v := tea.NewView(s)
	v.AltScreen = true
	return v
}

func base(m Model) string {
	tw, th := m.terminalWidth, m.terminalHeight
	if tw == 0 || th == 0 {
		return ""
	}

	kb := keyboard(m.activeSize, m.activeLayout, m.pressedKeys)
	kw := 0

	if !m.showAllInfo {
		return lipgloss.Place(tw, th, lipgloss.Center, lipgloss.Center, kb)
	}

	for line := range strings.SplitSeq(kb, "\n") {
		if w := lipgloss.Width(line); w > kw {
			kw = w
		}
	}

	bar := statusBar(m, kw)
	leg := legends(kw)
	content := leg + "\n" + kb + "\n" + bar

	return lipgloss.Place(tw, th, lipgloss.Center, lipgloss.Center, content)
}

func legends(width int) string {
	type legend struct {
		name  string
		style lipgloss.Style
	}

	items := []legend{
		{name: "pinky", style: fingerStyle[pinky]},
		{name: "ring", style: fingerStyle[ring]},
		{name: "middle", style: fingerStyle[middle]},
		{name: "index", style: fingerStyle[index]},
		{name: "thumb", style: fingerStyle[thumb]},
		{name: "any", style: fingerStyle[any]},
	}

	symbol := "•︎"

	sb := strings.Builder{}
	for _, legend := range items {
		fmt.Fprintf(&sb, "%s %s ", legend.style.Render(symbol), statusBarStyle.Render(legend.name))
	}
	s := sb.String()

	sw := width - lipgloss.Width(s)
	spacer := strings.Repeat(" ", max(0, sw))

	return lipgloss.JoinHorizontal(lipgloss.Bottom, s, spacer)
}

func statusBar(m Model, width int) string {
	size := m.help.Styles.FullKey.Render(fmt.Sprintf("%d%%", m.activeSize))
	layout := m.help.Styles.FullDesc.Render(" •︎", m.activeLayout)

	actives := lipgloss.JoinHorizontal(lipgloss.Bottom, size, "", layout)
	bindings := m.help.View(commands)

	sw := width - lipgloss.Width(actives) - lipgloss.Width(bindings)
	spacer := strings.Repeat(" ", max(0, sw))

	return lipgloss.JoinHorizontal(lipgloss.Top, actives, spacer, bindings)
}

func overlay(bg string, ov string, x, y int) string {
	bgLines := strings.Split(bg, "\n")
	overlayLines := strings.Split(ov, "\n")

	for oy, ol := range overlayLines {
		by := y + oy
		if by < 0 || by >= len(bgLines) {
			continue
		}
		bl := bgLines[by]
		bgW := ansi.StringWidth(bl)
		w := ansi.StringWidth(ol)

		if x < 0 || x >= bgW {
			continue
		}

		prefix := ansi.Cut(bl, 0, x)
		var suffix string
		if x+w < bgW {
			suffix = ansi.Cut(bl, x+w, bgW)
		}
		bgLines[by] = prefix + ol + suffix
	}

	return strings.Join(bgLines, "\n")
}
