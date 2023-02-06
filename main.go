package main

import (
	"changeme/pkg/app"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the application structure
	application := app.Create()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "converter",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        application.Startup,
		Bind: []any{
			application,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
