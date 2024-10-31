package server

import "context"

type TransactionLogic interface {
	CreateTransactionsFromCSV(ctx context.Context, filepath string) ([]Transaction, error)
}

type TransactionLogicProvider struct {
	r TransactionRepoProvider
}
