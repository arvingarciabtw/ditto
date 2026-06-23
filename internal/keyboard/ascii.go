// Package keyboard renders the ASCII keyboard visualization and manages
// layout remapping (qwerty, dvorak, colemak, etc.) and custom layout files.
package keyboard

import (
	"strings"

	lipgloss "charm.land/lipgloss/v2"
	ansi "github.com/charmbracelet/x/ansi"
)

func Keyboard(size int, layout string, pressedKeys map[uint16]bool, fingerStyle, fingerActive map[Finger]lipgloss.Style) string {
	rows, ok := sizes[size]
	if !ok {
		return ""
	}
	layoutMap := layouts[layout]
	shiftHeld := pressedKeys[42] || pressedKeys[54]
	shiftMap := shiftMaps[layout]

	remapped := make([][]key, len(rows))
	for i, row := range rows {
		remapped[i] = applyLayout(row, layoutMap)
	}

	evCodeToOrigLabel := make(map[uint16]string)
	for _, row := range rows {
		for _, k := range row {
			evCodeToOrigLabel[k.evCode] = k.label
		}
	}

	labelCount := make(map[string]int)
	for _, row := range remapped {
		for _, k := range row {
			labelCount[k.label]++
		}
	}

	pressedByEvCode := make(map[uint16]bool)
	pressedByLabel := make(map[string]bool)
	for code, down := range pressedKeys {
		if !down {
			continue
		}
		label, ok := evCodeToOrigLabel[code]
		if !ok {
			pressedByEvCode[code] = true
			continue
		}
		switch count := labelCount[label]; {
		case count > 1:
			pressedByEvCode[code] = true
		case count == 1:
			pressedByLabel[label] = true
		default:
			pressedByEvCode[code] = true
		}
	}

	var lines []string
	for i, keys := range remapped {
		pressed := make([]bool, len(keys))
		for j, k := range keys {
			if pressedByEvCode[k.evCode] || pressedByLabel[k.label] {
				pressed[j] = true
			}
		}
		if shiftHeld && shiftMap != nil {
			for j := range keys {
				if newLabel, ok := shiftMap[keys[j].label]; ok {
					keys[j].label = newLabel
				}
			}
		}
		if i == 0 {
			lines = append(lines, topLine(keys))
		}
		lines = append(lines, midLine(keys, pressed, fingerStyle, fingerActive))
		if i < len(remapped)-1 {
			lines = append(lines, divLine(keys, fingerStyle))
		} else {
			lines = append(lines, botLine(keys))
		}
	}
	return strings.Join(lines, "\n")
}

func applyLayout(keys []key, layoutMap map[string]string) []key {
	result := make([]key, len(keys))
	for i, k := range keys {
		if layoutMap != nil {
			if newLabel, ok := layoutMap[k.label]; ok {
				k.label = newLabel
			}
		}
		result[i] = k
	}
	return result
}

func topLine(keys []key) string {
	var b strings.Builder
	b.WriteByte(',')
	for _, k := range keys {
		b.WriteString(strings.Repeat("-", k.width))
		b.WriteByte(',')
	}
	return b.String()
}

func midLine(keys []key, pressed []bool, fingerStyle, fingerActive map[Finger]lipgloss.Style) string {
	var b strings.Builder
	for i, k := range keys {
		label := k.label
		if k.divLabel != "" {
			label = ""
		}

		isPressed := i < len(pressed) && pressed[i]

		if i == 0 {
			if isPressed {
				b.WriteString(fingerActive[k.finger].Render("|"))
			} else {
				b.WriteByte('|')
			}
		}

		if isPressed {
			b.WriteString(fingerActive[k.finger].Render(centerLabel(label, k.width)))
		} else {
			b.WriteString(fingerStyle[k.finger].Render(centerLabel(label, k.width)))
		}

		if k.rightless {
			b.WriteByte(' ')
		} else {
			nextPressed := i+1 < len(pressed) && pressed[i+1]
			if isPressed || nextPressed {
				f := k.finger
				if nextPressed && !isPressed {
					f = keys[i+1].finger
				}
				b.WriteString(fingerActive[f].Render("|"))
			} else {
				b.WriteByte('|')
			}
		}
	}
	return b.String()
}

func divLine(keys []key, fingerStyle map[Finger]lipgloss.Style) string {
	var b strings.Builder
	b.WriteByte('|')
	for _, k := range keys {
		if k.gap {
			if k.divLabel != "" {
				b.WriteString(fingerStyle[k.finger].Render(centerLabel(k.divLabel, k.width)))
			} else {
				b.WriteString(strings.Repeat(" ", k.width))
			}
			if k.rightless {
				b.WriteByte(',')
			} else {
				b.WriteByte('\'')
			}
			continue
		}
		b.WriteString(strings.Repeat("-", k.width))
		if k.leftless {
			b.WriteByte(',')
		} else {
			b.WriteByte('\'')
		}
	}
	return b.String()
}

func botLine(keys []key) string {
	var b strings.Builder
	b.WriteByte('\'')
	for _, k := range keys {
		b.WriteString(strings.Repeat("-", k.width))
		if k.leftless {
			b.WriteByte(',')
		} else {
			b.WriteByte('\'')
		}
	}
	return b.String()
}

func centerLabel(s string, width int) string {
	vw := ansi.StringWidth(s)
	if vw >= width {
		return s
	}
	total := width - vw
	left := total / 2
	right := total - left
	return strings.Repeat(" ", left) + s + strings.Repeat(" ", right)
}
