// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: transactions.sql

package db

import (
	"context"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (
  user_id, 
  amount,
  type,
  note,
  category_id,
  wallet_id
) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING id, user_id, amount, note, created_at, updated_at, type, category_id, wallet_id
`

type CreateTransactionParams struct {
	UserID     int64           `json:"user_id"`
	Amount     int64           `json:"amount"`
	Type       Transactiontype `json:"type"`
	Note       string          `json:"note"`
	CategoryID int64           `json:"category_id"`
	WalletID   int64           `json:"wallet_id"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, createTransaction,
		arg.UserID,
		arg.Amount,
		arg.Type,
		arg.Note,
		arg.CategoryID,
		arg.WalletID,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Amount,
		&i.Note,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Type,
		&i.CategoryID,
		&i.WalletID,
	)
	return i, err
}

const deleteTransaction = `-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1
`

func (q *Queries) DeleteTransaction(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTransaction, id)
	return err
}

const getTransaction = `-- name: GetTransaction :one
SELECT id, user_id, amount, note, created_at, updated_at, type, category_id, wallet_id FROM transactions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransaction(ctx context.Context, id int64) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, getTransaction, id)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Amount,
		&i.Note,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Type,
		&i.CategoryID,
		&i.WalletID,
	)
	return i, err
}

const listTransactions = `-- name: ListTransactions :many
SELECT id, user_id, amount, note, created_at, updated_at, type, category_id, wallet_id FROM transactions 
WHERE user_id = $1
ORDER BY created_at
LIMIT $2 OFFSET $3
`

type ListTransactionsParams struct {
	UserID int64 `json:"user_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTransactions(ctx context.Context, arg ListTransactionsParams) ([]Transaction, error) {
	rows, err := q.db.QueryContext(ctx, listTransactions, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Amount,
			&i.Note,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Type,
			&i.CategoryID,
			&i.WalletID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTransaction = `-- name: UpdateTransaction :one
UPDATE transactions SET
  amount = $1,
  type = $2,
  category_id = $3,
  wallet_id = $4
WHERE id = $5
RETURNING id, user_id, amount, note, created_at, updated_at, type, category_id, wallet_id
`

type UpdateTransactionParams struct {
	Amount     int64           `json:"amount"`
	Type       Transactiontype `json:"type"`
	CategoryID int64           `json:"category_id"`
	WalletID   int64           `json:"wallet_id"`
	ID         int64           `json:"id"`
}

func (q *Queries) UpdateTransaction(ctx context.Context, arg UpdateTransactionParams) (Transaction, error) {
	row := q.db.QueryRowContext(ctx, updateTransaction,
		arg.Amount,
		arg.Type,
		arg.CategoryID,
		arg.WalletID,
		arg.ID,
	)
	var i Transaction
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Amount,
		&i.Note,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Type,
		&i.CategoryID,
		&i.WalletID,
	)
	return i, err
}
