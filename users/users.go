package users

import (
	"sqlgo/model"

	"gorm.io/gorm"
)

type UsersSystem struct {
	DB *gorm.DB
}

func (us *UsersSystem) ListUsers() ([]model.User, error) {
	var result = make([]model.User, 0)
	var qry = us.DB.Table("pelanggan").Find(&result)
	var err = qry.Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
