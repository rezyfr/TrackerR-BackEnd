package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"

	"github.com/rezyfr/Trackerr-BackEnd/util"
)

// Store provide all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("tx err: %v, rollback err: %v", err, rollbackErr)
		}
		return err
	}
	return tx.Commit()
}

type NewTransactionTxParams struct {
	Amount     int64           `json:"amount"`
	Type       Transactiontype `json:"type"`
	UserID     int64           `json:"user_id"`
	CategoryID int64           `json:"category_id"`
	WalletID   int64           `json:"wallet_id"`
}

type NewTransactionTxResult struct {
	Transaction Transaction `json:"transaction"`
	Wallet      Wallet      `json:"wallet"`
	Category    Category    `json:"category"`
}

func (store *Store) CreateTransactionTx(ctx context.Context, arg NewTransactionTxParams) (NewTransactionTxResult, error) {
	var result NewTransactionTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transaction, err = q.CreateTransaction(ctx, CreateTransactionParams{
			Amount:     arg.Amount,
			Type:       arg.Type,
			UserID:     util.NullInt(int(arg.UserID)),
			WalletID:   util.NullInt(int(arg.WalletID)),
			CategoryID: arg.CategoryID,
		})
		if err != nil {
			return err
		}
		result.Category, err = q.GetCategory(ctx, arg.CategoryID)
		if err != nil {
			return err
		}

		if arg.Type == TransactiontypeDEBIT {
			arg.Amount = -arg.Amount
		}

		result.Wallet, err = q.AddWalletBalance(ctx, AddWalletBalanceParams{
			ID:     arg.WalletID,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}

type UpdateWalletTxParams struct {
	ID     int64 `json:"id"`
	Amount int64 `json:"amount"`
}

type UpdateWalletTxResult struct {
	Wallet      Wallet      `json:"wallet"`
	Transaction Transaction `json:"transaction"`
}

func (store *Store) UpdateWalletTx(ctx context.Context, arg UpdateWalletTxParams) (UpdateWalletTxResult, error) {
	var result UpdateWalletTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		currentWallet, err := q.GetWallet(ctx, arg.ID)
		if err != nil {
			return err
		}
		if currentWallet.Balance == arg.Amount {
			return errors.New("balance is the same")
		}
		result.Wallet, err = q.UpdateWalletBalance(ctx, UpdateWalletBalanceParams{
			ID:     arg.ID,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}
		transactionAmount := currentWallet.Balance - result.Wallet.Balance
		transactionType := TransactiontypeDEBIT

		if transactionAmount < 0 {
			transactionType = TransactiontypeCREDIT
		}

		result.Transaction, err = q.CreateTransaction(ctx, CreateTransactionParams{
			Amount:     int64(math.Abs(float64(transactionAmount))),
			Type:       transactionType,
			UserID:     currentWallet.UserID,
			WalletID:   util.NullInt(int(arg.ID)),
			CategoryID: 1,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}
