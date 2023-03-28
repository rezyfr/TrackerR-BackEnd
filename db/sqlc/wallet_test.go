package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/rezyfr/Trackerr-BackEnd/util"
	"github.com/stretchr/testify/require"
)

func createRandomWallet(t *testing.T) Wallet {
	arg := CreateWalletParams{
		UserID:  util.NullInt(1),
		Balance: util.RandomAmount(),
		Icon:    util.RandomString(5),
		Name:    util.RandomString(5),
	}

	wallet, err := testQueries.CreateWallet(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, wallet)

	require.Equal(t, arg.Balance, wallet.Balance)
	require.Equal(t, arg.Icon, wallet.Icon)
	require.Equal(t, arg.UserID, wallet.UserID)
	require.Equal(t, arg.Name, wallet.Name)

	return wallet
}

func TestCreateWallet(t *testing.T) {
	arg := CreateWalletParams{
		UserID:  util.NullInt(1),
		Balance: util.RandomAmount(),
		Icon:    util.RandomString(5),
	}

	wallet, err := testQueries.CreateWallet(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, wallet)

	require.Equal(t, arg.Balance, wallet.Balance)
	require.Equal(t, arg.Icon, wallet.Icon)
	require.Equal(t, arg.UserID, wallet.UserID)
}

func TestListWallets(t *testing.T) {
	for i := 0; i < 3; i++ {
		createRandomWallet(t)
	}

	arg := util.NullInt(1)

	wallets, err := testQueries.ListWallets(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, wallets)

	for _, wallet := range wallets {
		require.NotEmpty(t, wallet)
	}
}

func TestUpdateWallet(t *testing.T) {
	wallet1 := createRandomWallet(t)

	arg := UpdateWalletParams{
		ID:      wallet1.ID,
		Balance: util.RandomAmount(),
		Icon:    util.RandomString(5),
	}

	wallet2, err := testQueries.UpdateWallet(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, wallet2)

	require.Equal(t, arg.Balance, wallet2.Balance)
	require.Equal(t, arg.Icon, wallet2.Icon)
	require.Equal(t, arg.ID, wallet2.ID)
}

func TestDeleteWalelt(t *testing.T) {
	wallet1 := createRandomWallet(t)

	err := testQueries.DeleteWallet(context.Background(), wallet1.ID)
	require.NoError(t, err)

	wallet2, err := testQueries.GetWallet(context.Background(), wallet1.ID)
	require.Error(t, err)
	require.Empty(t, wallet2)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}
