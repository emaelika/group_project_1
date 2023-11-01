package main

import (
	"fmt"
	"sqlgo/auth"
	"sqlgo/config"
	"sqlgo/customers"
	"sqlgo/model"
	"sqlgo/products"
	"sqlgo/users"
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
	var us = users.UsersSystem{DB: auth.DB}
	var cs = customers.CustomersSystem{DB: auth.DB}
	fmt.Println("\nSelamat Datang Di Tokoku App!")
	menuUtama(&auth, &ps, &us, &cs)

}

func menuUtama(auth *auth.AuthSystem, ps *products.ProductssSystem, us *users.UsersSystem, cs *customers.CustomersSystem) {

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
					menuAdmin(auth, ps, us, cs, user)
				} else if user.Role == "pegawai" {
					menuPegawai(auth, ps, cs, user)
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

func menuAdmin(auth *auth.AuthSystem, ps *products.ProductssSystem, us *users.UsersSystem, cs *customers.CustomersSystem, user model.User) {
	for {

		fmt.Println("\n    === MENU ADMIN  ===")
		fmt.Println("(1)  :> Tambah Pegawai")
		fmt.Println("(2)  :> Lihat Daftar Pegawai")
		fmt.Println("(3)  :> Hapus Pegawai")
		fmt.Println("(4)  :> Tambah Barang")
		fmt.Println("(5)  :> Lihat Daftar Barang")
		fmt.Println("(6)  :> Edit Barang")
		fmt.Println("(7)  :> Hapus Barang")
		fmt.Println("(8)  :> Tambah Customer")
		fmt.Println("(9)  :> Lihat Daftar Customer")
		fmt.Println("(10) :> Hapus Customer")
		fmt.Println("(11) :> Buat Nota Transaksi")
		fmt.Println("(12) :> Lihat Daftar Transaksi")
		fmt.Println("(13) :> Hapus Transaksi")
		fmt.Println("(99) :> Logout")
		fmt.Print("Masukkan Pilihan : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			auth.AddPegawai()
        case 2:
			pegawai, _ := us.ListUsers()
			fmt.Println("")
			for _, user := range pegawai {
				if user.Role != "admin" {
					fmt.Println(user.UserID, user.Username, user.Role)
				}

			}
			fmt.Print("\n\t9. Back\nMasukkan input: ")
			fmt.Scanln(&choice)
			if choice == 9 {
				break
			} else {
				fmt.Println("Coba lagi")
			}
		case 3:
			ps.AddProduct(user)
        case 4:
            ps.ViewProduct()
        case 5:
            ps.UpdateProduct()
        case 6:
            ps.DeleteProduct()
		case 99:
			fmt.Println("\nAdmin logout")
			return
		default:
			fmt.Println("\nPilihan tidak valid. Silakan pilih menu yang benar.")
		}
	}
}

func menuPegawai(auth *auth.AuthSystem, ps *products.ProductssSystem, cs *customers.CustomersSystem, user model.User) {
	for {

		fmt.Println("\n    === MENU PEGAWAI  ===")
		fmt.Println("(1) :> Tambah Barang")
		fmt.Println("(2) :> Edit Barang")
		fmt.Println("(3) :> Lihat Daftar Barang")
		fmt.Println("(4) :> Tambah Customer")
		fmt.Println("(5) :> Lihat Daftar Customer")
		fmt.Println("(6) :> Buat Nota Transaksi")
		fmt.Println("(7) :> Lihat Daftar Transaksi")
		fmt.Println("(9) :> ")
		fmt.Println("(99) :> Logout")
		fmt.Print("Masukkan Pilihan : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			ps.AddProduct(user)
        case 2:
            ps.UpdateProduct()
        case 3:
            ps.ViewProduct()
		case 4:
			cs.AddCustomer()
		case 5:
			cs.ListCustomers()
		
		case 0:
			fmt.Println("")

		case 99:
			fmt.Println("\nPegawai logout")
			return
		default:
			fmt.Println("\nPilihan tidak valid. Silakan pilih menu yang benar.")
		}
	}
}
