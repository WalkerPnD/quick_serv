package main

import
// "net/http"
// "github.com/andlabs/ui"

"github.com/labstack/echo"

func main() {
	e := echo.New()
	e.Static("/", "./")
	e.Logger.Fatal(e.Start(":8080"))
}
