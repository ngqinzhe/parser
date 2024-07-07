package service

import (
	"context"
	"fmt"
	"log/slog"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ngqinzhe/parser/clients/geth"
)

type Parser interface {
	// last parsed block
	GetCurrentBlock(ctx context.Context) uint64
	// add address to observer
	Subscribe(ctx context.Context, address string) bool
	// list of inbound or outbound transactions for an address
	GetTransactions(ctx context.Context, address string) []*types.Transaction
}

type EthParser struct {
	gethClient       geth.Client
	addressWhitelist map[string]interface{}
	lock             sync.RWMutex
}

func NewEthParser(gethCli geth.Client) Parser {
	return &EthParser{
		gethClient:       gethCli,
		addressWhitelist: make(map[string]interface{}),
		lock:             sync.RWMutex{},
	}
}

func (e *EthParser) GetCurrentBlock(ctx context.Context) uint64 {
	block, err := e.gethClient.GetBlockNumber(ctx)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("[EthParser][GetCurrentBlock] unable to get current block, err: %v", err))
		return 0
	}
	return block
}

func (e *EthParser) Subscribe(ctx context.Context, address string) bool {
	e.lock.Lock()
	defer e.lock.Unlock()
	if _, exist := e.addressWhitelist[address]; exist {
		slog.InfoContext(ctx, "[EthParser][Subscribe] address already subscribed")
		return false
	}
	e.addressWhitelist[address] = struct{}{}
	return true
}

func (e *EthParser) GetTransactions(ctx context.Context, address string) []*types.Transaction {
	e.lock.RLock()
	_, subscribed := e.addressWhitelist[address]
	e.lock.RUnlock()

	if !subscribed {
		slog.InfoContext(ctx, fmt.Sprintf("[EthParser][GetTransactions] address: %s is not subscribed, will not listen to transactions", address))
		return nil
	}

	currentBlockNum := e.GetCurrentBlock(ctx)
	if currentBlockNum == 0 {
		slog.ErrorContext(ctx, "[EthParser][GetTransactions] unable to get current block")
		return nil
	}
	slog.InfoContext(ctx, fmt.Sprintf("currentBlockNum: %d", currentBlockNum))
	transactions, err := e.gethClient.GetTransactionsByAddress(ctx, address, currentBlockNum)
	if err != nil {
		slog.ErrorContext(ctx, fmt.Sprintf("[EthParser][GetTransactions] unable to get transactions, err: %v", err))
		return nil
	}
	return transactions
}
