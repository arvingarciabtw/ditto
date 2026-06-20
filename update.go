package main

import (
	"strconv"
	"strings"

	tea "charm.land/bubbletea/v2"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		/*
			These shortcuts bypass the overlay dispatch below so the user
			can switch between layout and size lists directly, without
			having to close the current overlay first.
		*/
		switch msg.String() {
		case "ctrl+shift+l":
			m.showLayoutList = !m.showLayoutList
			m.showSizeList = false
			return m, nil
		case "ctrl+shift+s":
			m.showSizeList = !m.showSizeList
			m.showLayoutList = false
			return m, nil
		}

		switch {
		case m.showLayoutList:
			return m.handleLayoutListUpdate(msg)
		case m.showSizeList:
			return m.handleSizeListUpdate(msg)
		case m.showQuitDialog:
			return m.handleQuitDialogUpdate(msg)
		default:
			return m.handleGlobalKeys(msg)
		}
	case globalKeyMsg:
		m.pressedKeys[msg.code] = msg.down
	case tea.WindowSizeMsg:
		m.terminalWidth = msg.Width
		m.terminalHeight = msg.Height
		m.help.SetWidth(msg.Width)
	}

	return m, nil
}

func (m Model) handleLayoutListUpdate(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	var action listAction
	m.layoutList, action = m.layoutList.Update(msg)

	switch action {

	case listConfirm:
		m.activeLayout = strings.ToLower(m.layoutList.items[m.layoutList.selected])
		m.showLayoutList = false
		saveConfig(config{ActiveLayout: m.activeLayout, ActiveSize: m.activeSize})
	case listCancel:
		m.showLayoutList = false
	}

	return m, nil
}

func (m Model) handleSizeListUpdate(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	var action listAction
	m.sizeList, action = m.sizeList.Update(msg)

	switch action {

	case listConfirm:
		sizeStr := strings.TrimSuffix(m.sizeList.items[m.sizeList.selected], "%")
		if size, err := strconv.Atoi(sizeStr); err == nil {
			m.activeSize = size
		}
		m.showSizeList = false
		saveConfig(config{ActiveLayout: m.activeLayout, ActiveSize: m.activeSize})
	case listCancel:
		m.showSizeList = false
	}

	return m, nil
}

func (m Model) handleQuitDialogUpdate(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	var action dialogAction
	m.quitDialog, action = m.quitDialog.Update(msg)

	switch action {

	case dialogConfirm:
		return m, tea.Quit
	case dialogCancel:
		m.showQuitDialog = false
	}

	return m, nil
}

func (m Model) handleGlobalKeys(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {

	case "ctrl+shift+h":
		m.showAllInfo = !m.showAllInfo
	case "?":
		m.help.ShowAll = !m.help.ShowAll
	case "q", "esc":
		m.showQuitDialog = true
		m.quitDialog.selected = 0
	case "ctrl+c":
		return m, tea.Quit
	}

	return m, nil
}
