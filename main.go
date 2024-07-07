package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/ngqinzhe/parser/clients/geth"
)

func main() {
	ctx := context.Background()
	// init logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	gethCli := geth.NewGethClient()
	s := NewServer(gethCli)
	s.InitRoutes(ctx)
}
