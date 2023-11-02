package model

import (
	"time"
)

type Product struct {
	ProductID   uint `gorm:"primaryKey"`
	UserID uint
	Username      string
	ProductName string
	Description string
	Stok        int
	Price       float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
