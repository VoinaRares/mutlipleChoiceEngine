package main

import (
	"choice_engine/backend/traversal"
	"context"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Initialize(treePath string, playerPath string) error {
	return traversal.Initialize(treePath, playerPath)
}

func (a *App) GetCurrent() traversal.Option {
	return traversal.GetCurrent()
}

func (a *App) ChooseOption(optionId string) error {
	return traversal.ChooseOption(optionId)
}

func (a *App) GetOptions(optionId string) []traversal.Option {
	return traversal.GetOptions(optionId)
}
