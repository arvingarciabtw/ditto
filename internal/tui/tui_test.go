package tui

import (
	"testing"
	"time"

	tea "charm.land/bubbletea/v2"

	"github.com/arvingarciabtw/ditto/internal/config"
	"github.com/arvingarciabtw/ditto/internal/input"
	basepkg "github.com/arvingarciabtw/ditto/internal/keyboard/base"
)

func testModel(t *testing.T) Model {
	t.Helper()
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	return InitModel(config.Default())
}

func TestModel_init(t *testing.T) {
	m := testModel(t)
	cmd := m.Init()
	if cmd != nil {
		t.Errorf("expected nil cmd from Init, got %v", cmd)
	}
}

func TestModel_initHasLayout(t *testing.T) {
	m := testModel(t)
	if m.activeLayout != "qwerty" {
		t.Errorf("expected default layout qwerty, got %q", m.activeLayout)
	}
	if m.activeSize != 75 {
		t.Errorf("expected default size 75, got %d", m.activeSize)
	}
}

func updateModel(t *testing.T, m Model, msg tea.Msg) Model {
	t.Helper()
	result, _ := m.Update(msg)
	model, ok := result.(Model)
	if !ok {
		t.Fatalf("unexpected type %T from Update", result)
	}
	return model
}

func TestModel_windowSize(t *testing.T) {
	m := testModel(t)
	m = updateModel(t, m, tea.WindowSizeMsg{Width: 100, Height: 50})
	if m.terminalWidth != 100 {
		t.Errorf("expected width 100, got %d", m.terminalWidth)
	}
	if m.terminalHeight != 50 {
		t.Errorf("expected height 50, got %d", m.terminalHeight)
	}
}

func TestModel_keyPressDown(t *testing.T) {
	m := testModel(t)
	m = updateModel(t, m, input.KeyMsg{Code: basepkg.KEY_A, Down: true})
	if !m.pressedKeys[basepkg.KEY_A] {
		t.Error("expected key 30 to be pressed")
	}
}

func TestModel_keyPressUp(t *testing.T) {
	m := testModel(t)
	m.pressedKeys[basepkg.KEY_A] = true
	m = updateModel(t, m, input.KeyMsg{Code: basepkg.KEY_A, Down: false})
	if m.pressedKeys[basepkg.KEY_A] {
		t.Error("expected key 30 to be released")
	}
}

func TestModel_toggleLayoutList(t *testing.T) {
	m := testModel(t)
	m = updateModel(t, m, tea.KeyPressMsg{Code: 'l'})
	if !m.showLayoutList {
		t.Error("expected showLayoutList to be true after l")
	}
}

func TestModel_toggleSizeList(t *testing.T) {
	m := testModel(t)
	m = updateModel(t, m, tea.KeyPressMsg{Code: 's'})
	if !m.showSizeList {
		t.Error("expected showSizeList to be true after ctrl+shift+s")
	}
}

func TestModel_toggleStandardList(t *testing.T) {
	m := testModel(t)
	m = updateModel(t, m, tea.KeyPressMsg{Code: 'd'})
	if !m.showStandardList {
		t.Error("expected showStandardList to be true after d")
	}
}

func TestModel_standardListClosesOthers(t *testing.T) {
	m := testModel(t)
	m.showLayoutList = true
	m = updateModel(t, m, tea.KeyPressMsg{Code: 'd'})
	if !m.showStandardList {
		t.Error("expected showStandardList to be true")
	}
	if m.showLayoutList {
		t.Error("expected showLayoutList to be false when standard opens")
	}
}

func TestModel_layoutListClosesSizeList(t *testing.T) {
	m := testModel(t)
	m = updateModel(t, m, tea.KeyPressMsg{Code: 's'})
	m = updateModel(t, m, tea.KeyPressMsg{Code: 'l'})
	if !m.showLayoutList {
		t.Error("expected showLayoutList to be true")
	}
	if m.showSizeList {
		t.Error("expected showSizeList to be closed when layout opens")
	}
}

func TestModel_openQuitDialog(t *testing.T) {
	m := testModel(t)
	m = updateModel(t, m, tea.KeyPressMsg{Text: "q", Code: 'q'})
	if !m.showQuitDialog {
		t.Error("expected showQuitDialog to be true after pressing q")
	}
}

