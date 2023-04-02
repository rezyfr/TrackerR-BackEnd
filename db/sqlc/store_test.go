package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTransactionTx(t *testing.T) {
	store := NewStore(testDB)

	user1 := createRandomUser(t)

	wallet1 := createRandomWallet(t)

	// run n concurrent new transaction
	n := 5
	amount := int64(10000)

	errs := make(chan error)
	results := make(chan NewTransactionTxResult)

	for i := 0; i < n; i++ {
		go func() {
			ctx := context.Background()
			result, err := store.CreateTransactionTx(ctx, NewTransactionTxParams{
				Amount:     amount,
				Type:       TransactiontypeDEBIT,
				UserID:     user1.ID,
				CategoryID: 5,
				WalletID:   wallet1.ID,
			})
			errs <- err
			results <- result
		}()
	}

	// Check n results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// Check transaction
		trx := result.Transaction
		require.Equal(t, amount, trx.Amount)
		require.Equal(t, wallet1.ID, result.Wallet.ID)
		require.Equal(t, user1.ID, trx.UserID.Int64)
		require.NotZero(t, trx.ID)
		require.NotZero(t, trx.CreatedAt)
	}

	// Check wallet
	updatedWallet, err := testQueries.GetWallet(context.Background(), wallet1.ID)
	require.NoError(t, err)
	require.Equal(t, wallet1.Balance-int64(n*int(amount)), updatedWallet.Balance)
}

func TestUpdateWalletTx(t *testing.T) {
	store := NewStore(testDB)
	wallet := createRandomWallet(t)

	ctx := context.Background()
	result, err := store.UpdateWalletTx(ctx, UpdateWalletTxParams{
		ID:     wallet.ID,
		Amount: 20000,
	})

	if result.Transaction.Type == TransactiontypeDEBIT {
		result.Transaction.Amount = -result.Transaction.Amount
	}
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, result.Wallet.Balance, wallet.Balance+result.Transaction.Amount)
}
