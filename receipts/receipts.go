package receipts

import (
	"gorm.io/gorm"
)

type ReceiptsSystem struct {
	DB *gorm.DB
}

func (rs *ReceiptsSystem) CreateReceipt()
