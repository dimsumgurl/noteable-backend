package main

import (
	"fmt"

	"github.com/dimsumgurl/noteable-backend/pkg/app"
)

func main() {
	app, err := app.NewApp()
	if err != nil {
		fmt.Print(err)
	}
	err = app.Initialize()
	if err != nil {
		fmt.Print(err)
	}
	err = app.Run()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print("app finished running")

}
