package model

import (
	"time"
)

type Receipt struct {
	CustomerID    uint `gorm:"foreignKey"`
	TransactionID uint `gorm:"primaryKey"`
	ProductName   string

	Total        int
	Quantity     int
	CustomerName string

	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
