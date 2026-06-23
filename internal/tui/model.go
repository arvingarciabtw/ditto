// Package tui implements the Bubble Tea model, update, view loop
// along with the TUI styling and overlay components.
package tui

import (
	"fmt"

	"github.com/arvingarciabtw/ditto/internal/config"
	"github.com/arvingarciabtw/ditto/internal/keyboard"
	"github.com/arvingarciabtw/ditto/internal/tui/components"
)

type Model struct {
	activeLayout   string
	activeSize     int
	layoutList     components.ListModel
	sizeList       components.ListModel
	quitDialog     components.DialogModel
	showLayoutList bool
	showSizeList   bool
	showQuitDialog bool
	showAllInfo    bool
	pressedKeys    map[uint16]bool
	terminalWidth  int
	terminalHeight int
}

func InitModel() Model {
	cfg := config.LoadConfig()

	layoutList := components.ListModel{
		Items:       keyboard.LayoutListItems,
		Selected:    0,
		Title:       "Layouts",
		AccentColor: LayoutColor,
	}
	for i, item := range layoutList.Items {
		if item == cfg.ActiveLayout {
			layoutList.Selected = i
			break
		}
	}

	sizeList := components.ListModel{
		Items:       keyboard.LayoutSizeItems,
		Selected:    0,
		Title:       "Sizes",
		AccentColor: SizeColor,
	}
	for i, item := range sizeList.Items {
		if item == fmt.Sprintf("%d%%", cfg.ActiveSize) {
			sizeList.Selected = i
			break
		}
	}

	return Model{
		layoutList:     layoutList,
		sizeList:       sizeList,
		quitDialog:     components.DialogModel{AccentColor: QuitColor},
		activeLayout:   cfg.ActiveLayout,
		activeSize:     cfg.ActiveSize,
		showLayoutList: false,
		showSizeList:   false,
		showAllInfo:    true,
		pressedKeys:    make(map[uint16]bool),
	}
}
