package model

type TransactionDTO struct {
	Event    int `json:"event"`
	Account  int `json:"account"`
	Quantity int `json:"quantity"`
}

func (transactionDTO TransactionDTO) DTOtoTransaction() Transaction {
	return Transaction{Event: transactionDTO.Event, Account: transactionDTO.Account, Quantity: transactionDTO.Quantity}
}

type Transaction struct {
	Id       int `json:"id"`
	Event    int `json:"event"`
	Account  int `json:"account"`
	Quantity int `json:"quantity"`
}

type TransactionDB struct {
	Id      int `db:"id"`
	Event   int `db:"event"`
	Account int `db:"account"`
}

func (transactionDB TransactionDB) DBtoTransaction() Transaction {
	return Transaction{Id: transactionDB.Id, Event: transactionDB.Event, Account: transactionDB.Account}
}
