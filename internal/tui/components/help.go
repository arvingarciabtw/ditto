package components

import (
	bkey "charm.land/bubbles/v2/key"
)

type Bindings struct {
	Layout  bkey.Binding
	Size    bkey.Binding
	HideKey bkey.Binding
}

var Commands = Bindings{
	Layout: bkey.NewBinding(
		bkey.WithKeys("ctrl+shift+l"),
		bkey.WithHelp("^l", "layout"),
	),
	Size: bkey.NewBinding(
		bkey.WithKeys("ctrl+shift+s"),
		bkey.WithHelp("^s", "size"),
	),
	HideKey: bkey.NewBinding(
		bkey.WithKeys("ctrl+shift+h"),
		bkey.WithHelp("^h", "hide"),
	),
}
