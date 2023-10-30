package users

import (
	"fmt"
	"sqlgo/model"
	"gorm.io/gorm"
)

type UsersSystem struct {
	DB *gorm.DB
}

func (us *UsersSystem) MenuAdmin() {
	for {
		fmt.Println("\n    === Menu Admin  ===")
		fmt.Println("(1) :> Tambah Pegawai")
		fmt.Println("(2) :> Hapus Data Barang")
		fmt.Println("(3) :> Lihat Daftar Barang")
		fmt.Println("(0) :> Logout")
		fmt.Print("Masukkan pilihan : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			us.tambahPegawai()
		// case 2:
		// 	us.hapusDataBarang()
		// case 3:
		// 	us.tampilkanDaftarBarang()
		case 0:
			fmt.Println("\nAdmin logout")
			return
		default:
			fmt.Println("\nPilihan tidak valid. Silakan pilih menu yang benar.")
		}
	}
}


func (us *UsersSystem) tambahPegawai() {
	var newUser model.User

	fmt.Print("Masukkan username pegawai baru: ")
	fmt.Scanln(&newUser.Username)

	fmt.Print("Masukkan password pegawai baru: ")
	fmt.Scanln(&newUser.Password)

	newUser.Role = "pegawai"

	result := us.DB.Create(&newUser)

	if result.Error != nil {
		fmt.Println("Error saat menambahkan pegawai:", result.Error)
		return
	}

	fmt.Println("Pegawai baru berhasil ditambahkan!")
}


func (us *UsersSystem) MenuPegawai(string) {
	for {
        fmt.Println("\nMenu Pegawai:")
        fmt.Println("1. Tambah Transaksi")
        fmt.Println("2. Lihat Daftar Barang")
        fmt.Println("0. Logout")
        fmt.Print("Masukkan pilihan: ")

        var choice int
        fmt.Scanln(&choice)

        // switch choice {
        // case 1:
        //     us.tambahTransaksi(pegawai)
        // case 2:
        //     us.tampilkanDaftarBarang()
        // case 0:
        //     fmt.Printf("Pegawai '%s' logout.\n", pegawai)
        //     return
        // default:
        //     fmt.Println("Pilihan tidak valid. Silakan pilih menu yang benar.")
        // }
    }
}