func TestModel_quitCtrlC(t *testing.T) {
	m := testModel(t)
	_, cmd := m.Update(tea.KeyPressMsg{Mod: tea.ModCtrl, Code: 'c'})
	if cmd == nil {
		t.Error("expected tea.Quit cmd from ctrl+c")
	}
}

func TestModel_toggleInfo(t *testing.T) {
	m := testModel(t)
	if !m.showAllInfo {
		t.Error("expected showAllInfo to start true")
	}
	m = updateModel(t, m, tea.KeyPressMsg{Code: 'h'})
	if m.showAllInfo {
		t.Error("expected showAllInfo to be false after toggle")
	}
}

func TestModel_standardListConfirm(t *testing.T) {
	m := testModel(t)
	m.showStandardList = true
	m.standardList.Selected = 1
	m = updateModel(t, m, tea.KeyPressMsg{Code: tea.KeyEnter})
	if m.activeStandard != "iso" {
		t.Errorf("expected activeStandard iso, got %q", m.activeStandard)
	}
	if m.showStandardList {
		t.Error("expected standardList to close after confirm")
	}
}

func TestModel_escClosesOverlay(t *testing.T) {
	m := testModel(t)
	m.showLayoutList = true
	m = updateModel(t, m, tea.KeyPressMsg{Code: tea.KeyEscape})
	if m.showLayoutList {
		t.Error("expected layoutList overlay to close on esc")
	}
}

func TestModel_locked_blocksL(t *testing.T) {
	m := testModel(t)
	m.locked = true
	m = updateModel(t, m, tea.KeyPressMsg{Code: 'l'})
	if m.showLayoutList {
		t.Error("expected showLayoutList to be false when locked")
	}
}

func TestModel_locked_blocksS(t *testing.T) {
	m := testModel(t)
	m.locked = true
	m = updateModel(t, m, tea.KeyPressMsg{Code: 's'})
	if m.showSizeList {
		t.Error("expected showSizeList to be false when locked")
	}
}

func TestModel_locked_blocksC(t *testing.T) {
	m := testModel(t)
	m.locked = true
	m.activeStandard = "jis"
	m = updateModel(t, m, tea.KeyPressMsg{Code: 'c'})
	if m.kanaActive {
		t.Error("expected kanaActive to be false when locked")
	}
}

func TestModel_locked_blocksH(t *testing.T) {
	m := testModel(t)
	m.locked = true
	m = updateModel(t, m, tea.KeyPressMsg{Code: 'h'})
	if !m.showAllInfo {
		t.Error("expected showAllInfo to stay true when locked")
	}
}

func TestModel_toggleKeycastMode(t *testing.T) {
	m := testModel(t)
	m = updateModel(t, m, tea.KeyPressMsg{Code: 'm'})
	if !m.showModeList {
		t.Error("expected showModeList to be true after m")
	}
	if m.keycastMode {
		t.Error("expected keycastMode to still be false")
	}
	m = updateModel(t, m, tea.KeyPressMsg{Code: tea.KeyDown})
	m = updateModel(t, m, tea.KeyPressMsg{Code: tea.KeyEnter})
	if !m.keycastMode {
		t.Error("expected keycastMode to be true after selecting Keycast")
	}
	if m.showModeList {
		t.Error("expected showModeList to be false after confirm")
	}
	m = updateModel(t, m, tea.KeyPressMsg{Code: 'm'})
	if !m.showModeList {
		t.Error("expected showModeList to be true after second m")
	}
	m = updateModel(t, m, tea.KeyPressMsg{Code: tea.KeyUp})
	m = updateModel(t, m, tea.KeyPressMsg{Code: tea.KeyEnter})
	if m.keycastMode {
		t.Error("expected keycastMode to be false after selecting Default")
	}
	if m.showModeList {
		t.Error("expected showModeList to be false after confirm")
	}
}

func TestModel_keycastKeyDown(t *testing.T) {
	m := testModel(t)
	m.keycastMode = true
	m = updateModel(t, m, input.KeyMsg{Code: basepkg.KEY_A, Down: true})
	if len(m.keycastKeys) != 1 || m.keycastKeys[0].label != "A" {
		t.Errorf("expected keycastKeys [A], got %v", m.keycastKeys)
	}
}

