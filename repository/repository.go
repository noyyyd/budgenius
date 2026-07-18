package repository

import (
	_ "embed"

	"context"
	"database/sql"
	"time"
)

//go:embed schema.sql
var schema string

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) (*Repository, error) {
	_, err := db.Exec(schema)
	if err != nil {
		return nil, err
	}
	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) Budgets(ctx context.Context) ([]Budget, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, start, end
		FROM budget
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bugets []Budget
	for rows.Next() {
		var (
			b     Budget
			start int64
			end   int64
		)
		err := rows.Scan(&b.ID, &start, &end)
		if err != nil {
			return nil, err
		}
		b.Start = time.Unix(start, 0).Local()
		b.End = time.Unix(end, 0).Local()
		bugets = append(bugets, b)
	}
	return bugets, rows.Err()
}

func (r *Repository) Categories(ctx context.Context) ([]Category, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, name, type
		FROM category
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var c Category
		err := rows.Scan(&c.ID, &c.Name, &c.Type)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, rows.Err()
}

func (r *Repository) Transactions(ctx context.Context) ([]Transaction, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, ts, amount, comment, c.name
		FROM transaction t
			JOIN category c ON t.category_id = c.id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []Transaction
	for rows.Next() {
		var (
			t  Transaction
			ts int64
		)
		err := rows.Scan(&t.ID, &ts, &t.Amount, &t.Comment, &t.Category)
		if err != nil {
			return nil, err
		}
		t.TS = time.Unix(ts, 0).Local()
		transactions = append(transactions, t)
	}
	return transactions, rows.Err()
}
