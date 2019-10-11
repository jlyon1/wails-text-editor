package main

import (
	"io/ioutil"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func basic() string {
	return "Hello World!"
}

// Save saves to file
func Save(value string) {
	ioutil.WriteFile("file.txt", []byte(value), 0644)
}

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "My Project",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(Save)
	app.Bind(basic)
	app.Run()
}
