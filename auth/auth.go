package auth

import (
	"fmt"
	"sqlgo/model"

	"gorm.io/gorm"
)

type AuthSystem struct {
	DB *gorm.DB
}

func (as *AuthSystem) Login() (model.User, bool) {
	// Input data -> hp
	// cari nomor HP di database
	// jika ditemukan berarti data valid
	// jika tidak ditemukan berarti data tidak valid

	var currentUser = new(model.User)

	fmt.Print("Masukkan username: ")
	fmt.Scanln(&currentUser.Username)
	fmt.Print("Masukkan Password: ")
	fmt.Scanln(&currentUser.Password)

	// qry := as.DB.Table("pelanggan").Where("hp = ?", hp).Take(currentUser)
	qry := as.DB.Where("username = ? AND password = ?", currentUser.Username, currentUser.Password).Take(currentUser)

	err := qry.Error

	if err != nil {
		fmt.Println("login process error:", err.Error())
		return model.User{}, false
	}

	return *currentUser, true
}
