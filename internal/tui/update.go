package tui

import (
	"strconv"
	"strings"

	tea "charm.land/bubbletea/v2"

	"github.com/arvingarciabtw/ditto/internal/config"
	"github.com/arvingarciabtw/ditto/internal/input"
	basepkg "github.com/arvingarciabtw/ditto/internal/keyboard/base"
	"github.com/arvingarciabtw/ditto/internal/tui/components"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {
		case "l":
			if !m.locked {
				m.showLayoutList = !m.showLayoutList
				m.showSizeList = false
				m.showStandardList = false
			}
			return m, nil
		case "s":
			if !m.locked {
				m.showSizeList = !m.showSizeList
				m.showLayoutList = false
				m.showStandardList = false
			}
			return m, nil
		case "d":
			if !m.locked {
				m.showStandardList = !m.showStandardList
				m.showLayoutList = false
				m.showSizeList = false
			}
			return m, nil
		case "h":
			if !m.locked {
				m.showAllInfo = !m.showAllInfo
				_ = config.SaveConfig(m.saveConfig())
			}
			return m, nil
		}

		switch {
		case m.showLayoutList:
			return m.handleLayoutListUpdate(msg)
		case m.showSizeList:
			return m.handleSizeListUpdate(msg)
		case m.showStandardList:
			return m.handleStandardListUpdate(msg)
		case m.showQuitDialog:
			return m.handleQuitDialogUpdate(msg)
		default:
			return m.handleGlobalKeys(msg)
		}
	case input.KeyMsg:
		m.pressedKeys[msg.Code] = msg.Down
		if msg.Code == basepkg.KEY_CAPSLOCK && msg.Down {
			m.capsLock = !m.capsLock
		}
		if msg.Code == basepkg.KEY_KATAKANAHIRAGANA {
			m.kanaKeyHeld = msg.Down
		}
		if msg.Code == basepkg.KEY_HANGEUL {
			m.hangeulKeyHeld = msg.Down
		}
	case tea.WindowSizeMsg:
		m.terminalWidth = msg.Width
		m.terminalHeight = msg.Height
	}

	m.pressedKeys[basepkg.KEY_CAPSLOCK] = m.pressedKeys[basepkg.KEY_CAPSLOCK] || m.capsLock
	m.pressedKeys[basepkg.KEY_KATAKANAHIRAGANA] = m.kanaKeyHeld || m.kanaActive
	m.pressedKeys[basepkg.KEY_HANGEUL] = m.hangeulKeyHeld || m.hangeulActive

	return m, nil
}

func (m Model) handleLayoutListUpdate(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	var action components.ListAction
	m.layoutList, action = m.layoutList.Update(msg)

	switch action {

	case components.ListConfirm:
		m.activeLayout = strings.ToLower(m.layoutList.Items[m.layoutList.Selected])
		if strings.HasSuffix(m.activeLayout, " uk") || m.activeLayout == "qwertz" {
			m.activeStandard = "iso"
		}
		m.showLayoutList = false
		_ = config.SaveConfig(m.saveConfig())
		return m, nil
	case components.ListCancel:
		m.showLayoutList = false
		return m, nil
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
		_ = config.SaveConfig(m.saveConfig())
		return m, nil
	case components.ListCancel:
		m.showSizeList = false
		return m, nil
	}

	return m, nil
}

func (m Model) handleStandardListUpdate(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	var action components.ListAction
	m.standardList, action = m.standardList.Update(msg)

	switch action {

	case components.ListConfirm:
		m.activeStandard = m.standardList.Items[m.standardList.Selected]
		m.showStandardList = false
		m.kanaActive = false
		m.hangeulActive = false
		_ = config.SaveConfig(m.saveConfig())
		return m, nil
	case components.ListCancel:
		m.showStandardList = false
		return m, nil
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
		return m, nil
	}

	return m, nil
}

func (m Model) handleGlobalKeys(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {

	case "q", "esc":
		m.showQuitDialog = true
		m.quitDialog.Selected = 0
	case "ctrl+c":
		return m, tea.Quit
	case "c":
		if !m.locked {
			switch m.activeStandard {
			case "jis":
				m.kanaActive = !m.kanaActive
			case "ks":
				m.hangeulActive = !m.hangeulActive
			}
		}
	}

	return m, nil
}
