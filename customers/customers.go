package customers

import (
	"sqlgo/model"

	"gorm.io/gorm"
)

type CustomersSystem struct {
	DB *gorm.DB
}

func (cs *CustomersSystem) ListUsers() ([]model.User, error) {
	var result = make([]model.User, 0)
	var qry = cs.DB.Table("customers").Find(&result)
	var err = qry.Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
