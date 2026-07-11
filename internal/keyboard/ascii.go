/*
Package keyboard provides the public API for keyboard rendering and layout management:
ASCII keyboard display, logical layout maps, custom layout loading, and re-exports of
shared types and constants from the base and standards subpackages.
*/
package keyboard

import (
	"strings"
	"unicode/utf8"

	lipgloss "charm.land/lipgloss/v2"
	ansi "github.com/charmbracelet/x/ansi"

	"github.com/arvingarciabtw/ditto/internal/keyboard/base"
	"github.com/arvingarciabtw/ditto/internal/keyboard/standards"
)

func Render(layout string, size int, standard string, pressedKeys map[uint16]bool, fingerStyle, fingerActive map[Finger]lipgloss.Style) string {
	sd, ok := resolveStandard(standard)
	if !ok {
		return ""
	}
	rows, ok := sd.Sizes[size]
	if !ok {
		return ""
	}

	layoutMap := layouts[layout]
	remapped := make([][]Key, len(rows))
	for i, row := range rows {
		remapped[i] = applyLayout(row, layoutMap)
	}

	shiftHeld := pressedKeys[base.KEY_LEFTSHIFT] || pressedKeys[base.KEY_RIGHTSHIFT]
	altGrHeld := pressedKeys[base.KEY_RIGHTALT]
	kanaHeld := pressedKeys[base.KEY_KATAKANAHIRAGANA]
	hangeulHeld := pressedKeys[base.KEY_HANGEUL]

	shiftMap := sd.ShiftMap
	if lm, ok := shiftMaps[layout]; ok {
		shiftMap = lm
	}
	altGrMap := sd.AltGrMap
	if am, ok := altGrMaps[layout]; ok {
		altGrMap = am
	}

	pressed := resolvePressed(rows, remapped, pressedKeys)
	if kanaHeld {
		applyKana(remapped, shiftHeld)
	}
	if hangeulHeld {
		applyHangeul(remapped, shiftHeld)
	}
	applyModifiers(remapped, shiftHeld, shiftMap, altGrHeld, altGrMap)
	return renderKeys(remapped, pressed, fingerStyle, fingerActive)
}

func applyLayout(keys []Key, layoutMap map[string]string) []Key {
	result := make([]Key, len(keys))
	for i, k := range keys {
		if layoutMap != nil {
			if newLabel, ok := layoutMap[k.Label]; ok {
				k.Label = newLabel
			}
		}
		result[i] = k
	}
	return result
}

func resolveStandard(standard string) (Data, bool) {
	sd, ok := standards.All[standard]
	return sd, ok
}

func ResolveKeycastLabel(scancode uint16, layout, standard string, pressedKeys map[uint16]bool) (string, bool) {
	label, ok := base.EvCodeLabel[scancode]
	if !ok {
		return "", false
	}

	if layoutMap, ok := layouts[layout]; ok && layoutMap != nil {
		if newLabel, ok := layoutMap[label]; ok {
			label = newLabel
		}
	}

	shiftHeld := pressedKeys[base.KEY_LEFTSHIFT] || pressedKeys[base.KEY_RIGHTSHIFT]
	altGrHeld := pressedKeys[base.KEY_RIGHTALT]

	if !shiftHeld && !altGrHeld {
		return label, true
	}

	sd, ok := resolveStandard(standard)
	if !ok {
		return label, true
	}

	shiftMap := sd.ShiftMap
	if lm, ok := shiftMaps[layout]; ok {
		shiftMap = lm
	}
	altGrMap := sd.AltGrMap
	if am, ok := altGrMaps[layout]; ok {
		altGrMap = am
	}

	if altGrHeld && altGrMap != nil {
		if newLabel, ok := altGrMap[label]; ok {
			label = newLabel
		}
		if shiftHeld && utf8.RuneCountInString(label) == 1 {
			label = strings.ToUpper(label)
		}
	} else if shiftHeld && shiftMap != nil {
		if newLabel, ok := shiftMap[label]; ok {
			label = newLabel
		}
	}

	return label, true
}

