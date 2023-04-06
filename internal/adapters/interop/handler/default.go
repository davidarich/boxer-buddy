package handler

import "github.com/davidarich/boxer-buddy/internal/adapters/log"

// Interop event handler used to log unhandled events
type DefaultLogger struct {
	logger log.Logger
}

func (d *DefaultLogger) Handle(msg InteropMessage, response chan<- []byte) (err error) {
	d.logger.Warn(msg)
	d.logger.Warn("event type was not routed to a specific handler")
	return
}

func NewDefault(
	logger log.Logger,
) *DefaultLogger {
	return &DefaultLogger{
		logger: logger,
	}
}
