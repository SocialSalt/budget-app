package logic_test

import (
	"context"
	"testing"
	"time"

	"github.com/socialsalt/budget-app/internal/logic"
	"github.com/socialsalt/budget-app/internal/model"
	"github.com/socialsalt/budget-app/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var exptectedTransactions []model.Transaction = []model.Transaction{
	{
		Date:          time.Date(2024, 11, 4, 0, 0, 0, 0, nil),
		Company:       "Internet Transfer from Jpmorgan Chase Bank, NaDda Account x8237",
		Category:      "Money Transfer",
		Amount:        50000,
		AccountNumber: " 1345",
		Institution:   "American Express",
	},
}

func TestCreateTransactionsFromCSV(t *testing.T) {

	ctrl := gomock.NewController(t)
	r := mocks.NewMockTransactionDA(ctrl)
	r.EXPECT().CreateTransactions(gomock.Any(), gomock.Any())
	tl := logic.TransactionLogicLayer{TransactionDAL: r}

	testFilePath := "./test_transactions.csv"
	transactions, err := tl.CreateTransactionsFromCVS(context.TODO(), testFilePath)

	assert.NoError(t, err)
	assert.Equal(t, 0, len(transactions), "")
}
