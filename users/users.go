package users

import (
	"fmt"
	"sqlgo/model"

	"gorm.io/gorm"
)

type UsersSystem struct {
	DB *gorm.DB
}

func (us *UsersSystem) AddPegawai() {
    for {
        var pegawai model.User
        var choice int

        fmt.Print("Masukkan Username Pegawai Baru : ")
        fmt.Scanln(&pegawai.Username)

        fmt.Print("Masukkan Password Pegawai Baru : ")
        fmt.Scanln(&pegawai.Password)

        pegawai.Role = "pegawai"

        result := us.DB.Create(&pegawai)

        if result.Error != nil {
            fmt.Println("Error saat menambahkan pegawai:", result.Error)
            return
        }

        fmt.Printf("\nPegawai '%s' berhasil ditambahkan!\n", pegawai.Username)

        fmt.Print("\n\t(1) :> Tambah Pegawai Lagi \n\t(9) :> Kembali \nMasukkan Pilihan : ")
        fmt.Scanln(&choice)
        if choice == 9 {
            return
        } else if choice != 1 {
            fmt.Println("Coba lagi")
            continue
        }
    }
}


func (us *UsersSystem) ViewPegawai() {
	var choice int
	var pegawai []model.User
	result := us.DB.Find(&pegawai)
	if result.Error != nil {
		fmt.Println("Error saat mengambil data pegawai:", result.Error)
		return
	}
	fmt.Println("")
	for _, user := range pegawai {
		if user.Role != "admin" {
			fmt.Printf("\nNama : %s\nStatus : %s\n", user.Username, user.Role)
		}
	}
	fmt.Print("\n\t(9) :> Kembali\nMasukkan Pilihan : ")
	fmt.Scanln(&choice)
	if choice == 9 {
		return
	} else {
		fmt.Println("Coba lagi")
	}
}


func (us *UsersSystem) DeletePegawai() {
    for {
        var pegawai model.User
        var username string
        var choice int

        fmt.Print("Masukkan Username Pegawai yang akan dihapus : ")
        fmt.Scanln(&username)

        result := us.DB.Where("username = ?", username).First(&pegawai)
        if result.Error != nil {
            fmt.Println("\nUsername pegawai tidak ditemukan.")
            fmt.Print("\n\t(1) :> Hapus Pegawai \n\t(9) :> Kembali \nMasukkan Pilihan : ")
            fmt.Scanln(&choice)
            if choice == 9 {
                return
            } else if choice != 1 {
                fmt.Println("Coba lagi")
                continue
            }
        } else {
            result = us.DB.Delete(&pegawai)
            if result.Error != nil {
                fmt.Println("Error saat menghapus pegawai:", result.Error)
                return
            }

            fmt.Printf("\nPegawai '%s' berhasil dihapus!\n", username)

            fmt.Print("\n\t(1) :> Hapus Pegawai Lagi \n\t(9) :> Kembali \nMasukkan Pilihan : ")
            fmt.Scanln(&choice)
            if choice == 9 {
                return
            } else if choice != 1 {
                fmt.Println("Coba lagi")
                continue
            }
        }
    }
}
