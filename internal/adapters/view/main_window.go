package view

import (
	"github.com/webview/webview"
)

type MainWindow struct {
	title   string
	webview webview.WebView
}

func (mw *MainWindow) Open(addr string) {
	debug := true
	mw.webview = webview.New(debug)
	defer mw.webview.Destroy()
	mw.webview.SetTitle(mw.title)
	mw.webview.SetSize(400, 800, webview.HintNone)
	mw.webview.Navigate(addr)
	mw.webview.Run()
}

func NewMainWindow(title string) *MainWindow {
	return &MainWindow{
		title: title,
	}
}
