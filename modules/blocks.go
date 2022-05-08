package modules

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	Models "github.com/jahankohan/geth_mempool/models"
)

type ClientHandler struct {
	*ethclient.Client
}

func (client ClientHandler) GetLatestBlock() *Models.Block {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	header, _ := client.HeaderByNumber(context.Background(), nil)
	blockNumber := big.NewInt(header.Number.Int64())
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	_block := &Models.Block{
		BlockNumber: 		block.Number().Int64(),
		Timestamp: 			block.Time(),
		Difficulty: 		block.Difficulty().Uint64(),
		Hash: 				block.Hash().String(),
		TransactionsCount:	len(block.Transactions()),
		Transactions: 		[]Models.Transaction{},
	}

	for _, tx := range block.Transactions() {
		_block.Transactions = append(_block.Transactions, Models.Transaction{
			Hash: 		tx.Hash().String(),
			Value: 		tx.Value().String(),
			Gas: 		tx.Gas(),
			GasPrice:	tx.GasPrice().Uint64(),
			Nonce: 		tx.Nonce(),
			To: 		tx.To().String(),
		})
	}
	return _block
}

func (client ClientHandler) GetTxByHash(hash common.Hash) *Models.Transaction {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	
	tx, pending, err := client.TransactionByHash(context.Background(), hash)
	if err != nil {
		fmt.Println(err)
	}

	return &Models.Transaction{
		Hash:		tx.Hash().String(),
		Value:		tx.Value().String(), 
		Gas:		tx.Gas(),
		GasPrice:	tx.GasPrice().Uint64(),
		Nonce:		tx.Nonce(),
		To:			tx.To().String(),
		Pending:	pending,
	}
}