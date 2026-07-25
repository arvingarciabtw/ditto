package components

import (
	bkey "charm.land/bubbles/v2/key"
)

type Bindings struct {
	Layout      bkey.Binding
	Size        bkey.Binding
	Standard    bkey.Binding
	HideKey     bkey.Binding
	Kana        bkey.Binding
	Hangeul     bkey.Binding
	Keycast     bkey.Binding
	Finger      bkey.Binding
	Visual      bkey.Binding
	Quit        bkey.Binding
	KeyBindings bkey.Binding
}

var Commands = Bindings{
	Size: bkey.NewBinding(
		bkey.WithKeys("s"),
		bkey.WithHelp("s", "size"),
	),
	Layout: bkey.NewBinding(
		bkey.WithKeys("l"),
		bkey.WithHelp("l", "layout"),
	),
	Standard: bkey.NewBinding(
		bkey.WithKeys("d"),
		bkey.WithHelp("d", "standard"),
	),
	Visual: bkey.NewBinding(
		bkey.WithKeys("v"),
		bkey.WithHelp("v", "visual"),
	),
	Keycast: bkey.NewBinding(
		bkey.WithKeys("m"),
		bkey.WithHelp("m", "mode"),
	),
	Kana: bkey.NewBinding(
		bkey.WithKeys("c"),
		bkey.WithHelp("c", "chars"),
	),
	Hangeul: bkey.NewBinding(
		bkey.WithKeys("c"),
		bkey.WithHelp("c", "chars"),
	),
	Finger: bkey.NewBinding(
		bkey.WithKeys("f"),
		bkey.WithHelp("f", "finger"),
	),
	HideKey: bkey.NewBinding(
		bkey.WithKeys("h"),
		bkey.WithHelp("h", "hide"),
	),
	Quit: bkey.NewBinding(
		bkey.WithKeys("q"),
		bkey.WithHelp("q", "quit"),
	),
	KeyBindings: bkey.NewBinding(
		bkey.WithKeys("?"),
		bkey.WithHelp("?", "key bindings"),
	),
}
