package customers

import (
<<<<<<< HEAD
	"sqlgo/model"

	"gorm.io/gorm"
)

type CustomersSystem struct {
	DB *gorm.DB
}

func (cs *CustomersSystem) ListCustomers() ([]model.Customer, error) {
	var result = make([]model.Customer, 0)
	var qry = cs.DB.Table("customers").Find(&result)
	var err = qry.Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
=======
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
>>>>>>> 542e56a07d9e9108743285783f53c27e4af0dc69
