package handler

import (
	"github.com/davidarich/boxer-buddy/internal/adapters/log"
	"github.com/davidarich/boxer-buddy/internal/ports"
)

const (
	SETTINGS_SAVE_MSG  string = "save"
)

// Handler for Settings interop messages
type Settings struct {
	logger      log.Logger
	viewManager ports.ViewManager
}

func (e *Settings) Handle(msg InteropMessage, response chan<- []byte) (err error) {
	switch msg.Cmd {
	case SETTINGS_SAVE_MSG:
		e.handleSave(msg)
	}
	return
}

func (e *Settings) handleSave(msg InteropMessage) {
	// todo: implement settings persistance
}

func NewSettings(
	logger log.Logger,
	viewManager ports.ViewManager,
) *Settings {
	return &Settings{
		logger:      logger,
		viewManager: viewManager,
	}
}
