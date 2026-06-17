package main

import (
	list "charm.land/bubbles/v2/list"
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

var keyboardLayoutItems = []list.Item{
	item{title: "QWERTY", desc: "The default"},
	item{title: "DVORAK", desc: "For nerds"},
	item{title: "COLEMAK", desc: "For mega nerds"},
	item{title: "COLEMAK-DH", desc: "Modern variant with better thumb positioning"},
	item{title: "WORKMAN", desc: "Reduces lateral finger movement"},
	item{title: "AZERTY", desc: "Standard French layout"},
}

var keyboardSizeItems = []list.Item{
	item{title: "60%", desc: "61 keys"},
	item{title: "65%", desc: "68 keys"},
	item{title: "75%", desc: "84 keys"},
	item{title: "80%", desc: "87 keys"},
	item{title: "96%", desc: "100 keys"},
	item{title: "100%", desc: "104 keys"},
}