func resolvePressed(rows, remapped [][]Key, pressedKeys map[uint16]bool) [][]bool {
	evCodeToOrigLabel := make(map[uint16]string)
	for _, row := range rows {
		for _, k := range row {
			evCodeToOrigLabel[k.EvCode] = k.Label
		}
	}

	labelCount := make(map[string]int)
	for _, row := range remapped {
		for _, k := range row {
			labelCount[k.Label]++
		}
	}

	evCodeCount := make(map[uint16]int)
	for _, row := range rows {
		for _, k := range row {
			evCodeCount[k.EvCode]++
		}
	}

	pressedByEvCode := make(map[uint16]bool)
	pressedByLabel := make(map[string]bool)
	for code, down := range pressedKeys {
		if !down {
			continue
		}
		if evCodeCount[code] > 1 {
			pressedByEvCode[code] = true
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

	pressed := make([][]bool, len(remapped))
	for i, keys := range remapped {
		pressed[i] = make([]bool, len(keys))
		for j, k := range keys {
			if pressedByEvCode[k.EvCode] || pressedByLabel[k.Label] {
				pressed[i][j] = true
			}
		}
	}
	return pressed
}

func applyModifiers(keys [][]Key, shiftHeld bool, shiftMap map[string]string, altGrHeld bool, altGrMap map[string]string) {
	for _, row := range keys {
		if altGrHeld && altGrMap != nil {
			for j := range row {
				if newLabel, ok := altGrMap[row[j].Label]; ok {
					row[j].Label = newLabel
				}
			}
			if shiftHeld {
				for j := range row {
					if len(row[j].Label) == 1 {
						row[j].Label = strings.ToUpper(row[j].Label)
					}
				}
			}
		} else if shiftHeld && shiftMap != nil {
			for j := range row {
				if newLabel, ok := shiftMap[row[j].Label]; ok {
					row[j].Label = newLabel
				}
			}
		}
	}
}

func applyKana(keys [][]Key, shiftHeld bool) {
	for _, row := range keys {
		for j := range row {
			label := row[j].Label
			if shiftHeld {
				if v, ok := standards.KanaSmallMap[label]; ok {
					row[j].Label = v
					continue
				}
			}
			if v, ok := standards.KanaLayout[label]; ok {
				row[j].Label = v
			}
		}
	}
}

func applyHangeul(keys [][]Key, shiftHeld bool) {
	for _, row := range keys {
		for j := range row {
			label := row[j].Label
			if shiftHeld {
				if tensed, ok := standards.HangeulTensed[label]; ok {
					row[j].Label = tensed
					continue
				}
			}
			if v, ok := standards.HangeulLayout[label]; ok {
				row[j].Label = v
			}
		}
	}
}

func renderKeys(keys [][]Key, pressed [][]bool, fingerStyle, fingerActive map[Finger]lipgloss.Style) string {
	var lines []string
	for i, kr := range keys {
		if i == 0 {
			lines = append(lines, topLine(kr))
		}
		lines = append(lines, midLine(kr, pressed[i], fingerStyle, fingerActive))
		if i < len(keys)-1 {
			lines = append(lines, divLine(kr, fingerStyle))
		} else {
			lines = append(lines, botLine(kr))
		}
	}
	return strings.Join(lines, "\n")
}

func topLine(keys []Key) string {
	var b strings.Builder
	b.WriteByte(',')
	for _, k := range keys {
		b.WriteString(strings.Repeat("-", k.Width))
		b.WriteByte(',')
	}
	return b.String()
}

func midLine(keys []Key, pressed []bool, fingerStyle, fingerActive map[Finger]lipgloss.Style) string {
	var b strings.Builder
	for i, k := range keys {
		label := k.Label
		if k.DivLabel != "" {
			label = ""
		}

		isPressed := i < len(pressed) && pressed[i]

		if i == 0 {
			if isPressed {
				b.WriteString(fingerActive[k.Finger].Render("|"))
			} else {
				b.WriteByte('|')
			}
		}

		if isPressed {
			b.WriteString(fingerActive[k.Finger].Render(centerLabel(label, k.Width)))
		} else {
			b.WriteString(fingerStyle[k.Finger].Render(centerLabel(label, k.Width)))
		}

		if k.Rightless {
			b.WriteByte(' ')
		} else {
			nextPressed := i+1 < len(pressed) && pressed[i+1]
			if isPressed || nextPressed {
				f := k.Finger
				if nextPressed && !isPressed {
					f = keys[i+1].Finger
				}
				b.WriteString(fingerActive[f].Render("|"))
			} else {
				b.WriteByte('|')
			}
		}
	}
	return b.String()
}

func divLine(keys []Key, fingerStyle map[Finger]lipgloss.Style) string {
	var b strings.Builder
	b.WriteByte('|')
	for _, k := range keys {
		if k.Gap {
			if k.DivLabel != "" {
				b.WriteString(fingerStyle[k.Finger].Render(centerLabel(k.DivLabel, k.Width)))
			} else {
				b.WriteString(strings.Repeat(" ", k.Width))
			}
			if k.Rightless {
				b.WriteByte(',')
			} else {
				b.WriteByte('\'')
			}
			continue
		}
		b.WriteString(strings.Repeat("-", k.Width))
		if k.Leftless {
			b.WriteByte(',')
		} else {
			b.WriteByte('\'')
		}
	}
	return b.String()
}

func botLine(keys []Key) string {
	var b strings.Builder
	b.WriteByte('\'')
	for _, k := range keys {
		b.WriteString(strings.Repeat("-", k.Width))
		if k.Leftless {
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
