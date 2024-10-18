package server

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	// "log"
	// "os"
)

type TransactionRepo interface {
	CreateTransactions(ctx context.Context, transactions []Transaction) ([]Transaction, error)
	ListTransactions(ctx context.Context, filter TransactionFilter) ([]Transaction, error)
}

type TransactionRepoProvider struct {
	DB *sql.DB
}

func (r *TransactionRepoProvider) CreateTransactions(ctx context.Context, transactions []Transaction) ([]Transaction, error) {
	tx, err := r.DB.BeginTx(ctx, &sql.TxOptions{ReadOnly: false})
	if err != nil {
		return []Transaction{}, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO transactions(date, company, category, amount, account_number, institution, full_description, date_added) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return []Transaction{}, err
	}
	defer stmt.Close()

	for i, t := range transactions {
		result, err := stmt.Exec(t.Date, t.Company, t.Category, t.Amount, t.AccountNumber, t.Institution, t.FullDescription, t.DateAdded)
		if err != nil {
			return []Transaction{}, err
		}
		id, _ := result.LastInsertId()
		fmt.Printf("%#v\n", id)
		transactions[i].ID = id
	}

	if err := tx.Commit(); err != nil {
		return []Transaction{}, err
	}

	return transactions, nil
}
