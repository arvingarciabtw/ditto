package components

import (
	"testing"
)

func hasKey(keys []string, target string) bool {
	for _, k := range keys {
		if k == target {
			return true
		}
	}
	return false
}

func TestCommands_LayoutKeys(t *testing.T) {
	keys := Commands.Layout.Keys()
	if !hasKey(keys, "ctrl+shift+l") {
		t.Errorf("expected ctrl+shift+l in Layout keys, got %v", keys)
	}
}

func TestCommands_LayoutHelp(t *testing.T) {
	if Commands.Layout.Help().Key != "^l" {
		t.Errorf("expected ^l, got %q", Commands.Layout.Help().Key)
	}
	if Commands.Layout.Help().Desc != "layout" {
		t.Errorf("expected layout, got %q", Commands.Layout.Help().Desc)
	}
}

func TestCommands_SizeKeys(t *testing.T) {
	keys := Commands.Size.Keys()
	if !hasKey(keys, "ctrl+shift+s") {
		t.Errorf("expected ctrl+shift+s in Size keys, got %v", keys)
	}
}

func TestCommands_SizeHelp(t *testing.T) {
	if Commands.Size.Help().Key != "^s" {
		t.Errorf("expected ^s, got %q", Commands.Size.Help().Key)
	}
	if Commands.Size.Help().Desc != "size" {
		t.Errorf("expected size, got %q", Commands.Size.Help().Desc)
	}
}

func TestCommands_HideKeyKeys(t *testing.T) {
	keys := Commands.HideKey.Keys()
	if !hasKey(keys, "ctrl+shift+h") {
		t.Errorf("expected ctrl+shift+h in HideKey keys, got %v", keys)
	}
}

func TestCommands_HideKeyHelp(t *testing.T) {
	if Commands.HideKey.Help().Key != "^h" {
		t.Errorf("expected ^h, got %q", Commands.HideKey.Help().Key)
	}
	if Commands.HideKey.Help().Desc != "hide" {
		t.Errorf("expected hide, got %q", Commands.HideKey.Help().Desc)
	}
}

func TestCommands_threeBindings(t *testing.T) {
	// Just verify Commands has all three bindings
	if Commands.Layout.Keys() == nil {
		t.Error("Layout binding should have keys")
	}
	if Commands.Size.Keys() == nil {
		t.Error("Size binding should have keys")
	}
	if Commands.HideKey.Keys() == nil {
		t.Error("HideKey binding should have keys")
	}
}
