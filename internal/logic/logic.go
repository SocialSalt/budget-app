//go:generate mockgen -source=logic.go -destination=../../mocks/logic.go -package=mocks
package logic

import (
	"context"
	"encoding/csv"
	"os"

	"github.com/socialsalt/budget-app/internal/data_access"
	"github.com/socialsalt/budget-app/internal/model"
	"github.com/socialsalt/budget-app/internal/utils"
)

type TransactionLogic interface {
	CreateTransactionsFromCSV(ctx context.Context, filepath string) ([]model.Transaction, error)
}

type TransactionLogicLayer struct {
	TransactionDAL dataaccess.TransactionDA
}

func (t *TransactionLogicLayer) CreateTransactionsFromCVS(ctx context.Context, filepath string) ([]model.Transaction, error) {

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	transactions, err := utils.ParseTransactionCSV(data)
	if err != nil {
		return nil, err
	}
	transactions, err = t.TransactionDAL.CreateTransactions(ctx, transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
