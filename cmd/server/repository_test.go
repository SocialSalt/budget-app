package server_test

import (
	"context"
	"testing"
	"time"

	"github.com/socialsalt/budget-app/cmd/server"
)

var now time.Time = time.Now()
var testTransactions = []server.Transaction{
	{
		Date:            now,
		Company:         "test company",
		Category:        "groceries",
		Amount:          123456,
		AccountNumber:   "1234",
		Institution:     "idk",
		FullDescription: "this is a sentance",
		DateAdded:       now,
	},
	{
		Date:            now.Add(-1 * time.Hour * 10),
		Company:         "other test company",
		Category:        "misc",
		Amount:          789,
		AccountNumber:   "5678",
		Institution:     "idk",
		FullDescription: "this is a sentance",
		DateAdded:       now.Add(-1 * time.Hour * 10),
	},
	{
		Date:            now.Add(-1 * time.Hour * 5),
		Company:         "test company",
		Amount:          2,
		AccountNumber:   "1234",
		Institution:     "idk",
		FullDescription: "this is a sentance",
		DateAdded:       now.Add(-1 * time.Hour * 5),
	},
	{
		Date:            now.Add(time.Hour * 5),
		Company:         "test company",
		Amount:          10,
		AccountNumber:   "1234",
		Institution:     "idk",
		FullDescription: "this is a sentance",
		DateAdded:       now.Add(time.Hour * 5),
	},
}

func TestCreateTransactions(t *testing.T) {
	db := InitTestDB(t)

	transactionRepo := server.TransactionRepoProvider{DB: db}

	transactions, err := transactionRepo.CreateTransactions(context.TODO(), testTransactions)
	ok(t, err)
	assert(t, transactions[0].ID == 1, "failed to create transaction correctly")
	assert(t, transactions[1].ID == 2, "failed to create transaction correctly")
	assert(t, transactions[2].Category == "", "failed to create transaction correctly")
}

func TestListTransactions(t *testing.T) {
	db := InitTestDB(t)
	transactionRepo := server.TransactionRepoProvider{DB: db}

	_, err := transactionRepo.CreateTransactions(context.TODO(), testTransactions)
	ok(t, err)

	miscCat := "misc"
	tf := server.TransactionFilter{
		Category: &miscCat,
	}
	transactions, err := transactionRepo.ListTransactions(context.TODO(), tf)
	assert(t, len(transactions) == 1, "Failed to get only misc cat")
	assert(t, transactions[0].Category == "misc", "Failed to get only misc cat")

	var minAmount int64 = 11
	var maxAmount int64 = 123450
	tf = server.TransactionFilter{
		MinAmount: &minAmount,
		MaxAmount: &maxAmount,
	}
	transactions, err = transactionRepo.ListTransactions(context.TODO(), tf)
	assert(t, len(transactions) == 1, "Failed to filter by min or max amount")
	assert(t, transactions[0].Amount == 789, "Failed to filter by min or max amount")

	var after time.Time = time.Now().Add(-1 * time.Hour)
	var before time.Time = time.Now().Add(time.Hour)
	tf = server.TransactionFilter{
		StartDate: &after,
		EndDate:   &before,
	}
	transactions, err = transactionRepo.ListTransactions(context.TODO(), tf)
	assert(t, len(transactions) == 1, "Failed to filter by time")
	assert(t, transactions[0].Date.Sub(now) == 0, "Failed to filter by time")

	tf = server.TransactionFilter{}
	transactions, err = transactionRepo.ListTransactions(context.TODO(), tf)
	assert(t, len(transactions) == 4, "Failed to list all")
}
