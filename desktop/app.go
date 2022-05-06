package main

import (
	"context"

	"github.com/marksteve/silk/store"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	s, err := store.NewStore("silk")
	if err != nil {
		panic(err)
	}
	a.ctx = context.WithValue(ctx, "store", s)
	runtime.LogInfo(ctx, "Store loaded")
	runtime.EventsEmit(ctx, "startup")
}

func (a *App) GetFibers() ([]store.Fiber, error) {
	s := a.ctx.Value("store").(*store.Store)
	return s.GetFibers("silk")
}

func (a *App) GetDbOptions() interface{} {
	s := a.ctx.Value("store").(*store.Store)
	return s.GetDbOptions("silk")
}
