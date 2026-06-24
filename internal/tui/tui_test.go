package tui

import (
	"testing"

	tea "charm.land/bubbletea/v2"

	"github.com/arvingarciabtw/ditto/internal/evdev"
)

func testModel(t *testing.T) Model {
	t.Helper()
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	return InitModel()
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

func TestModel_windowSize(t *testing.T) {
	m := testModel(t)
	result, _ := m.Update(tea.WindowSizeMsg{Width: 100, Height: 50})
	m = result.(Model)
	if m.terminalWidth != 100 {
		t.Errorf("expected width 100, got %d", m.terminalWidth)
	}
	if m.terminalHeight != 50 {
		t.Errorf("expected height 50, got %d", m.terminalHeight)
	}
}

func TestModel_keyPressDown(t *testing.T) {
	m := testModel(t)
	result, _ := m.Update(evdev.KeyMsg{Code: 30, Down: true})
	m = result.(Model)
	if !m.pressedKeys[30] {
		t.Error("expected key 30 to be pressed")
	}
}

func TestModel_keyPressUp(t *testing.T) {
	m := testModel(t)
	m.pressedKeys[30] = true
	result, _ := m.Update(evdev.KeyMsg{Code: 30, Down: false})
	m = result.(Model)
	if m.pressedKeys[30] {
		t.Error("expected key 30 to be released")
	}
}

func TestModel_toggleLayoutList(t *testing.T) {
	m := testModel(t)
	result, _ := m.Update(tea.KeyPressMsg{Mod: tea.ModCtrl | tea.ModShift, Code: 'l'})
	m = result.(Model)
	if !m.showLayoutList {
		t.Error("expected showLayoutList to be true after ctrl+shift+l")
	}
}

func TestModel_toggleSizeList(t *testing.T) {
	m := testModel(t)
	result, _ := m.Update(tea.KeyPressMsg{Mod: tea.ModCtrl | tea.ModShift, Code: 's'})
	m = result.(Model)
	if !m.showSizeList {
		t.Error("expected showSizeList to be true after ctrl+shift+s")
	}
}

func TestModel_layoutListClosesSizeList(t *testing.T) {
	m := testModel(t)
	result, _ := m.Update(tea.KeyPressMsg{Mod: tea.ModCtrl | tea.ModShift, Code: 's'})
	m = result.(Model)
	result, _ = m.Update(tea.KeyPressMsg{Mod: tea.ModCtrl | tea.ModShift, Code: 'l'})
	m = result.(Model)
	if !m.showLayoutList {
		t.Error("expected showLayoutList to be true")
	}
	if m.showSizeList {
		t.Error("expected showSizeList to be closed when layout opens")
	}
}

func TestModel_openQuitDialog(t *testing.T) {
	m := testModel(t)
	result, _ := m.Update(tea.KeyPressMsg{Text: "q", Code: 'q'})
	m = result.(Model)
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
	result, _ := m.Update(tea.KeyPressMsg{Mod: tea.ModCtrl | tea.ModShift, Code: 'h'})
	m = result.(Model)
	if m.showAllInfo {
		t.Error("expected showAllInfo to be false after toggle")
	}
}

func TestModel_escClosesOverlay(t *testing.T) {
	m := testModel(t)
	m.showLayoutList = true
	result, _ := m.Update(tea.KeyPressMsg{Code: tea.KeyEscape})
	m = result.(Model)
	if m.showLayoutList {
		t.Error("expected layoutList overlay to close on esc")
	}
}
