
package users

import (
	"time"
)

type User struct {
	UserID    uint `gorm:"primaryKey"`
	Username  string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}