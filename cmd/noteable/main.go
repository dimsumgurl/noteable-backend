package main

import (
	"fmt"

	"github.com/dimsumgurl/noteable-backend/pkg/app"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		//do something
	}
	err = app.Run()
	if err != nil {
		// do something
	}
	fmt.Print("app finished running")
}
