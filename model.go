package main

import (
	"fmt"

	help "charm.land/bubbles/v2/help"
)

type Model struct {
	activeLayout   string
	activeSize     int
	help           help.Model
	layoutList     listModel
	sizeList       listModel
	quitDialog     dialogModel
	showLayoutList bool
	showSizeList   bool
	showQuitDialog bool
	showAllInfo    bool
	pressedKeys     map[uint16]bool
	terminalWidth   int
	terminalHeight  int
}

func initModel() Model {
	cfg := loadConfig()

	layoutList := listModel{
		items:       layoutListItems,
		selected:    0,
		title:       "Layouts",
		accentColor: layoutColor,
	}
	for i, item := range layoutList.items {
		if item == cfg.ActiveLayout {
			layoutList.selected = i
			break
		}
	}

	sizeList := listModel{
		items:       layoutSizeItems,
		selected:    0,
		title:       "Sizes",
		accentColor: sizeColor,
	}
	for i, item := range sizeList.items {
		if item == fmt.Sprintf("%d%%", cfg.ActiveSize) {
			sizeList.selected = i
			break
		}
	}

	initModel := Model{
		layoutList:     layoutList,
		sizeList:       sizeList,
		quitDialog:     dialogModel{accentColor: quitColor},
		activeLayout:   cfg.ActiveLayout,
		activeSize:     cfg.ActiveSize,
		showLayoutList: false,
		showSizeList:   false,
		showAllInfo:    true,
		help:           help.New(),
		pressedKeys:    make(map[uint16]bool),
	}
	initModel.help.Styles = help.Styles{
		FullKey:        statusBarStyle,
		FullDesc:       statusBarStyle,
		FullSeparator:  statusBarStyle,
		ShortKey:       statusBarStyle,
		ShortDesc:      statusBarStyle,
		ShortSeparator: statusBarStyle,
		Ellipsis:       statusBarStyle,
	}
	initModel.help.FullSeparator = initModel.help.ShortSeparator

	return initModel
}
