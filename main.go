package main

import (
	"fmt"
	"runtime"

	"github.com/walker-walks/quick_serv/lib"
)

func main() {
	command := "milk"
	if runtime.GOOS == "windows" {
		command += ".bat"
	}
	fmt.Println(command)

	appErr := lib.StartQuickServ()
	if appErr != nil {
		panic(appErr)
	}
}
