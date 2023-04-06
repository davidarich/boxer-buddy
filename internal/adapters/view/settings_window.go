package view

import (
	"github.com/webview/webview"
)

type SettingsWindow struct {
	title   string
	webview webview.WebView
}

func (mw *SettingsWindow) Open(addr string) {
	debug := true
	mw.webview = webview.New(debug)
	defer mw.webview.Destroy()
	mw.webview.SetTitle(mw.title)
	mw.webview.SetSize(900, 700, webview.HintNone)
	mw.webview.Navigate(addr)
	mw.webview.Run()
}

func NewSettingsWindow() *SettingsWindow {
	return &SettingsWindow{
		title: "Settings",
	}
}
