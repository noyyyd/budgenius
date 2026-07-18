package main

import (
	"context"
	"database/sql"
	"embed"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed schema.sql
var schema string

func main() {
	db, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		panic(err)
	}

	// Create an instance of the app structure
	app, err := NewApp(db, schema)
	if err != nil {
		println("Error:", err.Error())
	}

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "budgenius",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnShutdown: func(ctx context.Context) {
			db.Close()
		},
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
