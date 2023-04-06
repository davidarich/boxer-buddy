package multibox

import "github.com/davidarich/boxer-buddy/internal/domain"

type ActiveProfile struct {
	IsEnabled   bool
	GameProfile domain.GameProfile
	GameClient  *domain.GameClient
}

func NewActiveProfile(gameProfile domain.GameProfile) *ActiveProfile {
	return &ActiveProfile{
		GameProfile: gameProfile,
		GameClient:  domain.NewGameClient(),
	}
}
