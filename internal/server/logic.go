package server

import (
	"context"
	"encoding/csv"
	"os"
)

type TransactionLogic interface {
	CreateTransactionsFromCSV(ctx context.Context, filepath string) ([]Transaction, error)
}

type TransactionLogicProvider struct {
	r TransactionRepo
}

func (t *TransactionLogicProvider) CreateTransactionsFromCVS(ctx context.Context, filepath string) ([]Transaction, error) {

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	transactions, err := ParseTransactionCSV(data)
	if err != nil {
		return nil, err
	}
	transactions, err = t.r.CreateTransactions(ctx, transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
