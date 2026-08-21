package main

import (
	"budgenius/repository"
	"fmt"
	"log"
	"time"
)

type CreateTransactionRequest struct {
	Date       string
	Amount     int64
	Comment    string
	CategoryID int64
}

func (a *App) CreateTransaction(t CreateTransactionRequest) error {
	log.Printf("Create transaction: %+v\n", t)

	date, err := time.Parse(time.DateOnly, t.Date)
	if err != nil {
		log.Println("Error parsing:", err)
		return err
	}

	err = a.repository.CreateTransaction(a.ctx, date, t.Amount, t.Comment, t.CategoryID)
	if err != nil {
		log.Println("Error creating category:", err)
		return err
	}

	log.Println("Transaction created successfully")

	return nil
}

type Transaction struct {
	ID       int64
	Date     string
	Amount   string
	Comment  string
	IsIncome bool
	Category string
}

func (a *App) Transactions() ([]Transaction, error) {
	log.Println("Request transactions")

	transactions, err := a.repository.Transactions(a.ctx)
	if err != nil {
		log.Println("Error getting transactions:", err)
		return nil, err
	}

	t := make([]Transaction, len(transactions))
	for i, transaction := range transactions {
		t[i] = Transaction{
			ID:       transaction.ID,
			Date:     transaction.TS.Format(time.DateOnly),
			Amount:   FormatMoney(transaction.Amount), // TODO преобразовывать копейки в рубли на фронте
			Comment:  transaction.Comment,
			IsIncome: transaction.Category.Type == repository.CategoryTypeIncome,
			Category: transaction.Category.Name,
		}
	}

	log.Printf("Successfully got %d transactions\n", len(t))

	return t, nil
}

func FormatMoney(amount int64) string {
	rubles := amount / 100
	kopecks := amount % 100

	return fmt.Sprintf("%d,%02d", rubles, kopecks)
}

func (a *App) DeleteTransaction(id int64) error {
	log.Printf("Delete transaction with id %d\n", id)

	err := a.repository.DeleteTransaction(a.ctx, id)
	if err != nil {
		log.Println("Error deleting transaction:", err)
		return err
	}

	log.Printf("Transaction with id %d deleted successfully\n", id)

	return nil
}
