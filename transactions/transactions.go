package transactions

import (
	"time"
)

type Transaction struct {
	TransactionID uint `gorm:"primaryKey"`
	UserID        uint
	CustomerID    uint
	ProductID uint
	Quantity int
	date time.Time
	Total         float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}