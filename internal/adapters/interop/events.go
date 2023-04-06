package interop

import (
	"github.com/davidarich/boxer-buddy/internal/services/multibox"
)

/*
events outgoing to UI
*/
type ActiveProfileUiEvent struct {
	EventType      string
	Action         string // actions: replace
	ActiveProfiles []*multibox.ActiveProfile
}

func NewActiveProfileUiEvent() *ActiveProfileUiEvent {
	return &ActiveProfileUiEvent{}
}
