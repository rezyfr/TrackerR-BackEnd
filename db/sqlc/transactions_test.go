package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/rezyfr/Trackerr-BackEnd/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransaction(t *testing.T, userId int) Transaction {
	arg := CreateTransactionParams{
		Amount:     util.RandomAmount(),
		Type:       Transactiontype(util.RandomType()),
		UserID:     util.NullInt(userId),
		CategoryID: 5,
		WalletID:   util.NullInt(1),
	}

	transaction, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, arg.Amount, transaction.Amount)

	return transaction
}

func TestCreateTransaction(t *testing.T) {
	arg := CreateTransactionParams{
		Amount:     util.RandomAmount(),
		Type:       Transactiontype(util.RandomType()),
		UserID:     util.NullInt(1),
		CategoryID: 5,
		WalletID:   util.NullInt(1),
	}

	trx, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, trx)

	require.Equal(t, arg.Type, trx.Type)
	require.Equal(t, arg.Amount, trx.Amount)
}

func TestListTransactions(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomTransaction(t, 1)
	}

	arg := ListTransactionsParams{
		Limit:  5,
		Offset: 5,
	}

	transactions, err := testQueries.ListTransactions(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transactions)
	require.Len(t, transactions, 5)

	for _, trx := range transactions {
		require.NotEmpty(t, trx)
	}
}

func TestUpdateTransaction(t *testing.T) {
	trx := createRandomTransaction(t, 1)

	arg := UpdateTransactionParams{
		ID:         trx.ID,
		Amount:     util.RandomAmount(),
		Type:       Transactiontype(util.RandomType()),
		CategoryID: trx.CategoryID,
		WalletID:   trx.WalletID,
	}

	trx2, err := testQueries.UpdateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, trx2)

	require.Equal(t, arg.ID, trx2.ID)
	require.Equal(t, arg.Amount, trx2.Amount)
	require.Equal(t, arg.Type, trx2.Type)
	require.Equal(t, arg.CategoryID, trx2.CategoryID)
	require.Equal(t, arg.WalletID, trx2.WalletID)
}

func TestDeleteTransaction(t *testing.T) {
	trx := createRandomTransaction(t, 1)

	err := testQueries.DeleteTransaction(context.Background(), trx.ID)
	require.NoError(t, err)

	trx2, err := testQueries.GetTransaction(context.Background(), trx.ID)
	require.Error(t, err)
	require.Empty(t, trx2)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
