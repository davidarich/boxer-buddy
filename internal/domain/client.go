package domain

// represents the state of a game client program
type GameClient struct {
	Process      *Process
	WindowBounds *Window
}

// returns bool to indicate if a client is running
func (gc *GameClient) IsRunning() bool {
	if gc.Process != nil {
		return gc.Process.Id > 0
	}
	return false
}

func (gc *GameClient) ClearProcess() {
	gc.Process = nil
}

func (gc *GameClient) ClearWindow() {
	gc.WindowBounds = nil
}

func NewGameClient() *GameClient {
	return &GameClient{}
}
