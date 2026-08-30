package tui

import (
	"strconv"
	"strings"
	"time"

	tea "charm.land/bubbletea/v2"

	"github.com/arvingarciabtw/ditto/internal/config"
	"github.com/arvingarciabtw/ditto/internal/input"
	"github.com/arvingarciabtw/ditto/internal/keyboard"
	basepkg "github.com/arvingarciabtw/ditto/internal/keyboard/base"
	"github.com/arvingarciabtw/ditto/internal/tui/components"
)

/*
keycastFadeMsg asks the update loop to remove expired keycast entries after a
key press without blocking event processing.
*/
type keycastFadeMsg struct {
	version int
}

/*
Init satisfies tea.Model; Ditto has no asynchronous command to start because
keyboard capture is owned by the command entrypoint.
*/
func (m Model) Init() tea.Cmd {
	return nil
}

/*
Update routes terminal, keyboard, and timer messages into model transitions so
Bubble Tea remains the single owner of mutable UI state.
*/
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {
		case "l":
			if !m.locked && !m.keycastMode {
				m.showLayoutList = !m.showLayoutList
				m.showSizeList = false
				m.showStandardList = false
				m.showModeList = false
				m.showHelpList = false
			}
			return m, nil
		case "s":
			if !m.locked && !m.keycastMode {
				m.showSizeList = !m.showSizeList
				m.showLayoutList = false
				m.showStandardList = false
				m.showModeList = false
				m.showHelpList = false
			}
			return m, nil
		case "d":
			if !m.locked && !m.keycastMode {
				m.showStandardList = !m.showStandardList
				m.showLayoutList = false
				m.showSizeList = false
				m.showModeList = false
				m.showHelpList = false
			}
			return m, nil
		case "h":
			if !m.locked {
				m.showAllInfo = !m.showAllInfo
				_ = config.Save(m.saveConfig())
			}
			return m, nil
		case "m":
			m.showModeList = !m.showModeList
			if m.keycastMode {
				m.modeList.Selected = 1
			} else {
				m.modeList.Selected = 0
			}
			m.showLayoutList = false
			m.showSizeList = false
			m.showStandardList = false
			m.showQuitDialog = false
			m.showHelpList = false
			return m, cmd
		case "?":
			m.showHelpList = !m.showHelpList
			m.showLayoutList = false
			m.showSizeList = false
			m.showStandardList = false
			m.showModeList = false
			m.showQuitDialog = false
			return m, cmd
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
		case m.showModeList:
			return m.handleModeListUpdate(msg)
		case m.showHelpList:
			return m.handleHelpListUpdate(msg)
		default:
			return m.handleGlobalKeys(msg)
		}
	case input.KeyEvent:
		var down bool
		switch msg.State {
		case input.KeyStateReleased:
			down = false
		case input.KeyStatePressed, input.KeyStateRepeated:
			down = true
		default:
			return m, nil
		}

		m.pressedKeys[msg.Code] = down
		if m.keycastMode && down && !isKeycastModifier(msg.Code) {
			if label, ok := keyboard.ResolveKeycastLabel(msg.Code, m.activeLayout, m.activeStandard, m.pressedKeys, m.capsLock); ok {
				m.keycastFadeVer++
				fng := basepkg.EvCodeFinger[msg.Code]
				entry := keycastEntry{label: label, version: m.keycastFadeVer, finger: fng, pressedAt: time.Now()}
				m.keycastKeys = append(m.keycastKeys, entry)
				if len(m.keycastKeys) > 5 {
					m.keycastKeys = m.keycastKeys[1:]
				}
				ver := entry.version
				cmd = func() tea.Msg {
					time.Sleep(1500 * time.Millisecond)
					return keycastFadeMsg{version: ver}
				}
			}
		}
		if msg.Code == basepkg.KEY_CAPSLOCK && msg.State == input.KeyStatePressed {
			m.capsLock = !m.capsLock
		}
		if msg.Code == basepkg.KEY_KATAKANAHIRAGANA {
			m.kanaKeyHeld = down
		}
		if msg.Code == basepkg.KEY_HANGEUL {
			m.hangeulKeyHeld = down
		}
	case keycastFadeMsg:
		now := time.Now()
		var kept []keycastEntry
		for _, e := range m.keycastKeys {
			if now.Sub(e.pressedAt) < 1500*time.Millisecond {
				kept = append(kept, e)
			}
		}
		m.keycastKeys = kept
		return m, nil
	case tea.WindowSizeMsg:
		m.terminalWidth = msg.Width
		m.terminalHeight = msg.Height
	}

	m.pressedKeys[basepkg.KEY_CAPSLOCK] = m.pressedKeys[basepkg.KEY_CAPSLOCK] || m.capsLock
	m.pressedKeys[basepkg.KEY_KATAKANAHIRAGANA] = m.kanaKeyHeld || m.kanaActive
	m.pressedKeys[basepkg.KEY_HANGEUL] = m.hangeulKeyHeld || m.hangeulActive

	return m, cmd
}

