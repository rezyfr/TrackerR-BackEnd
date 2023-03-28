package util

import (
	"database/sql"
	"math/rand"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func RandomFullName() string {
	return RandomString(10) + " " + RandomString(10)
}

func RandomEmail() string {
	return RandomString(10) + "@gmail.com"
}

func RandomAmount() int64 {
	return RandomInt(1000, 1000000)
}

func RandomType() Transactiontype {
	randomInt := RandomInt(1, 10)
	switch {
	case randomInt%2 == 0:
		return TransactiontypeDEBIT
	case randomInt%2 == 1:
		return TransactiontypeCREDIT
	}
	return TransactiontypeCREDIT
}

func NullInt(n int) sql.NullInt64 {
	return sql.NullInt64{
		Int64: int64(n),
		Valid: true,
	}
}

type Transactiontype string

const (
	TransactiontypeDEBIT    Transactiontype = "DEBIT"
	TransactiontypeCREDIT   Transactiontype = "CREDIT"
	TransactiontypeTRANSFER Transactiontype = "TRANSFER"
)
