package main

import (
	// "net/http"
	"time"

	"github.com/andlabs/ui"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.Static("/", "./")

	uiErr := ui.Main(func() {
		startBtn := ui.NewButton("Start")
		stopBtn := ui.NewButton("Stop")
		greeting := ui.NewLabel("greet")
		box := ui.NewVerticalBox()
		window := ui.NewWindow("Quick Serv", 200, 100, false)
		window.SetChild(box)

		box.Append(greeting, false)
		box.Append(startBtn, false)
		box.Append(stopBtn, false)

		// Events
		startBtn.OnClicked(func(*ui.Button) {
			e.Logger.Fatal(e.Start(":8080"))
		})

		stopBtn.OnClicked(func(*ui.Button) {
			e.Shutdown(1 * time.Second)
		})

		window.OnClosing(func(*ui.Window) bool {
			// e.ShutdownTimeout = 1 * time.Second
			e.Shutdown(1 * time.Second)
			ui.Quit()
			return true
		})

		window.Show()
	})
	if uiErr != nil {
		panic(uiErr)
	}

}
