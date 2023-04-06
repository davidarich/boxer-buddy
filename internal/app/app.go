package app

import (
	"fmt"
	"os"

	"github.com/davidarich/boxer-buddy/internal/adapters/interop"
	"github.com/davidarich/boxer-buddy/internal/adapters/interop/handler"
	"github.com/davidarich/boxer-buddy/internal/adapters/log"
	"github.com/davidarich/boxer-buddy/internal/adapters/os/process"
	"github.com/davidarich/boxer-buddy/internal/adapters/os/user"
	"github.com/davidarich/boxer-buddy/internal/adapters/settings"
	"github.com/davidarich/boxer-buddy/internal/adapters/view"
	"github.com/davidarich/boxer-buddy/internal/services/multibox"
)

type App struct {
	Name    string
	Version string
}

func (a *App) Run() {
	cfg := settings.NewPersistentAdapter()
	err := cfg.Load()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	l := log.NewZapLogFactory()
	p := process.NewWindowsAdapter(l.Logger())
	w := user.NewWindowsAdapter()
	m := multibox.NewEngine(l.Logger(), p, w)
	v := view.NewWindowManager(cfg)

	// interop
	dH := handler.NewDefault(l.Logger())
	eH := handler.NewEngine(cfg, l.Logger(), m)
	cH := handler.NewGameClient(l.Logger(), m)
	sH := handler.NewSettings(l.Logger(), v)
	r := interop.NewEventRouter(l.Logger(), dH, eH, cH, sH)
	s := interop.NewServer(l.Logger(), r, cfg)
	m.SetActiveProfileDispatcher(s)

	// start server for UI via event based interop
	go s.StartIO()

	// start UI
	v.OpenMain(fmt.Sprintf("%s - %s", a.Name, a.Version))
}

func NewApp() *App {
	return &App{
		Name:    "Boxer Buddy",
		Version: "v0.2.0 (alpha)",
	}
}
