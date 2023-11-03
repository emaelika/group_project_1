package model

import (
    "time"
)

type Transaction struct {
    TransactionID uint `gorm:"primaryKey"`
    UserID        uint `gorm:"foreignKey"`
    CustomerID    uint `gorm:"foreignKey"`
    ProductID  uint
    Quantity int
    Price       float64
    Total     float64
    CreatedAt time.Time
    UpdatedAt time.Time
}