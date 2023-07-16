package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var Window fyne.Window

func main() {
	myApp := app.New()
	Window = myApp.NewWindow("My App")

	Window.SetContent(GetTabs())
	Window.ShowAndRun()
}
