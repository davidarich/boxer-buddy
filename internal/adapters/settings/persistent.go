package settings

import (
	"encoding/json"
	"os"

	"github.com/davidarich/boxer-buddy/internal/domain"
)

// adapter which saves or loads settings to/from disk
type PersistentAdapter struct {
	cfg *domain.Settings
}

func (pa *PersistentAdapter) Get() (cfg *domain.Settings, err error) {
	cfg = pa.cfg
	return
}

// reads settings from disk
func (pa *PersistentAdapter) Load() (err error) {
	cfgData, err := os.ReadFile("settings.json")
	if err != nil {
		return
	}

	cfg := domain.NewSettings()
	err = json.Unmarshal(cfgData, cfg)
	if err != nil {
		return
	}
	pa.cfg = cfg
	return
}

// saves settings to disk
func (pa *PersistentAdapter) Save(cfg *domain.Settings) (err error) {
	cfgData, err := json.Marshal(cfg)
	if err != nil {
		return
	}

	err = os.WriteFile("settings.json", cfgData, 0)
	return
}

func NewPersistentAdapter() *PersistentAdapter {
	return &PersistentAdapter{}
}
