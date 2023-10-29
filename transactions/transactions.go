package transactions

import (
	"gorm.io/gorm"
)

type TransactionsSystem struct {
	DB *gorm.DB
}
