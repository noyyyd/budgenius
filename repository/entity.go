package repository

import "time"

type CategoryType string

const (
	CategoryTypeIncome   = "income"
	CategoryTypeExpenses = "expenses"
)

func (t CategoryType) IsValid() bool {
	switch t {
	case CategoryTypeExpenses, CategoryTypeIncome:
		return true
	}
	return false
}

type Budget struct {
	ID    int64
	Name  string
	Start time.Time
	End   time.Time
}

type Category struct {
	ID   int64
	Name string
	Type CategoryType
}

type Transaction struct {
	ID       int64
	TS       time.Time
	Amount   int64
	Comment  string
	Category Category
}
