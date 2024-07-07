package geth

import (
	"context"
	"fmt"
	"log/slog"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Client interface {
	GetBlockNumber(ctx context.Context) (uint64, error)
	GetTransactionsByAddress(ctx context.Context, address string, blockNum uint64) ([]*types.Transaction, error)
}

type ethClientImpl struct {
	cli *ethclient.Client
}

func NewGethClient() Client {
	cli, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		panic(err)
	}
	return &ethClientImpl{
		cli: cli,
	}
}

func (e *ethClientImpl) GetBlockNumber(ctx context.Context) (uint64, error) {
	blockNum, err := e.cli.BlockNumber(ctx)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("[ethClientImpl][GetBlockNumber] failed, err:%v", err))
		return 0, err
	}
	return blockNum, nil
}

// GetTransactionsByAddress will get all the transactions from the given address for the past 1000 blocks (limitation)
func (e *ethClientImpl) GetTransactionsByAddress(ctx context.Context, address string, currentBlockNum uint64) ([]*types.Transaction, error) {
	startBlockNum := currentBlockNum - 1000
	endBlockNum := currentBlockNum

	slog.InfoContext(ctx, fmt.Sprintf("[ethClientImpl][GetTransactionsByAddress] startBlockNum: %d, endBlockNum: %d", startBlockNum, endBlockNum))
	var allTransactions []*types.Transaction
	for i := startBlockNum; i <= endBlockNum; i++ {
		blockNum := big.NewInt(int64(i))
		block, err := e.cli.BlockByNumber(ctx, blockNum)
		if err != nil {
			slog.WarnContext(ctx, fmt.Sprintf("[ethClientImpl][GetTransactionsByAddress] unable to get block: %s", blockNum.String()))
			continue
		}
		allTransactions = append(allTransactions, block.Transactions()...)
	}

	return allTransactions, nil
}
