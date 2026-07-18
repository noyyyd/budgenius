package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

const appName = "budgenius"

var (
	dataDir string // ldflags

	//go:embed all:frontend/dist
	assets embed.FS
)

func main() {
	var err error
	if dataDir == "" {
		dataDir, err = appDataDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	}

	app := NewApp()

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
		OnShutdown:       app.shutdown,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

func appDataDir() (string, error) {
	switch runtime.GOOS {
	case "linux":
		base := os.Getenv("XDG_DATA_HOME")
		if base == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				return "", err
			}
			base = filepath.Join(home, ".local", "share")
		}
		return filepath.Join(base, appName), nil
	case "darwin":
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(
			home,
			"Library",
			"Application Support",
			appName,
		), nil
	case "windows":
		return filepath.Join(
			os.Getenv("LOCALAPPDATA"),
			appName,
		), nil
	}

	return "", fmt.Errorf("unsupported OS")
}
