package tui

import (
	"fmt"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
	ansi "github.com/charmbracelet/x/ansi"

	"github.com/arvingarciabtw/ditto/internal/keyboard"
	"github.com/arvingarciabtw/ditto/internal/tui/components"
)

func (m Model) View() tea.View {
	const overlayWidth = 30

	var s string
	if m.keycastMode {
		s = keycastView(m)
	} else {
		s = base(m)
	}

	hasOverlay := m.showLayoutList || m.showSizeList || m.showStandardList || m.showQuitDialog || m.showModeList

	if hasOverlay {
		tw, th := m.terminalWidth, m.terminalHeight
		var ov string
		var h int

		border := KeyboardBorder
		if m.keycastBoxDraw {
			border = KeyboardBorderBoxDraw
		}
		overlayBase := lipgloss.NewStyle().Border(border).Padding(1, 3)

		switch {
		case m.showLayoutList:
			h = min(th, max(11, min(th-4, 13)))
			m.layoutList.VisibleCount = h - 8
			ov = overlayBase.BorderForeground(LayoutColor).Width(overlayWidth).Height(h).Render(m.layoutList.View(StatusBarStyle))
		case m.showSizeList:
			h = min(th, max(11, min(th-4, 13)))
			m.sizeList.VisibleCount = h - 8
			ov = overlayBase.BorderForeground(SizeColor).Width(overlayWidth).Height(h).Render(m.sizeList.View(StatusBarStyle))
		case m.showStandardList:
			h = min(th, max(11, min(th-4, 13)))
			m.standardList.VisibleCount = h - 8
			ov = overlayBase.BorderForeground(StandardColor).Width(overlayWidth).Height(h).Render(m.standardList.View(StatusBarStyle))
		case m.showQuitDialog:
			h = 8
			ov = overlayBase.BorderForeground(QuitBorderColor).Width(overlayWidth).Height(h).Render(m.quitDialog.View())
		case m.showModeList:
			h = 8
			m.modeList.VisibleCount = 2
			ov = overlayBase.BorderForeground(ModeColor).Width(overlayWidth).Height(h).Render(m.modeList.View(StatusBarStyle))
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

	kb := keyboard.Render(m.activeLayout, m.activeSize, m.activeStandard, m.pressedKeys, FingerStyle, FingerActive)
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

	top := topBar(m, kw)
	bar := statusBar(m, kw)
	content := top + "\n" + kb + "\n" + bar

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

func topBar(m Model, width int) string {
	std := StatusBarStyle.Render(m.activeStandard)
	if m.locked {
		std = std + " " + StatusBarStyle.Render("•") + " " + StatusBarStyle.Render("locked")
	}

	size := StatusBarStyle.Render(fmt.Sprintf("%d%%", m.activeSize))
	layout := StatusBarStyle.Render(" • " + m.activeLayout + " • ")
	actives := lipgloss.JoinHorizontal(lipgloss.Bottom, size, layout, std)

	sw := width - lipgloss.Width(actives)
	spacer := strings.Repeat(" ", max(0, sw))

	return lipgloss.JoinHorizontal(lipgloss.Bottom, spacer, actives)
}

func statusBar(m Model, width int) string {
	bindings := renderBindings(components.Commands, m.activeStandard)
	sw := width - lipgloss.Width(bindings)
	spacer := strings.Repeat(" ", max(0, sw))
	return lipgloss.JoinHorizontal(lipgloss.Top, spacer, bindings)
}

func renderBindings(c components.Bindings, activeStandard string) string {
	standardDesc := c.Standard.Help().Desc

	parts := []string{
		StatusBarStyle.Render(c.Layout.Help().Key) + " " + StatusBarStyle.Render(c.Layout.Help().Desc),
		StatusBarStyle.Render(c.Size.Help().Key) + " " + StatusBarStyle.Render(c.Size.Help().Desc),
		StatusBarStyle.Render(c.Standard.Help().Key) + " " + StatusBarStyle.Render(standardDesc),
		StatusBarStyle.Render(c.Keycast.Help().Key) + " " + StatusBarStyle.Render(c.Keycast.Help().Desc),
	}

	switch activeStandard {
	case "jis":
		parts = append(parts, StatusBarStyle.Render(c.Kana.Help().Key)+" "+StatusBarStyle.Render(c.Kana.Help().Desc))
	case "ks":
		parts = append(parts, StatusBarStyle.Render(c.Hangeul.Help().Key)+" "+StatusBarStyle.Render(c.Hangeul.Help().Desc))
	}

	parts = append(parts, StatusBarStyle.Render(c.HideKey.Help().Key)+" "+StatusBarStyle.Render(c.HideKey.Help().Desc))
	return strings.Join(parts, StatusBarStyle.Render(" • "))
}

func keycastView(m Model) string {
	tw, th := m.terminalWidth, m.terminalHeight
	if tw == 0 || th == 0 {
		return ""
	}

	now := time.Now()
	var entries []keycastEntry
	for _, e := range m.keycastKeys {
		if now.Sub(e.pressedAt) < 1500*time.Millisecond {
			entries = append(entries, e)
		}
	}

	labels := make([]string, len(entries))
	for i, e := range entries {
		labels[i] = e.label
	}
	labels = fitLabelsToWidth(labels, tw)

	var kept []keycastEntry
	for _, l := range labels {
		for _, e := range entries {
			if e.label == l {
				kept = append(kept, e)
				break
			}
		}
	}

	row := keycastBoxRow(kept, m.keycastFingerColors, m.keycastBoxDraw)

	showBar := m.showAllInfo
	if !showBar {
		if row == "" {
			return lipgloss.Place(tw, th, lipgloss.Center, lipgloss.Center, "")
		}
		return lipgloss.Place(tw, th, lipgloss.Center, lipgloss.Center, row)
	}

	help := "h hide • m mode • q quit"
	if m.keycastMode {
		help = "b box draw • f finger color • m mode • h hide •  q quit"
	}
	cmd := StatusBarStyle.Render(help)
	cmdLine := lipgloss.Place(tw, 1, lipgloss.Center, lipgloss.Center, cmd)

	availH := th - 2
	if row == "" {
		return lipgloss.Place(tw, availH, lipgloss.Center, lipgloss.Center, "") + "\n" + cmdLine
	}

	keyArea := lipgloss.Place(tw, availH, lipgloss.Center, lipgloss.Center, row)
	return keyArea + "\n" + cmdLine
}

func fitLabelsToWidth(labels []string, maxWidth int) []string {
	if len(labels) == 0 {
		return nil
	}
	for len(labels) > 1 {
		tops := make([]string, len(labels))
		for i, l := range labels {
			w := lipgloss.Width(l)
			tops[i] = "," + strings.Repeat("-", w+2) + ","
		}
		if lipgloss.Width(strings.Join(tops, " ")) <= maxWidth {
			break
		}
		labels = labels[1:]
	}
	return labels
}

func keycastBoxRow(entries []keycastEntry, useColors, boxDraw bool) string {
	if len(entries) == 0 {
		return ""
	}
	tops := make([]string, len(entries))
	mids := make([]string, len(entries))
	bots := make([]string, len(entries))
	for i, e := range entries {
		l := e.label
		w := lipgloss.Width(l)
		if boxDraw {
			if useColors {
				s := FingerStyle[e.finger]
				tops[i] = s.Render("╭" + strings.Repeat("─", w+2) + "╮")
				mids[i] = s.Render("│ " + l + " │")
				bots[i] = s.Render("╰" + strings.Repeat("─", w+2) + "╯")
			} else {
				tops[i] = "╭" + strings.Repeat("─", w+2) + "╮"
				mids[i] = "│ " + l + " │"
				bots[i] = "╰" + strings.Repeat("─", w+2) + "╯"
			}
		} else {
			if useColors {
				s := FingerStyle[e.finger]
				tops[i] = s.Render("," + strings.Repeat("-", w+2) + ",")
				mids[i] = s.Render("| " + l + " |")
				bots[i] = s.Render("'" + strings.Repeat("-", w+2) + "'")
			} else {
				tops[i] = "," + strings.Repeat("-", w+2) + ","
				mids[i] = "| " + l + " |"
				bots[i] = "'" + strings.Repeat("-", w+2) + "'"
			}
		}
	}
	return strings.Join(tops, " ") + "\n" + strings.Join(mids, " ") + "\n" + strings.Join(bots, " ")
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

		prefix := ansi.Truncate(bl, x, "")
		pw := ansi.StringWidth(prefix)
		if pw < x {
			prefix += strings.Repeat(" ", x-pw)
		}

		var suffix string
		startCol := x + w
		if startCol < bgW {
			suffix = ansi.Cut(bl, startCol, bgW)
			sw := ansi.StringWidth(suffix)
			expected := bgW - startCol
			if sw > expected {
				suffix = ansi.Truncate(suffix, expected, "")
			} else if sw < expected {
				suffix += strings.Repeat(" ", expected-sw)
			}
		}
		bgLines[by] = prefix + ol + suffix
	}

	return strings.Join(bgLines, "\n")
}
