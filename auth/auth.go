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
	var hp string
	fmt.Print("Masukkan Hp")
	fmt.Scanln(&hp)

	qry := as.DB.Table("pelanggan").Where("hp = ?", hp).Take(currentUser)

	err := qry.Error

	if err != nil {
		fmt.Println("login process error:", err.Error())
		return model.User{}, false
	}

	return *currentUser, true
}

func (as *AuthSystem) Register() (model.User, bool) {
	var newUser = new(model.User)
	fmt.Print("Masukkan Hp")
	fmt.Scanln(&newUser.HP)
	fmt.Print("Masukkan Nama")
	fmt.Scanln(&newUser.Nama)
	fmt.Print("Masukkan Tanggal Lahir")
	fmt.Scanln(&newUser.TanggalLahir)

	err := as.DB.Table("pelanggan").Create(newUser).Error
	if err != nil {
		fmt.Println("input error:", err.Error())
		return model.User{}, false
	}

	return *newUser, true
}
