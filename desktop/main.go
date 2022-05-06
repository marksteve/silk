package main

import (
	"context"
	"embed"
	"fmt"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed frontend/dist
var assets embed.FS

func main() {

	app := &App{}

	err := wails.Run(&options.App{
		Title:      "Basic Demo",
		Width:      1024,
		Height:     768,
		Assets:     &assets,
		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}

type App struct {
	ctx context.Context
}

func (b *App) startup(ctx context.Context) {
	b.ctx = ctx
}

func (b *App) shutdown(ctx context.Context) {}

func (b *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
