package server

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
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
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO transactions(date, company, category, amount, account_number, institution, full_description, date_added) VALUES(?, ?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for i, t := range transactions {
		result, err := stmt.Exec(t.Date, t.Company, t.Category, t.Amount, t.AccountNumber, t.Institution, t.FullDescription, t.DateAdded)
		if err != nil {
			return nil, err
		}
		id, _ := result.LastInsertId()
		transactions[i].ID = id
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *TransactionRepoProvider) ListTransactions(ctx context.Context, tf TransactionFilter) ([]Transaction, error) {
	tx, err := r.DB.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	wheres := make([]string, 0, 6)

	if tf.Category != nil {
		wheres = append(wheres, fmt.Sprintf("category = '%s'", *tf.Category))
	}
	if tf.AccountNumber != nil {
		wheres = append(wheres, fmt.Sprintf("account_number = '%s'", *tf.AccountNumber))
	}
	if tf.MinAmount != nil {
		wheres = append(wheres, fmt.Sprintf("amount >= '%d'", *tf.MinAmount))
	}
	if tf.MaxAmount != nil {
		wheres = append(wheres, fmt.Sprintf("amount <= '%d'", *tf.MaxAmount))
	}
	if tf.StartDate != nil {
		wheres = append(wheres, fmt.Sprintf("date >= '%v'", *tf.StartDate))
	}
	if tf.EndDate != nil {
		wheres = append(wheres, fmt.Sprintf("date <= '%v'", *tf.EndDate))
	}

	whereClause := strings.Join(wheres, " AND ")
	if len(whereClause) > 0 {
		whereClause = fmt.Sprintf("WHERE %s", whereClause)
	}

	stmt, err := tx.Prepare(fmt.Sprintf("SELECT * FROM transactions %s;", whereClause))
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := make([]Transaction, 0)
	for rows.Next() {
		var t Transaction
		err := rows.Scan(&t.ID, &t.Date, &t.Company, &t.Category, &t.Amount, &t.AccountNumber, &t.Institution, &t.FullDescription, &t.DateAdded)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, t)
	}

	return transactions, nil
}