/*
handleLayoutListUpdate applies navigation or a selected logical layout while
keeping layout-specific standard changes and persistence together.
*/
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
		_ = config.Save(m.saveConfig())
		return m, nil
	case components.ListCancel:
		m.showLayoutList = false
		return m, nil
	}

	return m, nil
}

/*
handleSizeListUpdate applies navigation or a selected keyboard size and saves
confirmed changes after closing the picker.
*/
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
		_ = config.Save(m.saveConfig())
		return m, nil
	case components.ListCancel:
		m.showSizeList = false
		return m, nil
	}

	return m, nil
}

/*
handleStandardListUpdate applies navigation or a physical standard, clearing
overlays that are only meaningful for the previously selected standard.
*/
func (m Model) handleStandardListUpdate(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	var action components.ListAction
	m.standardList, action = m.standardList.Update(msg)

	switch action {

	case components.ListConfirm:
		m.activeStandard = m.standardList.Items[m.standardList.Selected]
		m.showStandardList = false
		m.kanaActive = false
		m.hangeulActive = false
		_ = config.Save(m.saveConfig())
		return m, nil
	case components.ListCancel:
		m.showStandardList = false
		return m, nil
	}

	return m, nil
}

/*
handleQuitDialogUpdate translates dialog confirmation into tea.Quit while a
cancellation returns control to the keyboard view.
*/
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

/*
handleModeListUpdate switches between full-keyboard and keycast modes and
clears history that should not carry across mode changes.
*/
func (m Model) handleModeListUpdate(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	var action components.ListAction
	m.modeList, action = m.modeList.Update(msg)

	switch action {

	case components.ListConfirm:
		m.keycastMode = m.modeList.Selected == 1
		m.showModeList = false
		m.keycastKeys = nil
	case components.ListCancel:
		m.showModeList = false
	}

	return m, nil
}

/*
handleHelpListUpdate delegates help navigation and closes the overlay when the
list reports cancellation.
*/
func (m Model) handleHelpListUpdate(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	var action components.ListAction
	m.helpList, action = m.helpList.Update(msg)

	switch action {
	case components.ListCancel:
		m.showHelpList = false
		return m, nil
	}

	return m, nil
}

/*
handleGlobalKeys handles shortcuts that apply when no overlay owns keyboard
input, including quitting and standard-specific display toggles.
*/
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
	case "f":
		if m.keycastMode {
			m.keycastFingerColors = !m.keycastFingerColors
		}
	case "v":
		switch m.activeVisual {
		case config.VisualASCII:
			m.activeVisual = config.VisualBoxDraw
		case config.VisualBoxDraw:
			m.activeVisual = config.VisualASCII
		}
		_ = config.Save(m.saveConfig())
	}

	return m, nil
}

/*
isKeycastModifier identifies keys that affect another key's label but should
not create standalone entries in keycast history.
*/
func isKeycastModifier(code uint16) bool {
	switch code {
	case basepkg.KEY_LEFTSHIFT, basepkg.KEY_RIGHTSHIFT,
		basepkg.KEY_LEFTCTRL, basepkg.KEY_RIGHTCTRL,
		basepkg.KEY_LEFTALT, basepkg.KEY_RIGHTALT,
		basepkg.KEY_LEFTMETA, basepkg.KEY_RIGHTMETA,
		basepkg.KEY_FN,
		basepkg.KEY_CAPSLOCK, basepkg.KEY_NUMLOCK, basepkg.KEY_SCROLLLOCK:
		return true
	}
	return false
}
