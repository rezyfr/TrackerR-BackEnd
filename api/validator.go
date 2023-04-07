package api

import (
	"github.com/go-playground/validator/v10"
	db "github.com/rezyfr/Trackerr-BackEnd/db/sqlc"
)

var validType validator.Func = func(fl validator.FieldLevel) bool {
	if trxType, ok := fl.Field().Interface().(string); ok {
		switch trxType {
		case string(db.TransactiontypeCREDIT), string(db.TransactiontypeDEBIT), string(db.TransactiontypeTRANSFER):
			return true
		}
		return false
	}
	return false
}
