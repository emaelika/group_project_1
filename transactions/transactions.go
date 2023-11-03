package transactions

import (
	"sqlgo/model"

	"gorm.io/gorm"
)

type TransactionsSystem struct {
	DB *gorm.DB
}

func (ts *TransactionsSystem) AddTransaction(user model.User, produk model.Product, nota model.Receipt) {
	var newTrans model.Transaction
	newTrans.TransactionID = nota.TransactionID
	newTrans.CustomerID = nota.CustomerID

}
