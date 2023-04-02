package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/rezyfr/Trackerr-BackEnd/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransaction(t *testing.T) (Transaction, error) {
	var transaction Transaction
	user, err := testQueries.CreateUser(context.Background(), CreateUserParams{
		Email:    util.RandomEmail(),
		FullName: util.RandomString(5),
	})
	if err != nil {
		return transaction, err
	}
	category, err := testQueries.CreateCategory(context.Background(), CreateCategoryParams{
		UserID: util.NullInt(int(user.ID)),
		Type:   Transactiontype(util.RandomType()),
		Icon:   util.RandomString(5),
		Name:   util.RandomString(5),
	})
	if err != nil {
		return transaction, err
	}
	wallet, err := testQueries.CreateWallet(context.Background(), CreateWalletParams{
		UserID: util.NullInt(int(user.ID)),
		Name:   util.RandomString(5),
	})
	if err != nil {
		return transaction, err
	}
	arg := CreateTransactionParams{
		Amount:     util.RandomAmount(),
		Type:       Transactiontype(util.RandomType()),
		UserID:     util.NullInt(int(user.ID)),
		CategoryID: category.ID,
		WalletID:   util.NullInt(int(wallet.ID)),
	}

	transaction, err = testQueries.CreateTransaction(context.Background(), arg)
	if err != nil {
		return transaction, err
	}

	return transaction, err
}

func TestCreateTransaction(t *testing.T) {
	user, err := testQueries.CreateUser(context.Background(), CreateUserParams{
		Email:    util.RandomEmail(),
		FullName: util.RandomString(5),
	})
	require.NoError(t, err)
	category, err := testQueries.CreateCategory(context.Background(), CreateCategoryParams{
		UserID: util.NullInt(int(user.ID)),
		Type:   Transactiontype(util.RandomType()),
		Icon:   util.RandomString(5),
		Name:   util.RandomString(5),
	})
	require.NoError(t, err)
	wallet, err := testQueries.CreateWallet(context.Background(), CreateWalletParams{
		UserID: util.NullInt(int(user.ID)),
		Name:   util.RandomString(5),
	})
	require.NoError(t, err)
	arg := CreateTransactionParams{
		Amount:     util.RandomAmount(),
		Type:       Transactiontype(util.RandomType()),
		UserID:     util.NullInt(int(user.ID)),
		CategoryID: category.ID,
		WalletID:   util.NullInt(int(wallet.ID)),
	}

	trx, err := testQueries.CreateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, trx)

	require.Equal(t, arg.Type, trx.Type)
	require.Equal(t, arg.Amount, trx.Amount)
}

func TestListTransactions(t *testing.T) {
	for i := 0; i < 5; i++ {
		createRandomTransaction(t)
	}

	arg := ListTransactionsParams{
		Limit:  5,
		Offset: 0,
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
	trx, err := createRandomTransaction(t)
	require.NoError(t, err)

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
	trx, err := createRandomTransaction(t)
	require.NoError(t, err)

	err2 := testQueries.DeleteTransaction(context.Background(), trx.ID)
	require.NoError(t, err2)

	trx2, err := testQueries.GetTransaction(context.Background(), trx.ID)
	require.Error(t, err)
	require.Empty(t, trx2)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
