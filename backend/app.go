package app

import (
	"context"
)

type App struct {
	Ctx context.Context
}

func Init() *App {
	return &App{}
}

func (a *App) Startup(ctx context.Context) {
	a.Ctx = ctx
}
