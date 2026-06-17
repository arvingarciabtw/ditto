package main

import (
	"strconv"
	"strings"

	list "charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
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

		case "ctrl+shift+h":
			m.showInfoBar = !m.showInfoBar
			return m, nil

		case "?":
			m.helpModel.ShowAll = !m.helpModel.ShowAll
			return m, nil

		case "enter":
			if m.showLayoutList && m.layoutList.FilterState() != list.Filtering {
				if selected, ok := m.layoutList.SelectedItem().(item); ok {
					m.activeLayout = strings.ToLower(selected.title)
				}
				m.showLayoutList = false
				return m, nil
			}
			if m.showSizeList && m.sizeList.FilterState() != list.Filtering {
				if selected, ok := m.sizeList.SelectedItem().(item); ok {
					sizeStr := strings.TrimSuffix(selected.title, "%")
					if size, err := strconv.Atoi(sizeStr); err == nil {
						m.activeSize = size
					}
				}
				m.showSizeList = false
				return m, nil
			}

		case "q":
			isFilteringLayout := m.showLayoutList && m.layoutList.FilterState() == list.Filtering
			isFilteringSize := m.showSizeList && m.sizeList.FilterState() == list.Filtering
			hasOpenMenu := m.showLayoutList || m.showSizeList

			if isFilteringLayout || isFilteringSize {
				break
			}
			if hasOpenMenu {
				m.showLayoutList = false
				m.showSizeList = false
				return m, nil
			}
			return m, tea.Quit

		case "ctrl+c":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		h, v := listFrameStyle.GetFrameSize()
		m.layoutList.SetSize(msg.Width-h, msg.Height-v)
		m.sizeList.SetSize(msg.Width-h, msg.Height-v)
		m.helpModel.SetWidth(msg.Width)
	}

	var cmd1, cmd2 tea.Cmd
	m.layoutList, cmd1 = m.layoutList.Update(msg)
	m.sizeList, cmd2 = m.sizeList.Update(msg)
	return m, tea.Batch(cmd1, cmd2)
}
