package server_test

import (
	"context"
	"testing"

	"github.com/socialsalt/budget-app/internal/server"
	"go.uber.org/mock/gomock"
)

func SetUpTransLogicProv(ctrl *gomock.Controller) server.TransactionLogicProvider {

	r := server.NewMockTransactionRepo(ctrl)
	return server.TransactionLogicProvider{R: r}
}

func TestCreateTransactionsFromCSV(t *testing.T) {

	ctrl := gomock.NewController(t)
	tl := SetUpTransLogicProv(ctrl)

	testFilePath := "./test_transactions.cvs"
	transactions, err := tl.CreateTransactionsFromCVS(context.TODO(), testFilePath)

	ok(t, err)
	assert(t, len(transactions) == 0, "")

}
