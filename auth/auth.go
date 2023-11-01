package auth

import (
	"fmt"
	"sqlgo/model"

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
			admin = model.User{
				Username: "admin",
				Password: "admin",
				Role:     "admin",
			}
			result := as.DB.Create(&admin)
			if result.Error != nil {
				fmt.Println("Error saat membuat Akun admin:", result.Error)
				return
			}
			fmt.Println("\nAkun Admin berhasil dibuat !")
		} else {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("")
	}
}

func (as *AuthSystem) Login() (model.User, bool) {
	var currentUser = new(model.User)

	fmt.Print("Masukkan Username : ")
	fmt.Scanln(&currentUser.Username)

	fmt.Print("Masukkan Password : ")
	var password string
	fmt.Scanln(&password)

	qry := as.DB.Where("username = ?", currentUser.Username).Take(currentUser)

	err := qry.Error

	if err != nil {
		fmt.Println("\nLogin Gagal Masuk. Username atau Password Salah !!")
		return model.User{}, false
	}

	if currentUser.Password != password {
		fmt.Println("\nLogin Gagal Masuk. Username atau Password Salah.")
		return model.User{}, false
	}

	fmt.Println("\nLogin Berhasil")
	if currentUser.Role == "admin" {
		fmt.Println("\nAdmin Berhasil Masuk!")
	} else if currentUser.Role == "pegawai" {
        fmt.Printf("\nBerhasil Masuk!\nPegawai : %s\n", currentUser.Username)
    } else {
        fmt.Println("\nPeran Tidak Valid")
    }
    return *currentUser, true
}

func (as *AuthSystem) AddPegawai() {
	var pegawai model.User

	fmt.Print("Masukkan Username Pegawai Baru : ")
    fmt.Scanln(&pegawai.Username)

    fmt.Print("Masukkan Password Pegawai Baru : ")
    fmt.Scanln(&pegawai.Password)

    pegawai.Role = "pegawai"

	result := as.DB.Create(&pegawai)

	if result.Error != nil {
        fmt.Println("Error saat menambahkan pegawai:", result.Error)
        return
    }
	
    fmt.Printf("\nPegawai '%s' berhasil ditambahkan!\n", pegawai.Username)
}
