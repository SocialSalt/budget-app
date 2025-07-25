package dataaccess_test

import (
	"context"
	"testing"
	"time"

	"github.com/socialsalt/budget-app/internal/data_access"
	"github.com/socialsalt/budget-app/internal/model"
	"github.com/stretchr/testify/assert"
)

var now time.Time = time.Now()
var testTransactions = []model.Transaction{
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

	transactionDAL := dataaccess.TransactionDAL{DB: db}

	transactions, err := transactionDAL.CreateTransactions(context.TODO(), testTransactions)
	assert.NoError(t, err)
	assert.Equal(t, int64(1), transactions[0].ID, "failed to create transaction correctly")
	assert.Equal(t, int64(2), transactions[1].ID, "failed to create transaction correctly")
	assert.Equal(t, "", transactions[2].Category, "failed to create transaction correctly")
}

func TestListTransactions(t *testing.T) {
	db := InitTestDB(t)
	transactionDAL := dataaccess.TransactionDAL{DB: db}

	_, err := transactionDAL.CreateTransactions(context.TODO(), testTransactions)
	assert.NoError(t, err)

	miscCat := "misc"
	tf := model.TransactionFilter{
		Category: &miscCat,
	}
	transactions, err := transactionDAL.ListTransactions(context.TODO(), tf)
	assert.Equal(t, 1, len(transactions), "Failed to get only misc cat")
	assert.Equal(t, "misc", transactions[0].Category, "Failed to get only misc cat")

	var minAmount int64 = 11
	var maxAmount int64 = 123450
	tf = model.TransactionFilter{
		MinAmount: &minAmount,
		MaxAmount: &maxAmount,
	}
	transactions, err = transactionDAL.ListTransactions(context.TODO(), tf)
	assert.Equal(t, 1, len(transactions), "Failed to filter by min or max amount")
	assert.Equal(t, int64(789), transactions[0].Amount, "Failed to filter by min or max amount")

	var after time.Time = time.Now().Add(-1 * time.Hour)
	var before time.Time = time.Now().Add(time.Hour)
	tf = model.TransactionFilter{
		StartDate: &after,
		EndDate:   &before,
	}
	transactions, err = transactionDAL.ListTransactions(context.TODO(), tf)
	assert.Equal(t, 1, len(transactions), "Failed to filter by time")
	assert.Equal(t, time.Duration(0), transactions[0].Date.Sub(now), "Failed to filter by time")

	tf = model.TransactionFilter{}
	transactions, err = transactionDAL.ListTransactions(context.TODO(), tf)
	assert.Equal(t, 4, len(transactions), "Failed to list all")
}
