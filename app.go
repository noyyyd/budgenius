package main

import (
	"budgenius/repository"
	"context"
	"database/sql"
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

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

	f, err := os.OpenFile(filepath.Join(dataDir, "app.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		a.fatalErrorDialog(ctx, err)
	}
	log.SetOutput(f)

	log.Println("Start application")

	dbPath := filepath.Join(dataDir, "db.sqlite")

	log.Println("Connect to database:", dbPath)
	a.db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Println("Error connection to database:", err)
		a.fatalErrorDialog(ctx, err)
	}
	a.repository, err = repository.NewRepository(a.db)
	if err != nil {
		log.Println("Error creating repository:", err)
		a.fatalErrorDialog(ctx, err)
	}
}

func (a *App) fatalErrorDialog(ctx context.Context, err error) {
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

type Budget struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Start string `json:"start"`
	End   string `json:"end"`
}

func (a *App) BudgetList() ([]Budget, error) {
	log.Println("Request budget list")

	budgets, err := a.repository.Budgets(a.ctx)
	if err != nil {
		log.Println("Error getting budgets:", err)
		return nil, err
	}

	b := make([]Budget, len(budgets))
	for i, budget := range budgets {
		b[i] = Budget{
			ID:    budget.ID,
			Name:  budget.Name,
			Start: budget.Start.Format(time.DateOnly),
			End:   budget.End.Format(time.DateOnly),
		}
	}

	log.Printf("Successfully got %d budgets\n", len(b))

	return b, nil
}

func (a *App) CreateBudget(budget Budget) error {
	start, err := time.Parse(time.DateOnly, budget.Start)
	if err != nil {
		log.Println("Error parsing:", err)
		return err
	}
	end, err := time.Parse(time.DateOnly, budget.End)
	if err != nil {
		log.Println("Error parsing:", err)
		return err
	}

	err = a.repository.CreateBudget(a.ctx, budget.Name, start, end)
	if err != nil {
		log.Println("Error creating budget:", err)
		return err
	}

	log.Printf("Budget %q with range %s and %s successfully created\n", budget.Name, budget.Start, budget.End)

	return nil
}
