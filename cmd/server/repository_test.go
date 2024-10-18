package server_test

import (
	"context"
	"testing"
	"time"

	"github.com/socialsalt/budget-app/cmd/server"
)

func TestCreateTransactions(t *testing.T) {
	db := InitTestDB(t)

	transactions := []server.Transaction{
		server.Transaction{
			Date:            time.Now(),
			Company:         "test company",
			Category:        "groceries",
			Amount:          123456,
			AccountNumber:   "1234",
			Institution:     "idk",
			FullDescription: "this is a sentance",
			DateAdded:       time.Unix(1234567, 0),
		},
		server.Transaction{
			Date:            time.Now(),
			Company:         "test company",
			Category:        "groceries",
			Amount:          123456,
			AccountNumber:   "1234",
			Institution:     "idk",
			FullDescription: "this is a sentance",
			DateAdded:       time.Unix(1234567, 0),
		},
	}

	transactionRepo := server.TransactionRepoProvider{DB: db}

	transactions, err := transactionRepo.CreateTransactions(context.TODO(), transactions)
	ok(t, err)
	assert(t, transactions[0].ID == 1, "failed to create transaction correctly")
	assert(t, transactions[1].ID == 2, "failed to create transaction correctly")
}
