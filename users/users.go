
package users

import (
	"time"
)

<<<<<<< HEAD
type UsersSystem struct {
	DB *gorm.DB
}

func (us *UsersSystem) ListUsers() ([]model.User, error) {
	var result = make([]model.User, 0)
	var qry = us.DB.Table("Users").Find(&result)
	var err = qry.Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
=======
type User struct {
	UserID    uint `gorm:"primaryKey"`
	Username  string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
>>>>>>> 542e56a07d9e9108743285783f53c27e4af0dc69
