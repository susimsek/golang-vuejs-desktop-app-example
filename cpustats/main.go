package main

import (
	"cpustats/pkg/sys"
	_ "embed"
	"github.com/wailsapp/wails"
)

//go:embed frontend/dist/app.js
var js string

//go:embed frontend/dist/app.css
var css string

func main() {

	stats := &sys.Stats{}

	app := wails.CreateApp(&wails.AppConfig{
		Width:     512,
		Height:    512,
		Title:     "CPU Usage",
		JS:        js,
		CSS:       css,
		Resizable: true,
		Colour:    "#131313",
	})
	app.Bind(stats)
	app.Run()
}
