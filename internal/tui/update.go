package tui

import (
	"strconv"
	"strings"

	tea "charm.land/bubbletea/v2"

	"github.com/arvingarciabtw/ditto/internal/config"
	"github.com/arvingarciabtw/ditto/internal/evdev"
	"github.com/arvingarciabtw/ditto/internal/keyboard"
	"github.com/arvingarciabtw/ditto/internal/tui/components"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+shift+l":
			m.showLayoutList = !m.showLayoutList
			m.showSizeList = false
			return m, nil
		case "ctrl+shift+s":
			m.showSizeList = !m.showSizeList
			m.showLayoutList = false
			return m, nil
		case "ctrl+shift+d":
			if m.activeStandard == keyboard.ANSI {
				m.activeStandard = keyboard.ISO
			} else {
				m.activeStandard = keyboard.ANSI
			}
			config.SaveConfig(config.Config{ActiveLayout: m.activeLayout, ActiveSize: m.activeSize, ActiveStandard: m.activeStandard})
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
	case evdev.KeyMsg:
		m.pressedKeys[msg.Code] = msg.Down
		if msg.Code == 58 && msg.Down {
			m.capsLock = !m.capsLock
		}
	case tea.WindowSizeMsg:
		m.terminalWidth = msg.Width
		m.terminalHeight = msg.Height
	}

	m.pressedKeys[58] = m.pressedKeys[58] || m.capsLock

	return m, nil
}

func (m Model) handleLayoutListUpdate(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	var action components.ListAction
	m.layoutList, action = m.layoutList.Update(msg)

	switch action {

	case components.ListConfirm:
		m.activeLayout = strings.ToLower(m.layoutList.Items[m.layoutList.Selected])
		if strings.HasSuffix(m.activeLayout, " uk") {
			m.activeStandard = keyboard.ISO
		}
		m.showLayoutList = false
		config.SaveConfig(config.Config{ActiveLayout: m.activeLayout, ActiveSize: m.activeSize, ActiveStandard: m.activeStandard})
	case components.ListCancel:
		m.showLayoutList = false
	}

	return m, nil
}

func (m Model) handleSizeListUpdate(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	var action components.ListAction
	m.sizeList, action = m.sizeList.Update(msg)

	switch action {

	case components.ListConfirm:
		sizeStr := strings.TrimSuffix(m.sizeList.Items[m.sizeList.Selected], "%")
		if size, err := strconv.Atoi(sizeStr); err == nil {
			m.activeSize = size
		}
		m.showSizeList = false
		config.SaveConfig(config.Config{ActiveLayout: m.activeLayout, ActiveSize: m.activeSize, ActiveStandard: m.activeStandard})
	case components.ListCancel:
		m.showSizeList = false
	}

	return m, nil
}

func (m Model) handleQuitDialogUpdate(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	var action components.DialogAction
	m.quitDialog, action = m.quitDialog.Update(msg)

	switch action {

	case components.DialogConfirm:
		return m, tea.Quit
	case components.DialogCancel:
		m.showQuitDialog = false
	}

	return m, nil
}

func (m Model) handleGlobalKeys(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {

	case "ctrl+shift+h":
		m.showAllInfo = !m.showAllInfo
	case "q", "esc":
		m.showQuitDialog = true
		m.quitDialog.Selected = 0
	case "ctrl+c":
		return m, tea.Quit
	}

	return m, nil
}
