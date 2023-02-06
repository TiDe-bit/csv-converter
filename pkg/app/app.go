package app

import (
	"changeme/pkg/converter"
	"changeme/pkg/options"
	"context"
)

// App struct
type App struct {
	ctx     context.Context
	options *converter.ConvertOptions
}

// Create creates a new App application struct
func Create() *App {
	return &App{}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.options = a.LoadOptions()
}

func (a *App) LoadOptions() *converter.ConvertOptions {
	return options.Load()
}

func (a *App) SpreadColumns() []string {
	return converter.SpreadColumns()
}

func (a *App) Convert() {
	options.Save(a.options)
	converter.Convert(a.options)
}

func (a *App) AddOption(from, to int, enableKDIfNew bool) {
	pair := converter.FromToPair{
		From: from,
		To:   to,
	}
	if a.options == nil {
		a.options = converter.WithOptions(enableKDIfNew)
	}
	a.options.AddOption(pair)
}

func (a *App) RemoveOption(from, to int) {
	a.options.RemoveOption(converter.FromToPair{From: from, To: to})
}
