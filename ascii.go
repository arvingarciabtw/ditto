package main

import (
	"strings"
)

func applyLayout(keys []Key, layoutMap map[string]string) []Key {
	if layoutMap == nil {
		return keys
	}
	result := make([]Key, len(keys))
	for i, k := range keys {
		if newLabel, ok := layoutMap[k.Label]; ok {
			k.Label = newLabel
		}
		result[i] = k
	}
	return result
}

func buildKeyboard(size int, layout string, pressedKeys map[uint16]bool) string {
	rows, ok := keyboardSizes[size]
	if !ok {
		return ""
	}
	layoutMap := keyboardLayouts[layout]

	evCodeToLabel := make(map[uint16]string)
	labelCount := make(map[string]int)
	for _, row := range rows {
		for _, k := range row {
			evCodeToLabel[k.EvCode] = k.Label
			labelCount[k.Label]++
		}
	}

	pressedByLabel := make(map[string]bool)
	pressedByEvCode := make(map[uint16]bool)
	for code, down := range pressedKeys {
		if !down {
			continue
		}
		label, ok := evCodeToLabel[code]
		if !ok {
			continue
		}
		if labelCount[label] > 1 {
			pressedByEvCode[code] = true
		} else {
			pressedByLabel[label] = true
		}
	}

	var lines []string
	for i, row := range rows {
		keys := applyLayout(row, layoutMap)
		pressed := make([]bool, len(keys))
		for j, k := range keys {
			if pressedByEvCode[k.EvCode] || pressedByLabel[k.Label] {
				pressed[j] = true
			}
		}
		if i == 0 {
			lines = append(lines, buildTopLine(keys))
		}
		lines = append(lines, buildMidLine(keys, pressed))
		if i < len(rows)-1 {
			lines = append(lines, buildDivLine(keys))
		} else {
			lines = append(lines, buildBotLine(keys))
		}
	}
	return strings.Join(lines, "\n")
}

func buildTopLine(keys []Key) string {
	var b strings.Builder
	b.WriteByte(',')
	for _, k := range keys {
		b.WriteString(strings.Repeat("-", k.Width))
		b.WriteByte(',')
	}
	return b.String()
}

func buildMidLine(keys []Key, pressed []bool) string {
	var b strings.Builder
	b.WriteByte('|')
	for i, k := range keys {
		label := k.Label
		if k.DivLabel != "" {
			label = ""
		}
		if i < len(pressed) && pressed[i] {
			b.WriteString(fingerActive[k.Finger].Render(centerLabel(label, k.Width)))
		} else {
			b.WriteString(fingerStyle[k.Finger].Render(centerLabel(label, k.Width)))
		}
		if k.Rightless {
			b.WriteByte(' ')
			continue
		}
		b.WriteByte('|')
	}
	return b.String()
}

func buildDivLine(keys []Key) string {
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

func buildBotLine(keys []Key) string {
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
