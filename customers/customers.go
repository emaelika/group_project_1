package customers

import (
	"time"
)

type Customer struct {
	CustomerID uint `gorm:"primaryKey"`
	CustomerName       string
	Address    string
	Phone      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}