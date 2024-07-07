package handler

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngqinzhe/parser/consts"
	"github.com/ngqinzhe/parser/dal/db_model"
	"github.com/ngqinzhe/parser/service"
)

type GetTransactionsHandler struct {
	parser service.Parser
}

func NewGetTransactionsHandler(parser service.Parser) *GetTransactionsHandler {
	return &GetTransactionsHandler{
		parser: parser,
	}
}

func (g *GetTransactionsHandler) Handle(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx = c.Request.Context()
		req := &db_model.GetTransactionsRequest{}
		if err := c.BindJSON(req); err != nil {
			slog.ErrorContext(ctx, fmt.Sprintf("[GetTransactionsHandler] unmarshal failed, err: %v", err))
			c.JSON(http.StatusBadRequest, &db_model.HttpErrorResponse{
				Code:    consts.UnmarshalError,
				Message: "unable to unmarshal to getTransactionRequest",
			})
			return
		}
		address := req.Address
		transactions := g.parser.GetTransactions(ctx, req.Address)
		if len(transactions) == 0 {
			slog.ErrorContext(ctx, "[GetTransactionsHandler] failed to get transactions")
			c.JSON(http.StatusInternalServerError, &db_model.HttpErrorResponse{
				Code:    consts.RpcError,
				Message: "failed to get transactions",
			})
			return
		}

		var allRelatedTransactions []*db_model.EthTransaction
		for _, transaction := range transactions {
			// if transaction is sent to current address
			if transaction.To().String() == address {
				allRelatedTransactions = append(allRelatedTransactions, &db_model.EthTransaction{
					Address:    address,
					IsSender:   false,
					Data:       transaction.Data(),
					Gas:        transaction.Gas(),
					Nonce:      transaction.Nonce(),
					CreateTime: transaction.Time(),
				})
			}
			// TODO: get transaction if transaction is sent "from" address
		}

		c.JSON(http.StatusOK, &db_model.GetTransactionsResponse{
			Transactions: allRelatedTransactions,
		})
	}
}
