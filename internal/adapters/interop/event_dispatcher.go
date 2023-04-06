package interop

import (
	"encoding/json"

	"github.com/davidarich/boxer-buddy/internal/adapters/log"
	"github.com/davidarich/boxer-buddy/internal/services/multibox"
	"github.com/davidarich/boxer-buddy/internal/ports"
)

type ActiveProfileUiEventDispatcher struct {
	interopWriteAdapter ports.InteropWriter
	logger              log.Logger
}

func (d *ActiveProfileUiEventDispatcher) Replace(activeProfiles []*multibox.ActiveProfile) {
	e := NewActiveProfileUiEvent()
	e.Action = "replace"
	e.ActiveProfiles = activeProfiles
	d.dispatch(e)
}

func (d *ActiveProfileUiEventDispatcher) dispatch(e *ActiveProfileUiEvent) {
	message, err := json.Marshal(e)
	if err != nil {
		d.logger.Error("error preparing activeProfile state message")
		return
	}

	if d.interopWriteAdapter != nil {
		d.interopWriteAdapter.Write(message)
		return
	}
	d.logger.Error("no interop dispatcher is set")
}

func NewActiveProfileUiEventDispatcher(
	logger log.Logger,
	interopWriteAdapter ports.InteropWriter,
) *ActiveProfileUiEventDispatcher {
	return &ActiveProfileUiEventDispatcher{
		interopWriteAdapter: interopWriteAdapter,
		logger:              logger,
	}
}
