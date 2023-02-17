package app

import (
	"changeme/pkg/converter"
	"changeme/pkg/options"
	"context"

	"github.com/sirupsen/logrus"
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

func (a *App) GetSavedOptions() []*converter.FromToPair {
	array := make([]*converter.FromToPair, len(a.options.FromToPairs))
	return array
}

func (a *App) LoadOptions() *converter.ConvertOptions {
	return options.Load()
}

func (a *App) SpreadColumns() []string {
	return converter.SpreadColumns()
}

func (a *App) Convert() {
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
	go options.Save(a.options)
	go logrus.Infof("got %+v", a.options)
}

func (a *App) RemoveOption(from, to int) {
	a.options.RemoveOption(converter.FromToPair{From: from, To: to})
}
