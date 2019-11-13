package main

import (
	"log"

	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {
	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	myContainers, err := NewContainers()
	if err != nil {
		log.Fatal(err)
	}

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 1024,
		Title:  "logOut",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(myContainers)
	// app.Bind(ContainerInspect)
	// app.Bind(ContainerLogs)
	app.Run()
}
