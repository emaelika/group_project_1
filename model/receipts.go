package model

import (
    "time"
)

type Receipt struct {
	TransactionID uint `gorm:"primaryKey"`
    CustomerID    uint `gorm:"foreignKey"`
    ProductID     uint `gorm:"foreignKey"`
    CustomerName  string
    ProductName   string
    Total         float64
    Quantity      int
    Price     float64
    CreatedAt time.Time
    UpdatedAt time.Time
}