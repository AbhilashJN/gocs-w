package main

import (
	_ "embed"

	"github.com/wailsapp/wails"
)

//go:embed frontend/build/static/js/main.js
var js string

//go:embed frontend/build/static/css/main.css
var css string

func main() {

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1024,
		Height:    768,
		Title:     "GOCS-W",
		JS:        js,
		CSS:       css,
		Colour:    "#131313",
		Resizable: true,
	})
	demo := NewDemo()
	app.Bind(demo)
	app.Run()
}
