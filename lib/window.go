package lib

import (
	"github.com/andlabs/ui"
	"github.com/skratchdot/open-golang/open"
)

// StartQuickServ starts the App
func StartQuickServ() error {
	return ui.Main(mainWindow)
}

func mainWindow() {
	server := NewServer()
	window := ui.NewWindow("Quick Serv", 200, 100, false)
	startBtn := ui.NewButton("Start")
	stopBtn := ui.NewButton("Stop")
	openBrowserBtn := ui.NewButton("Open Browser")
	status := ui.NewLabel("Server is Stoped")
	box := ui.NewVerticalBox()

	// Attaching to box of the main window
	window.SetChild(box)
	box.Append(status, false)
	box.Append(startBtn, false)
	box.Append(stopBtn, false)
	box.Append(openBrowserBtn, false)

	// Events
	startBtn.OnClicked(func(*ui.Button) {
		server.Run()
		status.SetText("Server is Running")
	})

	stopBtn.OnClicked(func(*ui.Button) {
		status.SetText("Server is Stopped")
		server.Stop()
	})

	openBrowserBtn.OnClicked(func(*ui.Button) {
		open.Run("http://localhost:" + server.Port)
	})

	window.OnClosing(func(*ui.Window) bool {

		server.Stop()
		ui.Quit()
		return true
	})

	window.Show()
}
