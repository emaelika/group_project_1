package products

import (
	"gorm.io/gorm"
)

type ProductsSystem struct {
	DB *gorm.DB
}
