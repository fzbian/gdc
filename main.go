package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

var Window fyne.Window

func main() {
	App := app.New()
	Window = App.NewWindow("My App")

	Window.SetContent(GetTabs())
	Window.ShowAndRun()
}
