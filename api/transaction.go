package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/rezyfr/Trackerr-BackEnd/db/sqlc"
)

type listTransactionRequest struct {
	Limit  int32 `form:"page_limit,default=10" binding:"max=50"`
	Offset int32 `form:"page_offset,default=0"`
	UserID int64 `form:"user_id,default=0"`
}

func (server *Server) listTransactions(ctx *gin.Context) {
	var req listTransactionRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListTransactionsParams{
		UserID: req.UserID,
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
