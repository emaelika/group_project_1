package model

import (
	"time"
)

type Product struct {
	ProductID   uint `gorm:"primaryKey"`
	UserID      uint
	ProductName string
	Description string
	Stok        int
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
