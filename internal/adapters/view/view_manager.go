package view

import "github.com/davidarich/boxer-buddy/internal/ports"

type ViewManager struct {
	main     *MainWindow
	settings ports.Settings
}

func (w *ViewManager) OpenMain(title string) {
	w.main = NewMainWindow(title)
	cfg, _ := w.settings.Get()
	w.main.Open(
		"http://" +
			cfg.UiOptions.GetAddress() +
			"?interop_host=" +
			cfg.InteropOptions.GetHost() +
			"&interop_port=" +
			cfg.InteropOptions.GetPort(),
	)
}

func NewWindowManager(settings ports.Settings) *ViewManager {
	return &ViewManager{
		settings: settings,
	}
}
