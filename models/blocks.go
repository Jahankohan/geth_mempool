package models


type Block struct {
	BlockNumber			int64			`json:"blockNumber"`
	Timestamp			uint64			`json:"Timestamp"`
	Difficulty			uint64			`json:"Difficulty"`
	Hash				string			`json:"Hash"`
	TransactionsCount	int				`json:"TransactionsCount"`
	Transactions		[]Transaction	`json:"Transactions"`
}

type Transaction struct {
	Hash		string	`json:"hash"`
	Value		string	`json:"value"`
	Gas			uint64	`json:"gas"`
	GasPrice	uint64	`json:"gasPrice"`
	Nonce		uint64	`json:"nonce"`
	To			string	`json:"to"`
	Pending		bool	`json:"pending"`
}

type TransferEthRequest struct{
	PrivKey	string	`json:"privKey"`
	To		string	`json:"to"`
	Amount	int64	`json:"amount"`
}

type HashResponse struct {
	Hash	string	`json:"hash"`
}

type BalanceResponse struct {
	Address	string	`json:"address"`
	Balance	string	`json:"balance"`
	Symbol	string	`json:"symbol"`
	Units	string	`json:"units"`
}

type Error struct {
	Code	uint64	`json:"code"`
	Message	string	`json:"message"`
}