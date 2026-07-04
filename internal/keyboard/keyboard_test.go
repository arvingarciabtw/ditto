package keyboard

import (
	"fmt"
	"strings"
	"testing"

	lipgloss "charm.land/lipgloss/v2"

	"github.com/arvingarciabtw/ditto/internal/keyboard/base"
	"github.com/arvingarciabtw/ditto/internal/keyboard/standards"
)

func emptyStyles() (map[Finger]lipgloss.Style, map[Finger]lipgloss.Style) {
	fs := make(map[Finger]lipgloss.Style)
	fa := make(map[Finger]lipgloss.Style)
	for _, f := range []Finger{Pinky, Ring, Middle, Index, Thumb, Any} {
		fs[f] = lipgloss.NewStyle()
		fa[f] = lipgloss.NewStyle()
	}
	return fs, fa
}

func TestCenterLabel_evenPadding(t *testing.T) {
	got := centerLabel("A", 5)
	want := "  A  "
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCenterLabel_oddPadding(t *testing.T) {
	got := centerLabel("AB", 5)
	// total=3, left=1, right=2 → " AB  "
	want := " AB  "
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCenterLabel_exactFit(t *testing.T) {
	got := centerLabel("ABC", 3)
	want := "ABC"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCenterLabel_overflow(t *testing.T) {
	got := centerLabel("ABCD", 3)
	want := "ABCD"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestCenterLabel_wideChars(t *testing.T) {
	got := centerLabel("⌘", 5)
	// ⌘ has visual width 2, so total=3, left=1, right=2
	wv := lipgloss.Width(got)
	if wv != 5 {
		t.Errorf("expected visual width 5, got %d for %q", wv, got)
	}
}

func TestApplyLayout_withMap(t *testing.T) {
	layoutMap := map[string]string{"A": "B", "C": "D"}
	keys := []Key{
		{Label: "A", Width: 3},
		{Label: "C", Width: 3},
		{Label: "E", Width: 3},
	}
	result := applyLayout(keys, layoutMap)
	if result[0].Label != "B" {
		t.Errorf("expected B, got %q", result[0].Label)
	}
	if result[1].Label != "D" {
		t.Errorf("expected D, got %q", result[1].Label)
	}
	if result[2].Label != "E" {
		t.Errorf("expected E (unchanged), got %q", result[2].Label)
	}
}

func TestApplyLayout_nilMap(t *testing.T) {
	keys := []Key{{Label: "A", Width: 3}}
	result := applyLayout(keys, nil)
	if result[0].Label != "A" {
		t.Errorf("expected A unchanged, got %q", result[0].Label)
	}
}

func TestApplyLayout_originalUnmodified(t *testing.T) {
	layoutMap := map[string]string{"A": "B"}
	keys := []Key{{Label: "A", Width: 3}}
	applyLayout(keys, layoutMap)
	if keys[0].Label != "A" {
		t.Errorf("applyLayout must not modify original slice")
	}
}

func TestTopLine_format(t *testing.T) {
	keys := []Key{
		{Label: "A", Width: 3},
		{Label: "B", Width: 3},
	}
	got := topLine(keys)
	want := ",---,---,"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestBotLine_format(t *testing.T) {
	keys := []Key{
		{Label: "A", Width: 3},
		{Label: "B", Width: 3},
	}
	got := botLine(keys)
	want := "'---'---'"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRender_unknownSize(t *testing.T) {
	fs, fa := emptyStyles()
	got := Render("qwerty", 999, "ansi", nil, fs, fa)
	if got != "" {
		t.Errorf("expected empty for unknown size, got %q", got)
	}
}

func TestRender_unknownLayout(t *testing.T) {
	fs, fa := emptyStyles()
	got := Render("nonexistent", 60, "ansi", nil, fs, fa)
	if got == "" {
		t.Error("expected non-empty rendering for unknown layout (falls back to labels)")
	}
}

func TestRender_size60_hasRows(t *testing.T) {
	fs, fa := emptyStyles()
	got := Render("qwerty", 60, "ansi", nil, fs, fa)
	lines := strings.Split(got, "\n")
	if len(lines) < 5 {
		t.Errorf("expected at least 5 lines, got %d", len(lines))
	}
}

func TestRender_size60_startsWithComma(t *testing.T) {
	fs, fa := emptyStyles()
	got := Render("qwerty", 60, "ansi", nil, fs, fa)
	lines := strings.Split(got, "\n")
	if len(lines) == 0 || !strings.HasPrefix(lines[0], ",") {
		t.Errorf("first line should start with ',', got %q", lines[0])
	}
}

func TestRender_allSizesRender(t *testing.T) {
	fs, fa := emptyStyles()
	for size := range standards.SizesANSI {
		t.Run(fmt.Sprintf("%d%%", size), func(t *testing.T) {
			got := Render("qwerty", size, "ansi", nil, fs, fa)
			if got == "" {
				t.Errorf("size %d produced empty output", size)
			}
		})
	}
}

func TestRender_allSizesStandard(t *testing.T) {
	fs, fa := emptyStyles()
	for size := range standards.SizesISO {
		t.Run(fmt.Sprintf("%d%%", size), func(t *testing.T) {
			got := Render("qwerty", size, "iso", nil, fs, fa)
			if got == "" {
				t.Errorf("ISO size %d produced empty output", size)
			}
		})
	}
}

func TestRender_Standard_hasEnterOnRow2(t *testing.T) {
	fs, fa := emptyStyles()
	got := Render("qwerty", 60, "iso", nil, fs, fa)
	if !strings.Contains(got, "Ent") {
		t.Error("ISO keyboard should have Enter key on row 2")
	}
}

func TestRender_Standard_hasHashKey(t *testing.T) {
	fs, fa := emptyStyles()
	got := Render("qwerty", 60, "iso", nil, fs, fa)
	if !strings.Contains(got, "#") {
		t.Error("ISO keyboard should have # key on row 3")
	}
}

func TestRender_size80_hasGaps(t *testing.T) {
	fs, fa := emptyStyles()
	got := Render("qwerty", 80, "ansi", nil, fs, fa)
	if !strings.Contains(got, "  ") {
		t.Error("size 80 should have gap spaces")
	}
}

func TestRender_Standard_allSizesDistinct(t *testing.T) {
	fs, fa := emptyStyles()
	for size := range standards.SizesANSI {
		t.Run(fmt.Sprintf("%d%%", size), func(t *testing.T) {
			ansi := Render("qwerty", size, "ansi", nil, fs, fa)
			std := Render("qwerty", size, "iso", nil, fs, fa)
			if ansi == std {
				t.Errorf("ISO size %d should differ from ANSI", size)
			}
		})
	}
}

func TestQWERTYUKShiftMap(t *testing.T) {
	fs, fa := emptyStyles()
	shifted := map[uint16]bool{base.KEY_LEFTSHIFT: true}
	got := Render("qwerty uk", 60, "ansi", shifted, fs, fa)
	if !strings.Contains(got, "\"") {
		t.Error("UK layout should show \" when shift+2 is held")
	}
	if strings.Contains(got, "#") {
		t.Error("UK layout should not show # when shift is held (should be £ on 3)")
	}
	if !strings.Contains(got, "£") {
		t.Error("UK layout should show £ when shift+3 is held")
	}
}

func TestQWERTYUKAltGr(t *testing.T) {
	fs, fa := emptyStyles()
	altGr := map[uint16]bool{base.KEY_RIGHTALT: true}
	got := Render("qwerty uk", 60, "ansi", altGr, fs, fa)
	if !strings.Contains(got, "Á") {
		t.Error("UK layout should show Á when AltGr+A is held")
	}
	if !strings.Contains(got, "€") {
		t.Error("UK layout should show € when AltGr+4 is held")
	}
}

func TestQWERTYUKShiftAltGr(t *testing.T) {
	fs, fa := emptyStyles()
	both := map[uint16]bool{base.KEY_LEFTSHIFT: true, base.KEY_RIGHTALT: true}
	got := Render("qwerty uk", 60, "ansi", both, fs, fa)
	if !strings.Contains(got, "Á") {
		t.Error("UK layout should show Á when Shift+AltGr+A is held")
	}
	if strings.Contains(got, "á") {
		t.Error("UK layout should not show lowercase á when Shift+AltGr+A is held")
	}
}

func TestDvorakUKShiftMap(t *testing.T) {
	fs, fa := emptyStyles()
	shifted := map[uint16]bool{base.KEY_LEFTSHIFT: true}
	got := Render("dvorak uk", 60, "ansi", shifted, fs, fa)
	if !strings.Contains(got, "\"") {
		t.Error("Dvorak UK layout should show \" when shift+2 is held")
	}
	if !strings.Contains(got, "£") {
		t.Error("Dvorak UK layout should show £ when shift+3 is held")
	}
}

func TestDvorakUKAltGr(t *testing.T) {
	fs, fa := emptyStyles()
	altGr := map[uint16]bool{base.KEY_RIGHTALT: true}
	got := Render("dvorak uk", 60, "ansi", altGr, fs, fa)
	if !strings.Contains(got, "Á") {
		t.Error("Dvorak UK layout should show Á when AltGr+A (QWERTY position) is held")
	}
	if !strings.Contains(got, "€") {
		t.Error("Dvorak UK layout should show € when AltGr+4 is held")
	}
}

func TestQWERTYUK_BacktickShift(t *testing.T) {
	fs, fa := emptyStyles()
	shifted := map[uint16]bool{base.KEY_LEFTSHIFT: true}
	got := Render("qwerty uk", 60, "ansi", shifted, fs, fa)
	if !strings.Contains(got, "¬") {
		t.Error("UK layout should show ¬ when shift is held (shift+` = ¬)")
	}
}

func TestJIS_allSizesHaveJISKeys(t *testing.T) {
	fs, fa := emptyStyles()
	for size := range standards.SizesJIS {
		t.Run(fmt.Sprintf("%d%%", size), func(t *testing.T) {
			got := Render("qwerty", size, "jis", nil, fs, fa)
			if got == "" {
				t.Errorf("JIS size %d produced empty output", size)
			}
			jisKey := "全半"
			if size == 75 || size == 96 {
				jisKey = "全"
			}
			jisSingles := []string{jisKey, "無", "変", "仮"}
			for _, s := range jisSingles {
				if !strings.Contains(got, s) {
					t.Errorf("JIS size %d missing %q", size, s)
				}
			}
		})
	}
}

func TestQWERTYUK_BacktickAltGr(t *testing.T) {
	fs, fa := emptyStyles()
	altGr := map[uint16]bool{base.KEY_RIGHTALT: true}
	got := Render("qwerty uk", 60, "ansi", altGr, fs, fa)
	if !strings.Contains(got, "¦") {
		t.Error("UK layout should show ¦ when AltGr is held (AltGr+` = ¦)")
	}
}

func TestBacktickShift_NonUK(t *testing.T) {
	fs, fa := emptyStyles()
	shifted := map[uint16]bool{base.KEY_LEFTSHIFT: true}
	got := Render("qwerty", 60, "ansi", shifted, fs, fa)
	if !strings.Contains(got, "~") {
		t.Error("US layout should show ~ when shift+` is held")
	}
}
