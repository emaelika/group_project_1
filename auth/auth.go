package auth

import (
	"fmt"
	"sqlgo/model"
	"sqlgo/users"

	"gorm.io/gorm"
)

type AuthSystem struct {
	DB *gorm.DB
}



func (as *AuthSystem) SetupAdmin() {
	var admin model.User
	err := as.DB.Where("username = ? AND role = ?", "admin", "admin").First(&admin).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Admin tidak ditemukan, buat admin baru
			admin = model.User{
				Username: "admin",
				Password: "admin",
				Role:     "admin",
			}
			result := as.DB.Create(&admin)
			if result.Error != nil {
				fmt.Println("Error saat membuat admin:", result.Error)
				return
			}
			fmt.Println("Admin berhasil dibuat!")
		} else {
			// Error database lainnya
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("")
	}
}

func (as *AuthSystem) LoginAdmin() (model.User, bool) {
	var currentUser = new(model.User)

	fmt.Print("Masukkan username: ")
	fmt.Scanln(&currentUser.Username)

	fmt.Print("Masukkan Password: ")
	var password string
	fmt.Scanln(&password)

	usersSystem := users.UsersSystem{}

	qry := as.DB.Where("username = ?", currentUser.Username).Take(currentUser)

	err := qry.Error

	if err != nil {
		fmt.Println("\nAdmin Gagal Masuk. Username atau Password Salah !!")
		return model.User{}, false
	}

	if currentUser.Password != password {
		fmt.Println("\nAdmin Gagal Masuk. Username atau Password Salah.")
		return model.User{}, false
	}

	fmt.Println("\nAdmin Berhasil Masuk!")
	usersSystem.MenuAdmin()

	return *currentUser, true
}


func (as *AuthSystem) LoginPegawai() (model.User, bool) {
	var currentUser = new(model.User)

	fmt.Print("Masukkan username pegawai: ")
    var username string
    fmt.Scanln(&username)

	usersSystem := users.UsersSystem{}
	
	
	if as.isPegawaiValid(username) {
		fmt.Printf("Pegawai '%s' login berhasil!\n", username)
		usersSystem.MenuPegawai(username)
	} else {
		fmt.Println("Login pegawai gagal. Username pegawai tidak valid.")
	}
	return *currentUser, true

}

func (as *AuthSystem) isPegawaiValid(username string) bool {
	var user model.User
	err := as.DB.Where("username = ? AND role = ?", username, "pegawai").First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// User tidak ditemukan
			return false
		}
		// Error database lainnya
		fmt.Println("Error:", err)
		return false
	}
	// User ditemukan
	return true
}
