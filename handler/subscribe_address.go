package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ngqinzhe/parser/consts"
	"github.com/ngqinzhe/parser/dal/db_model"
	"github.com/ngqinzhe/parser/service"
)

type SubscribeAddressHandler struct {
	parser service.Parser
}

func NewSubscribeAddressHandler(parser service.Parser) *SubscribeAddressHandler {
	return &SubscribeAddressHandler{
		parser: parser,
	}
}

func (s *SubscribeAddressHandler) Handle(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx = c.Request.Context()
		req := &db_model.SubscribeAddressRequest{}
		if err := c.BindJSON(req); err != nil {
			c.JSON(http.StatusBadRequest, &db_model.HttpErrorResponse{
				Code:    consts.UnmarshalError,
				Message: "unable to unmarshal request to subscribeAddressRequest",
			})
			return
		}
		success := s.parser.Subscribe(ctx, req.Address)
		c.JSON(http.StatusOK, &db_model.SubscribeAddressResponse{
			SubscribeSuccess: success,
		})
	}
}
