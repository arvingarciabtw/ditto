package components

import (
	"image/color"
	"strings"
	"testing"

	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

func testList() ListModel {
	return ListModel{
		Items:        []string{"a", "b", "c"},
		Selected:     0,
		Title:        "Test",
		AccentColor:  color.RGBA{R: 255, G: 0, B: 0, A: 255},
		VisibleCount: 3,
	}
}

func TestListModel_initialSelection(t *testing.T) {
	l := testList()
	if l.Selected != 0 {
		t.Errorf("expected Selected=0, got %d", l.Selected)
	}
}

func TestListModel_navigateDown(t *testing.T) {
	l := testList()
	l, _ = l.Update(tea.KeyPressMsg{Code: 'j'})
	if l.Selected != 1 {
		t.Errorf("expected Selected=1, got %d", l.Selected)
	}
}

func TestListModel_navigateUp(t *testing.T) {
	l := testList()
	l.Selected = 2
	l, _ = l.Update(tea.KeyPressMsg{Code: 'k'})
	if l.Selected != 1 {
		t.Errorf("expected Selected=1, got %d", l.Selected)
	}
}

func TestListModel_staysAtTop(t *testing.T) {
	l := testList()
	l, _ = l.Update(tea.KeyPressMsg{Code: tea.KeyUp})
	_, action := l.Update(tea.KeyPressMsg{Code: tea.KeyUp})
	if action == ListCancel || action == ListConfirm {
		t.Errorf("up at top should not confirm or cancel, got action=%d", action)
	}
}

func TestListModel_staysAtBottom(t *testing.T) {
	l := testList()
	l.Selected = 2
	_, action := l.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	if action == ListCancel || action == ListConfirm {
		t.Errorf("down at bottom should not confirm or cancel, got action=%d", action)
	}
}

func TestListModel_confirmEnter(t *testing.T) {
	l := testList()
	_, action := l.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	if action != ListConfirm {
		t.Errorf("expected ListConfirm, got %d", action)
	}
}

func TestListModel_cancelQ(t *testing.T) {
	l := testList()
	_, action := l.Update(tea.KeyPressMsg{Code: 'q'})
	if action != ListCancel {
		t.Errorf("expected ListCancel, got %d", action)
	}
}

func TestListModel_cancelEsc(t *testing.T) {
	l := testList()
	_, action := l.Update(tea.KeyPressMsg{Code: tea.KeyEscape})
	if action != ListCancel {
		t.Errorf("expected ListCancel, got %d", action)
	}
}

func TestListModel_viewContainsTitle(t *testing.T) {
	l := testList()
	v := l.View(lipgloss.NewStyle())
	if !strings.Contains(v, "Test") {
		t.Errorf("view should contain title 'Test', got %q", v)
	}
}

func TestListModel_viewContainsHelp(t *testing.T) {
	l := testList()
	v := l.View(lipgloss.NewStyle())
	if !strings.Contains(v, "enter") {
		t.Errorf("view should contain help text, got %q", v)
	}
}

func TestListModel_emptyItems(t *testing.T) {
	l := ListModel{
		Items:        nil,
		Selected:     0,
		Title:        "Empty",
		AccentColor:  color.RGBA{R: 0, G: 0, B: 255, A: 255},
		VisibleCount: 3,
	}
	_, action := l.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	if action == ListConfirm || action == ListCancel {
		t.Error("navigation on empty list should not confirm or cancel")
	}
}
