// Package tui implements the Bubble Tea model, update, view loop
// along with the TUI styling and overlay components.
package tui

import (
	"github.com/arvingarciabtw/ditto/internal/config"
	"github.com/arvingarciabtw/ditto/internal/keyboard"
	"github.com/arvingarciabtw/ditto/internal/tui/components"
)

type Model struct {
	activeLayout     string
	activeSize       int
	activeStandard   string
	locked           bool
	layoutList       components.ListModel
	sizeList         components.ListModel
	standardList     components.ListModel
	quitDialog       components.DialogModel
	showLayoutList   bool
	showSizeList     bool
	showStandardList bool
	showQuitDialog   bool
	showAllInfo      bool
	pressedKeys      map[uint16]bool
	capsLock         bool
	kanaKeyHeld      bool
	kanaActive       bool
	hangeulKeyHeld   bool
	hangeulActive    bool
	terminalWidth    int
	terminalHeight   int
}

func InitModel(cfg config.Config) Model {
	layoutList := components.ListModel{
		Items:        keyboard.LayoutListItems,
		Selected:     0,
		Title:        "Layouts",
		AccentColor:  LayoutColor,
		VisibleCount: 3,
	}
	for i, item := range layoutList.Items {
		if item == cfg.ActiveLayout {
			layoutList.Selected = i
			break
		}
	}

	sizeList := components.ListModel{
		Items:        keyboard.LayoutSizeItems,
		Selected:     0,
		Title:        "Sizes",
		AccentColor:  SizeColor,
		VisibleCount: 3,
	}
	for i, s := range keyboard.Sizes {
		if s == cfg.ActiveSize {
			sizeList.Selected = i
			break
		}
	}

	standardList := components.ListModel{
		Items:        keyboard.StandardListItems,
		Selected:     0,
		Title:        "Standards",
		AccentColor:  StandardColor,
		VisibleCount: 3,
	}
	for i, item := range standardList.Items {
		if item == cfg.ActiveStandard {
			standardList.Selected = i
			break
		}
	}

	showAllInfo := true
	if cfg.ShowAllInfo != nil {
		showAllInfo = *cfg.ShowAllInfo
	}

	return Model{
		layoutList:       layoutList,
		sizeList:         sizeList,
		standardList:     standardList,
		quitDialog:       components.DialogModel{AccentColor: QuitColor},
		activeLayout:     cfg.ActiveLayout,
		activeSize:       cfg.ActiveSize,
		activeStandard:   cfg.ActiveStandard,
		locked:           cfg.Locked,
		showLayoutList:   false,
		showSizeList:     false,
		showStandardList: false,
		showAllInfo:      showAllInfo,
		pressedKeys:      make(map[uint16]bool),
	}
}

func (m Model) saveConfig() config.Config {
	v := m.showAllInfo
	return config.Config{
		ActiveLayout:   m.activeLayout,
		ActiveSize:     m.activeSize,
		ActiveStandard: m.activeStandard,
		Locked:         m.locked,
		ShowAllInfo:    &v,
	}
}
