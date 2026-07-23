package components

import (
	bkey "charm.land/bubbles/v2/key"
)

type Bindings struct {
	Layout   bkey.Binding
	Size     bkey.Binding
	Standard bkey.Binding
	HideKey  bkey.Binding
	Kana     bkey.Binding
	Hangeul  bkey.Binding
	Keycast  bkey.Binding
	Finger   bkey.Binding
}

var Commands = Bindings{
	Layout: bkey.NewBinding(
		bkey.WithKeys("l"),
		bkey.WithHelp("l", "layout"),
	),
	Size: bkey.NewBinding(
		bkey.WithKeys("s"),
		bkey.WithHelp("s", "size"),
	),
	Standard: bkey.NewBinding(
		bkey.WithKeys("d"),
		bkey.WithHelp("d", "standard"),
	),
	HideKey: bkey.NewBinding(
		bkey.WithKeys("h"),
		bkey.WithHelp("h", "hide"),
	),
	Kana: bkey.NewBinding(
		bkey.WithKeys("c"),
		bkey.WithHelp("c", "chars"),
	),
	Hangeul: bkey.NewBinding(
		bkey.WithKeys("c"),
		bkey.WithHelp("c", "chars"),
	),
	Keycast: bkey.NewBinding(
		bkey.WithKeys("m"),
		bkey.WithHelp("m", "mode"),
	),
	Finger: bkey.NewBinding(
		bkey.WithKeys("f"),
		bkey.WithHelp("f", "finger"),
	),
}
