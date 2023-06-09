// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: wallet.sql

package db

import (
	"context"
)

const addWalletBalance = `-- name: AddWalletBalance :one
UPDATE wallet SET
  balance = balance + $1
WHERE id = $2
RETURNING id, user_id, name, balance, icon, created_at, updated_at
`

type AddWalletBalanceParams struct {
	Amount int64 `json:"amount"`
	ID     int64 `json:"id"`
}

func (q *Queries) AddWalletBalance(ctx context.Context, arg AddWalletBalanceParams) (Wallet, error) {
	row := q.db.QueryRowContext(ctx, addWalletBalance, arg.Amount, arg.ID)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Balance,
		&i.Icon,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createWallet = `-- name: CreateWallet :one
INSERT INTO wallet (
  user_id, 
  name,
  balance,
  icon
) VALUES (
  $1, $2, $3, $4
) RETURNING id, user_id, name, balance, icon, created_at, updated_at
`

type CreateWalletParams struct {
	UserID  int64  `json:"user_id"`
	Name    string `json:"name"`
	Balance int64  `json:"balance"`
	Icon    string `json:"icon"`
}

func (q *Queries) CreateWallet(ctx context.Context, arg CreateWalletParams) (Wallet, error) {
	row := q.db.QueryRowContext(ctx, createWallet,
		arg.UserID,
		arg.Name,
		arg.Balance,
		arg.Icon,
	)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Balance,
		&i.Icon,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteWallet = `-- name: DeleteWallet :exec
DELETE FROM wallet
WHERE id = $1
`

func (q *Queries) DeleteWallet(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteWallet, id)
	return err
}

const getWallet = `-- name: GetWallet :one
SELECT id, user_id, name, balance, icon, created_at, updated_at FROM wallet
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetWallet(ctx context.Context, id int64) (Wallet, error) {
	row := q.db.QueryRowContext(ctx, getWallet, id)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Balance,
		&i.Icon,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getWalletForUpdate = `-- name: GetWalletForUpdate :one
SELECT id, user_id, name, balance, icon, created_at, updated_at FROM wallet
WHERE id = $1 FOR NO KEY UPDATE
`

func (q *Queries) GetWalletForUpdate(ctx context.Context, id int64) (Wallet, error) {
	row := q.db.QueryRowContext(ctx, getWalletForUpdate, id)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Balance,
		&i.Icon,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listWallets = `-- name: ListWallets :many
SELECT id, user_id, name, balance, icon, created_at, updated_at FROM wallet 
WHERE user_id = $1
ORDER BY name
`

func (q *Queries) ListWallets(ctx context.Context, userID int64) ([]Wallet, error) {
	rows, err := q.db.QueryContext(ctx, listWallets, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Wallet{}
	for rows.Next() {
		var i Wallet
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.Balance,
			&i.Icon,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateWallet = `-- name: UpdateWallet :one
UPDATE wallet SET
  name = $1,
  balance = $2,
  icon = $3
WHERE id = $4
RETURNING id, user_id, name, balance, icon, created_at, updated_at
`

type UpdateWalletParams struct {
	Name    string `json:"name"`
	Balance int64  `json:"balance"`
	Icon    string `json:"icon"`
	ID      int64  `json:"id"`
}

func (q *Queries) UpdateWallet(ctx context.Context, arg UpdateWalletParams) (Wallet, error) {
	row := q.db.QueryRowContext(ctx, updateWallet,
		arg.Name,
		arg.Balance,
		arg.Icon,
		arg.ID,
	)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Balance,
		&i.Icon,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateWalletBalance = `-- name: UpdateWalletBalance :one
UPDATE wallet SET
  balance = $1
WHERE id = $2
RETURNING id, user_id, name, balance, icon, created_at, updated_at
`

type UpdateWalletBalanceParams struct {
	Amount int64 `json:"amount"`
	ID     int64 `json:"id"`
}

func (q *Queries) UpdateWalletBalance(ctx context.Context, arg UpdateWalletBalanceParams) (Wallet, error) {
	row := q.db.QueryRowContext(ctx, updateWalletBalance, arg.Amount, arg.ID)
	var i Wallet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.Balance,
		&i.Icon,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
