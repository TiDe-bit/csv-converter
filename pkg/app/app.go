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

type MyFromToPair struct {
	From int
	To   int
}

func (m MyFromToPair) toConverterFromToPair() converter.FromToPair {
	pair := converter.FromToPair{}
	pair.From = m.From
	pair.To = m.To
	return pair
}

// Create creates a new App application struct
func Create() *App {
	return &App{}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.options = options.Load()
}

func (a *App) GetSavedOptions() []*converter.FromToPair {
	a.options = options.Load()
	array := make([]*converter.FromToPair, len(a.options.FromToPairs))
	logrus.Infof("saved options %+v", array)
	return array
}

func (a *App) SpreadColumns() []string {
	return converter.SpreadColumns()
}

func (a *App) Convert() {
	converter.Convert(a.options)
}

func (a *App) SaveOptions(options []MyFromToPair) {
	logrus.Infof("save this pls %+v", options)
	for _, option := range options {
		a.options.AddOption(option.toConverterFromToPair())
	}
}
