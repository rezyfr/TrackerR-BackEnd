package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/rezyfr/Trackerr-BackEnd/db/sqlc"
	"github.com/rezyfr/Trackerr-BackEnd/token"
)

type listTransactionRequest struct {
	Limit  int32 `form:"page_limit,default=10" binding:"max=50"`
	Offset int32 `form:"page_offset,default=0"`
}

func (server *Server) listTransactions(ctx *gin.Context) {
	var req listTransactionRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.ListTransactionsParams{
		UserID: int64(authPayload.UserID),
		Limit:  req.Limit,
		Offset: (req.Offset - 1) * req.Limit,
	}
	transactions, err := server.store.ListTransactions(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

type createTransactionRequest struct {
	Amount     int64  `json:"amount" binding:"required"`
	Note       string `json:"note" binding:"required"`
	Type       string `json:"type" binding:"required,trxtype"`
	CategoryID int64  `json:"category_id" binding:"required"`
	WalletID   int64  `json:"wallet_id" binding:"required"`
}

func (server *Server) createTransaction(ctx *gin.Context) {
	var req createTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !server.validateWallet(ctx, req.WalletID) {
		return
	}

	if !server.validateCategory(ctx, req.CategoryID) {
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.NewTransactionTxParams{
		Amount:     req.Amount,
		Note:       req.Note,
		Type:       db.Transactiontype(req.Type),
		CategoryID: req.CategoryID,
		UserID:     int64(authPayload.UserID),
		WalletID:   req.WalletID,
	}

	result, err := server.store.CreateTransactionTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validateWallet(ctx *gin.Context, walletID int64) bool {
	wallet, err := server.store.GetWallet(ctx, walletID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	return wallet.ID == walletID
}

func (server *Server) validateCategory(ctx *gin.Context, categoryId int64) bool {
	category, err := server.store.GetCategory(ctx, categoryId)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}

	return category.ID == categoryId
}
