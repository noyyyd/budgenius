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
		SELECT id, name, start, end
		FROM budget
		ORDER BY id DESC
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
		err := rows.Scan(&b.ID, &b.Name, &start, &end)
		if err != nil {
			return nil, err
		}
		b.Start = time.Unix(start, 0).Local()
		b.End = time.Unix(end, 0).Local()
		bugets = append(bugets, b)
	}
	return bugets, rows.Err()
}

func (r *Repository) CreateBudget(ctx context.Context, name string, start, end time.Time) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO budget (name, start, end)
		VALUES  (?, ?, ?)`,
		name, start.Unix(), end.Unix())
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) CreateCategory(ctx context.Context, name string, categoryType CategoryType) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO category (name, type)
		VALUES  (?, ?)`,
		name, categoryType)
	return err
}

func (r *Repository) UpdateCategory(ctx context.Context, id int64, name string, categoryType CategoryType) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE category SET name = ?, type = ?
		WHERE id = ?`,
		name, categoryType, id)
	return err
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

func (r *Repository) DeleteCategory(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM category WHERE id = ?`, id)
	return err
}

func (r *Repository) CreateTransaction(ctx context.Context, ts time.Time, amount int64, comment string, categoryID int64) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO money_transaction (ts, amount, comment, category_id)
		VALUES (?, ?, ?, ?)`,
		ts.Unix(), amount, comment, categoryID)
	return err
}

func (r *Repository) UpdateTransaction(ctx context.Context, id int64, ts time.Time, amount int64, comment string, categoryID int64) error {
	_, err := r.db.ExecContext(ctx, `
		UPDATE money_transaction SET ts = ?, amount = ?, comment = ?, category_id = ?
		WHERE id = ?`,
		ts.Unix(), amount, comment, categoryID, id)
	return err
}

func (r *Repository) Transactions(ctx context.Context) ([]Transaction, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT t.id, ts, amount, comment, c.id, c.name, c.type
		FROM money_transaction t
			JOIN category c ON t.category_id = c.id
	`) // TODO сортировка транзакций
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
		err := rows.Scan(&t.ID, &ts, &t.Amount, &t.Comment, &t.Category.ID, &t.Category.Name, &t.Category.Type)
		if err != nil {
			return nil, err
		}
		t.TS = time.Unix(ts, 0).Local()
		transactions = append(transactions, t)
	}
	return transactions, rows.Err()
}

func (r *Repository) DeleteTransaction(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM money_transaction WHERE id = ?`, id)
	return err
}
