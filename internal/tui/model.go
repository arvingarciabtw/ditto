/*
Package tui implements the Bubble Tea model, update, view loop
along with the TUI styling and overlay components.
*/
package tui

import (
	"time"

	"github.com/arvingarciabtw/ditto/internal/config"
	"github.com/arvingarciabtw/ditto/internal/keyboard"
	"github.com/arvingarciabtw/ditto/internal/tui/components"
)

/*
keycastEntry records a rendered key press and its display metadata so keycast
mode can color, order, and expire entries independently of held-key state.
*/
type keycastEntry struct {
	label     string
	version   int
	finger    keyboard.Finger
	pressedAt time.Time
}

/*
Model contains all persisted choices, transient input state, and overlay state
owned by the Bubble Tea update loop.
*/
type Model struct {
	activeLayout   string
	activeSize     int
	activeStandard string
	activeVisual   string
	locked         bool

	layoutList   components.ListModel
	sizeList     components.ListModel
	standardList components.ListModel
	modeList     components.ListModel
	helpList     components.ListModel
	quitDialog   components.DialogModel

	showLayoutList   bool
	showSizeList     bool
	showStandardList bool
	showModeList     bool
	showHelpList     bool
	showQuitDialog   bool
	showAllInfo      bool

	pressedKeys    map[uint16]bool
	capsLock       bool
	kanaKeyHeld    bool
	kanaActive     bool
	hangeulKeyHeld bool
	hangeulActive  bool

	terminalWidth  int
	terminalHeight int

	keycastMode         bool
	keycastKeys         []keycastEntry
	keycastFadeVer      int
	keycastFingerColors bool
}

/*
InitModel constructs the initial TUI state from validated configuration and
synchronizes each picker with its active value.
*/
func InitModel(cfg config.Config) Model {
	layoutList := components.ListModel{
		Items:        keyboard.LayoutListItems,
		Selected:     0,
		Title:        "Layouts",
		AccentColor:  LayoutColor,
		VisibleCount: 3,
	}
	for i, item := range layoutList.Items {
		if item == cfg.ActiveLayout {
			layoutList.Selected = i
			break
		}
	}

	sizeList := components.ListModel{
		Items:        keyboard.LayoutSizeItems,
		Selected:     0,
		Title:        "Sizes",
		AccentColor:  SizeColor,
		VisibleCount: 3,
	}
	for i, s := range keyboard.Sizes {
		if s == cfg.ActiveSize {
			sizeList.Selected = i
			break
		}
	}

	standardList := components.ListModel{
		Items:        keyboard.StandardListItems,
		Selected:     0,
		Title:        "Standards",
		AccentColor:  StandardColor,
		VisibleCount: 3,
	}
	for i, item := range standardList.Items {
		if item == cfg.ActiveStandard {
			standardList.Selected = i
			break
		}
	}

	showAllInfo := true
	if cfg.ShowAllInfo != nil {
		showAllInfo = *cfg.ShowAllInfo
	}

	helpItems := dedupBindings(components.Commands)

	return Model{
		layoutList:   layoutList,
		sizeList:     sizeList,
		standardList: standardList,
		modeList: components.ListModel{
			Items:        []string{"Default", "Keycast"},
			Selected:     0,
			Title:        "Mode",
			AccentColor:  ModeColor,
			VisibleCount: 2,
		},
		helpList: components.ListModel{
			Items:        helpItems,
			Selected:     0,
			Title:        "Key Bindings",
			AccentColor:  HelpColor,
			VisibleCount: 0,
			HideEnter:    true,
			FooterLeft:   true,
		},
		quitDialog: components.DialogModel{
			AccentColor: QuitColor,
			Prompt:      "Are you sure you want to quit?",
			LeftLabel:   "Quit",
			RightLabel:  "Cancel",
		},

		activeLayout:   cfg.ActiveLayout,
		activeSize:     cfg.ActiveSize,
		activeStandard: cfg.ActiveStandard,
		activeVisual:   cfg.ActiveVisual,
		locked:         cfg.Locked,

		showLayoutList:   false,
		showSizeList:     false,
		showStandardList: false,
		showAllInfo:      showAllInfo,

		pressedKeys: make(map[uint16]bool),
	}
}

/*
saveConfig extracts only persistent user choices from the model so transient
terminal, input, and overlay state never reaches the config file.
*/
func (m Model) saveConfig() config.Config {
	v := m.showAllInfo

	return config.Config{
		ActiveLayout:   m.activeLayout,
		ActiveSize:     m.activeSize,
		ActiveStandard: m.activeStandard,
		ActiveVisual:   m.activeVisual,
		Locked:         m.locked,
		ShowAllInfo:    &v,
	}
}

/*
dedupBindings builds ordered help rows while collapsing bindings that share a
key, such as the standard-specific character overlay command.
*/
func dedupBindings(b components.Bindings) []string {
	/*
		binding pairs a displayed key with its description while preserving the
		deliberate help ordering before duplicate keys are removed.
	*/
	type binding struct {
		key  string
		desc string
	}

	order := []binding{
		{b.Size.Help().Key, b.Size.Help().Desc},
		{b.Layout.Help().Key, b.Layout.Help().Desc},
		{b.Standard.Help().Key, b.Standard.Help().Desc},
		{b.Visual.Help().Key, b.Visual.Help().Desc},
		{b.Keycast.Help().Key, b.Keycast.Help().Desc},
		{b.Kana.Help().Key, b.Kana.Help().Desc},
		{b.Finger.Help().Key, b.Finger.Help().Desc},
		{b.HideKey.Help().Key, b.HideKey.Help().Desc},
		{b.Quit.Help().Key, b.Quit.Help().Desc},
		{b.KeyBindings.Help().Key, b.KeyBindings.Help().Desc},
	}

	seen := make(map[string]bool)
	var items []string
	for _, b := range order {
		if seen[b.key] {
			continue
		}
		seen[b.key] = true
		items = append(items, b.key+"  "+b.desc)
	}

	return items
}
