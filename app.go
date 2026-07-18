package main

import (
	"budgenius/repository"
	"context"
	"database/sql"
	"fmt"
	"time"
)

// App struct
type App struct {
	ctx        context.Context
	repository *repository.Repository
}

// NewApp creates a new App application struct
func NewApp(db *sql.DB, schema string) (*App, error) {
	r, err := repository.NewRepository(db, schema)
	if err != nil {
		return nil, err
	}
	return &App{
		repository: r,
	}, nil
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

type Budget struct {
	ID    int64  `json:"id"`
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
			Start: budget.Start.Format(time.DateTime),
			End:   budget.End.Format(time.DateTime),
		})
	}

	return b, nil
}
