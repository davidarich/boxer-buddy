package handler

import (
	"github.com/davidarich/boxer-buddy/internal/adapters/log"
	"github.com/davidarich/boxer-buddy/internal/services/multibox"
)

const (
	GAME_CLIENT_FOCUS_MSG string = "focus"
	GAME_CLIENT_START_MSG string = "start"
	GAME_CLIENT_STOP_MSG  string = "stop"
)

// Handler for GameClient interop messages
type GameClient struct {
	engine *multibox.Engine
	logger log.Logger
}

func (g *GameClient) Handle(msg InteropMessage, response chan<- []byte) (err error) {
	switch msg.Cmd {
	case GAME_CLIENT_FOCUS_MSG:
		g.handleFocus(msg)
	case GAME_CLIENT_START_MSG:
		g.handleStart(msg)
	case GAME_CLIENT_STOP_MSG:
		g.handleStop(msg)
	}
	return
}

func (g *GameClient) handleFocus(msg InteropMessage) {
	if len(msg.Args) < 1 {
		g.logger.Error("engine failed to focus game client: missing argument [0] - no profile name specified")
		return
	}
	g.engine.FocusGame(msg.Args[0])
}

func (g *GameClient) handleStart(msg InteropMessage) {
	if len(msg.Args) < 1 {
		g.logger.Error("engine failed to start game client: missing argument [0] - no profile name specified")
		return
	}
	g.engine.StartGame(msg.Args[0])
}

func (g *GameClient) handleStop(msg InteropMessage) {
	if len(msg.Args) < 1 {
		g.logger.Error("engine failed to stop game client: missing argument [0] - no profile name specified")
		return
	}
	g.engine.StopGame(msg.Args[0])
}

func NewGameClient(logger log.Logger, engine *multibox.Engine) *GameClient {
	return &GameClient{
		logger: logger,
		engine: engine,
	}
}
