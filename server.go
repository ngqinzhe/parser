package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/ngqinzhe/parser/clients/geth"
	"github.com/ngqinzhe/parser/handler"
	"github.com/ngqinzhe/parser/service"
)

type server struct {
	router  *gin.Engine
	gethCli geth.Client
}

func NewServer(gethCli geth.Client) *server {
	return &server{
		router:  gin.Default(),
		gethCli: gethCli,
	}
}

func (s *server) InitRoutes(ctx context.Context) {
	parser := service.NewEthParser(s.gethCli)

	s.router.POST("/subscribe", handler.NewSubscribeAddressHandler(parser).Handle(ctx))
	s.router.POST("/getTransactions", handler.NewGetTransactionsHandler(parser).Handle(ctx))

	s.router.Run("localhost:3000")
}
