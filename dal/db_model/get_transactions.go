package db_model

import "time"

type GetTransactionsRequest struct {
	Address string `json:"address"`
}

type GetTransactionsResponse struct {
	Transactions []*EthTransaction `json:"transactions"`
}

type EthTransaction struct {
	Address    string    `json:"address"`
	IsSender   bool      `json:"is_sender"`
	Data       []byte    `json:"data"`
	Gas        uint64    `json:"gas"`
	Nonce      uint64    `json:"nonce"`
	CreateTime time.Time `json:"create_time"`
}
