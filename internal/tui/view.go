package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	ansi "github.com/charmbracelet/x/ansi"

	"github.com/arvingarciabtw/ditto/internal/keyboard"
	"github.com/arvingarciabtw/ditto/internal/tui/components"
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
			ov = OverlayBase.BorderForeground(LayoutColor).Width(overlayWidth).Height(h).Render(m.layoutList.View(StatusBarStyle))
		case m.showSizeList:
			h = 10
			ov = OverlayBase.BorderForeground(SizeColor).Width(overlayWidth).Height(h).Render(m.sizeList.View(StatusBarStyle))
		case m.showQuitDialog:
			h = 8
			ov = OverlayBase.BorderForeground(QuitBorderColor).Width(overlayWidth).Height(h).Render(m.quitDialog.View())
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

	kb := keyboard.Keyboard(m.activeSize, m.activeLayout, m.pressedKeys, FingerStyle, FingerActive)
	kh := strings.Count(kb, "\n") + 1
	kw := 0
	for line := range strings.SplitSeq(kb, "\n") {
		if w := lipgloss.Width(line); w > kw {
			kw = w
		}
	}

	neededH, neededW := kh, kw
	if m.showAllInfo {
		neededH += 4
	}

	if th < neededH || tw < neededW {
		return warning(tw, th, neededW, neededH)
	}

	if !m.showAllInfo {
		return lipgloss.Place(tw, th, lipgloss.Center, lipgloss.Center, kb)
	}

	bar := statusBar(m, kw)
	leg := legends(kw)
	content := leg + "\n" + kb + "\n" + bar

	return lipgloss.Place(tw, th, lipgloss.Center, lipgloss.Center, content)
}

func warning(tw, th, nw, nh int) string {
	header := lipgloss.NewStyle().Foreground(QuitColor).Render("Terminal size too small:")
	size := WarningStyle.Render(fmt.Sprintf("Width = %d  Height = %d", tw, th))
	need := WarningAccent.Render("Needed for current config:")
	needSize := WarningStyle.Render(fmt.Sprintf("Width = %d  Height = %d", nw, nh))

	lines := []string{header, size, "", need, needSize}
	return lipgloss.Place(tw, th, lipgloss.Center, lipgloss.Center, strings.Join(lines, "\n"))
}

func legends(width int) string {
	type legend struct {
		name  string
		style lipgloss.Style
	}

	items := []legend{
		{name: "pinky", style: FingerStyle[keyboard.Pinky]},
		{name: "ring", style: FingerStyle[keyboard.Ring]},
		{name: "middle", style: FingerStyle[keyboard.Middle]},
		{name: "index", style: FingerStyle[keyboard.Index]},
		{name: "thumb", style: FingerStyle[keyboard.Thumb]},
		{name: "any", style: FingerStyle[keyboard.Any]},
	}

	symbol := "•︎"

	sb := strings.Builder{}
	for _, legend := range items {
		fmt.Fprintf(&sb, "%s %s ", legend.style.Render(symbol), StatusBarStyle.Render(legend.name))
	}
	s := sb.String()

	sw := width - lipgloss.Width(s)
	spacer := strings.Repeat(" ", max(0, sw))

	return lipgloss.JoinHorizontal(lipgloss.Bottom, s, spacer)
}

func statusBar(m Model, width int) string {
	size := StatusBarStyle.Render(fmt.Sprintf("%d%%", m.activeSize))
	layout := StatusBarStyle.Render(" •︎", m.activeLayout)

	actives := lipgloss.JoinHorizontal(lipgloss.Bottom, size, "", layout)
	bindings := renderBindings(components.Commands)

	sw := width - lipgloss.Width(actives) - lipgloss.Width(bindings)
	spacer := strings.Repeat(" ", max(0, sw))

	return lipgloss.JoinHorizontal(lipgloss.Top, actives, spacer, bindings)
}

func renderBindings(c components.Bindings) string {
	parts := []string{
		StatusBarStyle.Render(c.Layout.Help().Key) + " " + StatusBarStyle.Render(c.Layout.Help().Desc),
		StatusBarStyle.Render(c.Size.Help().Key) + " " + StatusBarStyle.Render(c.Size.Help().Desc),
		StatusBarStyle.Render(c.HideKey.Help().Key) + " " + StatusBarStyle.Render(c.HideKey.Help().Desc),
	}
	return strings.Join(parts, "  ")
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
