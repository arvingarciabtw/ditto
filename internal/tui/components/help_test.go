package components

import (
	"testing"

	bkey "charm.land/bubbles/v2/key"
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
	if !hasKey(keys, "l") {
		t.Errorf("expected l in Layout keys, got %v", keys)
	}
}

func TestCommands_LayoutHelp(t *testing.T) {
	if Commands.Layout.Help().Key != "l" {
		t.Errorf("expected ^l, got %q", Commands.Layout.Help().Key)
	}
	if Commands.Layout.Help().Desc != "layout" {
		t.Errorf("expected layout, got %q", Commands.Layout.Help().Desc)
	}
}

func TestCommands_SizeKeys(t *testing.T) {
	keys := Commands.Size.Keys()
	if !hasKey(keys, "s") {
		t.Errorf("expected s in Size keys, got %v", keys)
	}
}

func TestCommands_SizeHelp(t *testing.T) {
	if Commands.Size.Help().Key != "s" {
		t.Errorf("expected ^s, got %q", Commands.Size.Help().Key)
	}
	if Commands.Size.Help().Desc != "size" {
		t.Errorf("expected size, got %q", Commands.Size.Help().Desc)
	}
}

func TestCommands_HideKeyKeys(t *testing.T) {
	keys := Commands.HideKey.Keys()
	if !hasKey(keys, "h") {
		t.Errorf("expected h in HideKey keys, got %v", keys)
	}
}

func TestCommands_HideKeyHelp(t *testing.T) {
	if Commands.HideKey.Help().Key != "h" {
		t.Errorf("expected ^h, got %q", Commands.HideKey.Help().Key)
	}
	if Commands.HideKey.Help().Desc != "hide" {
		t.Errorf("expected hide, got %q", Commands.HideKey.Help().Desc)
	}
}

func TestCommands_StandardKeys(t *testing.T) {
	keys := Commands.Standard.Keys()
	if !hasKey(keys, "d") {
		t.Errorf("expected d in Standard keys, got %v", keys)
	}
}

func TestCommands_StandardHelp(t *testing.T) {
	if Commands.Standard.Help().Key != "d" {
		t.Errorf("expected d, got %q", Commands.Standard.Help().Key)
	}
	if Commands.Standard.Help().Desc != "std" {
		t.Errorf("expected std, got %q", Commands.Standard.Help().Desc)
	}
}

func TestCommands_KanaKeys(t *testing.T) {
	keys := Commands.Kana.Keys()
	if !hasKey(keys, "c") {
		t.Errorf("expected c in Kana keys, got %v", keys)
	}
}

func TestCommands_KanaHelp(t *testing.T) {
	if Commands.Kana.Help().Key != "c" {
		t.Errorf("expected c, got %q", Commands.Kana.Help().Key)
	}
	if Commands.Kana.Help().Desc != "chars" {
		t.Errorf("expected chars, got %q", Commands.Kana.Help().Desc)
	}
}

func TestCommands_HangeulKeys(t *testing.T) {
	keys := Commands.Hangeul.Keys()
	if !hasKey(keys, "c") {
		t.Errorf("expected c in Hangeul keys, got %v", keys)
	}
}

func TestCommands_HangeulHelp(t *testing.T) {
	if Commands.Hangeul.Help().Key != "c" {
		t.Errorf("expected c, got %q", Commands.Hangeul.Help().Key)
	}
	if Commands.Hangeul.Help().Desc != "chars" {
		t.Errorf("expected chars, got %q", Commands.Hangeul.Help().Desc)
	}
}

func TestCommands_allBindingsHaveKeys(t *testing.T) {
	bindings := []struct {
		name string
		b    bkey.Binding
	}{
		{"Layout", Commands.Layout},
		{"Size", Commands.Size},
		{"Standard", Commands.Standard},
		{"HideKey", Commands.HideKey},
		{"Kana", Commands.Kana},
		{"Hangeul", Commands.Hangeul},
	}
	for _, b := range bindings {
		if b.b.Keys() == nil {
			t.Errorf("%s binding should have keys", b.name)
		}
	}
}
