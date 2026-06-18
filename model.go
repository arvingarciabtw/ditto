package main

import (
	help "charm.land/bubbles/v2/help"
	list "charm.land/bubbles/v2/list"
)

type Model struct {
	layoutList     list.Model
	sizeList       list.Model
	activeLayout   string
	activeSize     int
	showLayoutList bool
	showSizeList   bool
	showInfoBar    bool
	helpModel      help.Model
	pressedKeys    map[uint16]bool
}

type GlobalKeyMsg struct {
	Code uint16
	Down bool
}

func getInitModel() Model {
	cfg := loadConfig()
	initModel := Model{
		layoutList:     list.New(keyboardLayoutItems, list.NewDefaultDelegate(), 0, 0),
		sizeList:       list.New(keyboardSizeItems, list.NewDefaultDelegate(), 0, 0),
		activeLayout:   cfg.ActiveLayout,
		activeSize:     cfg.ActiveSize,
		showLayoutList: false,
		showSizeList:   false,
		showInfoBar:    true,
		helpModel:      help.New(),
		pressedKeys:    make(map[uint16]bool),
	}
	initModel.layoutList.Title = "Layouts"
	initModel.layoutList.KeyMap.Quit.SetKeys("q")
	initModel.sizeList.Title = "Sizes"
	initModel.sizeList.KeyMap.Quit.SetKeys("q")

	return initModel
}
