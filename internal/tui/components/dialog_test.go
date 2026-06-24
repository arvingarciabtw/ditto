package components

import (
	"image/color"
	"testing"

	tea "charm.land/bubbletea/v2"
)

func testDialog() DialogModel {
	return DialogModel{
		Selected:    0,
		AccentColor: color.RGBA{R: 255, G: 0, B: 0, A: 255},
	}
}

func TestDialogModel_initialSelection(t *testing.T) {
	d := testDialog()
	if d.Selected != 0 {
		t.Errorf("expected Selected=0, got %d", d.Selected)
	}
}

func TestDialogModel_selectDown(t *testing.T) {
	d := testDialog()
	d, _ = d.Update(tea.KeyPressMsg{Code: tea.KeyDown})
	if d.Selected != 1 {
		t.Errorf("expected Selected=1, got %d", d.Selected)
	}
}

func TestDialogModel_selectUp(t *testing.T) {
	d := testDialog()
	d.Selected = 1
	d, _ = d.Update(tea.KeyPressMsg{Code: tea.KeyUp})
	if d.Selected != 0 {
		t.Errorf("expected Selected=0, got %d", d.Selected)
	}
}

func TestDialogModel_confirmOnEnter(t *testing.T) {
	d := testDialog()
	_, action := d.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	if action != DialogConfirm {
		t.Errorf("expected DialogConfirm on Enter when selected=0, got %d", action)
	}
}

func TestDialogModel_cancelOnEnter(t *testing.T) {
	d := testDialog()
	d.Selected = 1
	_, action := d.Update(tea.KeyPressMsg{Code: tea.KeyEnter})
	if action != DialogCancel {
		t.Errorf("expected DialogCancel on Enter when selected=1, got %d", action)
	}
}

func TestDialogModel_cancelOnQ(t *testing.T) {
	d := testDialog()
	_, action := d.Update(tea.KeyPressMsg{Text: "q", Code: 'q'})
	if action != DialogCancel {
		t.Errorf("expected DialogCancel on q, got %d", action)
	}
}

func TestDialogModel_cancelOnEsc(t *testing.T) {
	d := testDialog()
	_, action := d.Update(tea.KeyPressMsg{Code: tea.KeyEscape})
	if action != DialogCancel {
		t.Errorf("expected DialogCancel on esc, got %d", action)
	}
}
