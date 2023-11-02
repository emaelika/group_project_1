package model

import (
	"time"
)

type Transaction struct {
	TransactionID uint `gorm:"primaryKey"`
	UserID        uint `gorm:"foreignKey"`
	CustomerID    uint `gorm:"foreignKey"`
	ProductID     uint `gorm:"foreignKey"`
	Quantity      int

	Total     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
