package main

import (
	"fmt"
	"sqlgo/auth"
	"sqlgo/config"
	"sqlgo/model"
	"sqlgo/products"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		fmt.Println("Ada Masalah", err.Error())
		return
	}

	err = db.AutoMigrate(&model.User{}, &model.Customer{}, &model.Product{}, &model.Receipt{}, &model.Transaction{})
	if err != nil {
		fmt.Println("Ada Masalah", err.Error())
		return
	}
	var auth = auth.AuthSystem{DB: db}
	auth.SetupAdmin()
	var ps = products.ProductssSystem{DB: auth.DB}
	for {
		fmt.Println("\nSelamat Datang Di Tokoku App!")
		menuUtama(&auth, &ps)
	}
}

func menuUtama(auth *auth.AuthSystem, ps *products.ProductssSystem) {
	
	for {
		fmt.Println("\n     === MENU ===")
        fmt.Println("(1) :> Login")
        fmt.Println("(0) :> Exit")
        fmt.Print("Masukkan Pilihan : ")
		
        var inputMenu int
        fmt.Scanln(&inputMenu)
		
        switch inputMenu {
		case 1:
            user, loggedIn := auth.Login()
            if loggedIn {
				if user.Role == "admin" {
					menuAdmin(auth, ps, user)
                } else if user.Role == "pegawai" {
                    menuPegawai(auth, ps, user)
                }
            }
        case 0:
            fmt.Println("\nTerima kasih!")
            return
        default:
            fmt.Println("\nPilihan tidak valid. Silakan pilih menu yang benar.")
        }
    }
}

func menuAdmin(auth *auth.AuthSystem, ps *products.ProductssSystem, user model.User) {
	for {
        		fmt.Println("\n    === MENU ADMIN  ===")
		fmt.Println("(1) :> Tambah Pegawai")
		fmt.Println("(2) :> Tambah Barang")
		fmt.Println("(3) :> Edit Barang")
		fmt.Println("(4) :> Hapus Barang")
		fmt.Println("(5) :> Lihat Daftar Barang")
		fmt.Println("(6) :> Lihat Daftar Pegawai")
		fmt.Println("(0) :> Logout")
		fmt.Print("Masukkan Pilihan : ")

        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            auth.AddPegawai()
		case 2:
			ps.AddProduct(user)
        case 0:
            fmt.Println("\nAdmin logout")
            return
        default:
            fmt.Println("\nPilihan tidak valid. Silakan pilih menu yang benar.")
        }
    }
}

func menuPegawai(auth *auth.AuthSystem, ps *products.ProductssSystem, user model.User) {
	for {
        fmt.Println("\n    === MENU PEGAWAI  ===")
        fmt.Println("(1) :> Tambah Barang")
        fmt.Println("(2) :> Edit Barang")
        fmt.Println("(3) :> Tambah Customer")
        fmt.Println("(4) :> Buat Transaksi")
        fmt.Println("(5) :> Lihat Daftar Barang")
        fmt.Println("(0) :> Logout")
        fmt.Print("Masukkan Pilihan : ")

        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            ps.AddProduct(user)
        // case 2:
        //     auth.EditBarang()
        // case 3:
        //     auth.TambahCustomer()
        // case 4:
        //     auth.BuatTransaksi()
        // case 5:
        //     auth.LihatDaftarBarang()
        case 0:
            fmt.Println("\nPegawai logout")
            return
        default:
            fmt.Println("\nPilihan tidak valid. Silakan pilih menu yang benar.")
        }
    }
}
