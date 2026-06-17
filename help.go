package main

import (
	key "charm.land/bubbles/v2/key"
)

type keyMap struct {
	Layout  key.Binding
	Size    key.Binding
	HideKey key.Binding
	Help    key.Binding
	Quit    key.Binding
}

func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Layout},
		{k.Size},
		{k.HideKey},
		{k.Quit},
	}
}

var keys = keyMap{
	Layout: key.NewBinding(
		key.WithKeys("ctrl+shift+l"),
		key.WithHelp("ctrl+shift+l", "layout"),
	),
	Size: key.NewBinding(
		key.WithKeys("ctrl+shift+s"),
		key.WithHelp("ctrl+shift+s", "size"),
	),
	HideKey: key.NewBinding(
		key.WithKeys("ctrl+shift+h"),
		key.WithHelp("ctrl+shift+h", "hide Bar"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
