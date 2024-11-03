package server_test

import (
	"context"
	"testing"

	"github.com/socialsalt/budget-app/internal/server"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func SetUpTransLogicProv(ctrl *gomock.Controller) server.TransactionLogicProvider {

	r := server.NewMockTransactionRepo(ctrl)
	return server.TransactionLogicProvider{TransactionRepo: r}
}

func TestCreateTransactionsFromCSV(t *testing.T) {

	ctrl := gomock.NewController(t)
	tl := SetUpTransLogicProv(ctrl)

	testFilePath := "./test_transactions.cvs"
	transactions, err := tl.CreateTransactionsFromCVS(context.TODO(), testFilePath)

	assert.NoError(t, err)
	assert.Equal(t, 0, len(transactions), "")
}
