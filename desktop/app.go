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

	s, err := store.NewStore("silk-RJoboAy9")
	if err != nil {
		runtime.LogFatalf(ctx, "Failed to open: %s", err)
	}
	a.ctx = context.WithValue(ctx, "store", s)

	runtime.EventsEmit(ctx, "startup")
	runtime.LogInfo(ctx, "Store initialized")
}

func (a *App) shutdown(ctx context.Context) {
	s := a.ctx.Value("store").(*store.Store)
	s.Close()
}

func (a *App) GetFibers() ([]store.Fiber, error) {
	s := a.ctx.Value("store").(*store.Store)
	return s.GetFibers()
}

func (a *App) Weave(data []byte) error {
	s := a.ctx.Value("store").(*store.Store)
	return s.Weave(data)
}

func (a *App) GetDbOptions() interface{} {
	s := a.ctx.Value("store").(*store.Store)
	return s.GetDbOptions()
}
