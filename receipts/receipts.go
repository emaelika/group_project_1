package receipts

import (
	"time"
)

type Receipt struct {
	NoteID        uint `gorm:"primaryKey"`
	CustomerID uint
	TransactionID uint
	ProductName string
	Product       string
	Total int
	Quantity      int
	CustomerName string
	date time.Time
	Price         float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}