func TestModel_keycastKeyDownUp(t *testing.T) {
	m := testModel(t)
	m.keycastMode = true
	m = updateModel(t, m, input.KeyMsg{Code: basepkg.KEY_A, Down: true})
	m = updateModel(t, m, input.KeyMsg{Code: basepkg.KEY_A, Down: false})
	if len(m.keycastKeys) != 1 {
		t.Errorf("expected keycastKeys to have 1 key after key-up, got %v", m.keycastKeys)
	}
}

func TestModel_keycastFade(t *testing.T) {
	m := testModel(t)
	m.keycastMode = true
	m = updateModel(t, m, input.KeyMsg{Code: basepkg.KEY_A, Down: true})
	m.keycastKeys[0].pressedAt = time.Now().Add(-2 * time.Second)
	m = updateModel(t, m, keycastFadeMsg{version: 99})
	if len(m.keycastKeys) != 0 {
		t.Errorf("expected keycastKeys empty after fade, got %v", m.keycastKeys)
	}
}

func TestModel_keycastDuplicatePresses(t *testing.T) {
	m := testModel(t)
	m.keycastMode = true
	m = updateModel(t, m, input.KeyMsg{Code: basepkg.KEY_A, Down: true})
	v1 := m.keycastKeys[0].version
	m.keycastKeys[0].pressedAt = time.Now().Add(-2 * time.Second)
	m = updateModel(t, m, input.KeyMsg{Code: basepkg.KEY_A, Down: true})
	v2 := m.keycastKeys[1].version
	if len(m.keycastKeys) != 2 {
		t.Fatalf("expected 2 entries after two presses, got %v", m.keycastKeys)
	}
	if v1 == v2 {
		t.Error("expected different versions for each press")
	}
	m = updateModel(t, m, keycastFadeMsg{version: 99})
	if len(m.keycastKeys) != 1 || m.keycastKeys[0].version != v2 {
		t.Errorf("expected 1 entry with version %d after first fade, got %v", v2, m.keycastKeys)
	}
	m.keycastKeys[0].pressedAt = time.Now().Add(-2 * time.Second)
	m = updateModel(t, m, keycastFadeMsg{version: 99})
	if len(m.keycastKeys) != 0 {
		t.Errorf("expected empty after second fade, got %v", m.keycastKeys)
	}
}

func TestModel_keycastMaxFive(t *testing.T) {
	m := testModel(t)
	m.keycastMode = true
	for i := 0; i < 6; i++ {
		m = updateModel(t, m, input.KeyMsg{Code: basepkg.KEY_1 + uint16(i), Down: true})
	}
	if len(m.keycastKeys) != 5 {
		t.Errorf("expected 5 keys, got %d: %v", len(m.keycastKeys), m.keycastKeys)
	}
}

func TestModel_keycastToggleNotBlockedByLocked(t *testing.T) {
	m := testModel(t)
	m.locked = true
	m = updateModel(t, m, tea.KeyPressMsg{Code: 'm'})
	if !m.showModeList {
		t.Error("expected showModeList to be true even when locked")
	}
	m = updateModel(t, m, tea.KeyPressMsg{Code: tea.KeyDown})
	m = updateModel(t, m, tea.KeyPressMsg{Code: tea.KeyEnter})
	if !m.keycastMode {
		t.Error("expected keycastMode to be true even when locked")
	}
}

func TestModel_keycastQuitDialog(t *testing.T) {
	m := testModel(t)
	m.keycastMode = true
	m = updateModel(t, m, tea.KeyPressMsg{Text: "q", Code: 'q'})
	if !m.showQuitDialog {
		t.Error("expected quit dialog to open in keycast mode after q")
	}
}

func TestModel_keycastToggleClosesOverlays(t *testing.T) {
	m := testModel(t)
	m.showQuitDialog = true
	m = updateModel(t, m, tea.KeyPressMsg{Code: 'm'})
	if m.showQuitDialog {
		t.Error("expected quit dialog to close when opening mode list")
	}
	if !m.showModeList {
		t.Error("expected showModeList to be true when pressing m")
	}
}
