package keyboard

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/arvingarciabtw/ditto/internal/config"
)

type customLayoutFile struct {
	Map   map[string]string `json:"map"`
	Shift map[string]string `json:"shift,omitempty"`
}

func init() {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		return
	}
	dir := filepath.Join(cfgDir, config.DirName, "layouts")

	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".json") {
			continue
		}
		name := strings.TrimSuffix(entry.Name(), ".json")
		if name == "config" {
			continue
		}

		data, err := os.ReadFile(filepath.Join(dir, entry.Name()))
		if err != nil {
			continue
		}

		var clf customLayoutFile
		if err := json.Unmarshal(data, &clf); err != nil {
			continue
		}

		if clf.Map == nil {
			continue
		}

		layouts[name] = clf.Map
		if clf.Shift != nil {
			shiftMaps[name] = clf.Shift
		} else {
			shiftMaps[name] = usShift
		}
	}
}
