package handler

import (
	"github.com/davidarich/boxer-buddy/internal/adapters/log"
	"github.com/davidarich/boxer-buddy/internal/ports"
	"github.com/davidarich/boxer-buddy/internal/services/multibox"
)

const (
	ENGINE_START_MSG  string = "start"
	ENGINE_STOP_MSG   string = "stop"
	ENGINE_STATUS_MSG string = "status"
)

// Handler for Engine interop messages
type Engine struct {
	settings ports.Settings
	logger   log.Logger
	engine   *multibox.Engine
}

// handles a message and channel for responding
func (e *Engine) Handle(msg InteropMessage, response chan<- []byte) (err error) {
	switch msg.Cmd {
	case ENGINE_START_MSG:
		e.handleStart(msg)
	case ENGINE_STOP_MSG:
		e.handleStop(msg)
	case ENGINE_STATUS_MSG:
		e.handleStatus(msg)
	}
	return
}

func (e *Engine) handleStatus(msg InteropMessage) {
	e.engine.Status()
}

func (e *Engine) handleStart(msg InteropMessage) {
	// validate that a profile was specified
	if len(msg.Args) < 1 {
		e.logger.Error("engine failed to start: missing argument [0] - no profile specified")
		return
	}

	// load settings
	cfg, err := e.settings.Get()
	if err != nil {
		e.logger.Error("engine failed to start: error loading settings")
		return
	}

	// get the specified profile
	mb := cfg.GetGroupByName(msg.Args[0])

	// confirm that an entry was found
	if mb == nil {
		e.logger.Error("engine failed to start: profile not found")
		return
	}
	e.logger.Info("found specified profile: ", mb.Name)

	e.engine.Start(mb)
}

func (e *Engine) handleStop(msg InteropMessage) {
	e.engine.Stop()
}

func NewEngine(
	settings ports.Settings,
	logger log.Logger,
	multiboxEngine *multibox.Engine,
) *Engine {
	return &Engine{
		settings: settings,
		logger:   logger,
		engine:   multiboxEngine,
	}
}
