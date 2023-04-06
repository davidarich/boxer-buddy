package interop

import (
	"encoding/json"

	"github.com/davidarich/boxer-buddy/internal/adapters/interop/handler"
	"github.com/davidarich/boxer-buddy/internal/adapters/log"
)

const (
	ENGINE_EVENT      string = "engine"
	GAME_CLIENT_EVENT string = "game_client"
	SETTINGS_EVENT    string = "settings"
)

type EventRouter struct {
	defaultHandler    *handler.DefaultLogger
	engineHandler     *handler.Engine
	gameClientHandler *handler.GameClient
	settingsHandler   *handler.Settings

	logger log.Logger
}

// match the event to its handler
func (er *EventRouter) Route(event []byte, response chan []byte) (err error) {
	ev, err := mapMessageEvent(event)
	// case/switch for handling each message "type"
	switch ev.Type {
	case ENGINE_EVENT:
		go er.engineHandler.Handle(*ev, response)
	case GAME_CLIENT_EVENT:
		go er.gameClientHandler.Handle(*ev, response)
	case SETTINGS_EVENT:
		go er.settingsHandler.Handle(*ev, response)
	default:
		go er.defaultHandler.Handle(*ev, response)
	}

	return
}

func NewEventRouter(
	logger log.Logger,
	defaultHandler *handler.DefaultLogger,
	engineEventHandler *handler.Engine,
	gameClientEventHandler *handler.GameClient,
	viewManagerEventHandler *handler.Settings,
) *EventRouter {
	return &EventRouter{
		logger:            logger,
		engineHandler:     engineEventHandler,
		gameClientHandler: gameClientEventHandler,
		settingsHandler:   viewManagerEventHandler,
	}
}

// decode message from bytes to struct
func mapMessageEvent(msg []byte) (*handler.InteropMessage, error) {
	ev := &handler.InteropMessage{}
	err := json.Unmarshal(msg, ev)

	return ev, err
}
