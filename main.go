package main

import (
	"fmt"
	"runtime"

	"github.com/andlabs/ui"
	"github.com/skratchdot/open-golang/open"
	"github.com/walkerEpps/quick_serve/server"
)

func main() {
	command := "milk"
	if runtime.GOOS == "windows" {
		command += ".bat"
	}
	fmt.Println(command)
	// init := false

	server := server.New()

	uiErr := ui.Main(func() {
		startBtn := ui.NewButton("Start")
		stopBtn := ui.NewButton("Stop")
		openBrowserBtn := ui.NewButton("Open Browser")
		status := ui.NewLabel("Server is Stoped")
		box := ui.NewVerticalBox()
		window := ui.NewWindow("Quick Serv", 200, 100, false)
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
	})
	if uiErr != nil {
		panic(uiErr)
	}
}
