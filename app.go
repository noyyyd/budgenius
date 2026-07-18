package main

import (
	"budgenius/repository"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx        context.Context
	db         *sql.DB
	repository *repository.Repository
}

func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	f, err := os.OpenFile(filepath.Join(dataDir, "app.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		a.FatalErrorDialog(ctx, err)
	}
	log.SetOutput(f)

	log.Println("Start application")

	dbPath := filepath.Join(dataDir, "db.sqlite")

	log.Println("Connect to database:", dbPath)
	a.db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Println("Error connection to database:", err)
		a.FatalErrorDialog(ctx, err)
	}
	a.repository, err = repository.NewRepository(a.db)
	if err != nil {
		log.Println("Error creating repository:", err)
		a.FatalErrorDialog(ctx, err)
	}
}

func (a *App) FatalErrorDialog(ctx context.Context, err error) {
	runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:    runtime.ErrorDialog,
		Title:   "Ошибка",
		Message: err.Error(),
	})
	runtime.Quit(ctx)
}

func (a *App) shutdown(ctx context.Context) {
	log.Println("Stop application")
	a.db.Close()
	log.Println("Application successfully stopped")
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

type Budget struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Start string `json:"start"`
	End   string `json:"end"`
}

func (a *App) Budgets() ([]Budget, error) {
	budgets, err := a.repository.Budgets(a.ctx)
	if err != nil {
		return nil, err
	}

	var b []Budget
	for _, budget := range budgets {
		b = append(b, Budget{
			ID:    budget.ID,
			Name:  budget.Name,
			Start: budget.Start.Format(time.DateOnly),
			End:   budget.End.Format(time.DateOnly),
		})
	}

	return b, nil
}
