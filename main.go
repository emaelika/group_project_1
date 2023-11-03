package main

import (
	"fmt"
	"sqlgo/auth"
	"sqlgo/config"
	"sqlgo/customers"
	"sqlgo/model"
	"sqlgo/products"
	"sqlgo/receipts"
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
	var rs = receipts.ReceiptsSystem{DB: auth.DB}
	fmt.Println("\nSelamat Datang Di Tokoku App!")
	menuUtama(&auth, &ps, &us, &cs, &rs)

}

func menuUtama(auth *auth.AuthSystem, ps *products.ProductssSystem, us *users.UsersSystem, cs *customers.CustomersSystem, rs *receipts.ReceiptsSystem) {

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
					menuAdmin(auth, ps, us, cs, rs, user)
				} else if user.Role == "pegawai" {
					menuPegawai(auth, ps, cs, rs, user)
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

func menuAdmin(auth *auth.AuthSystem, ps *products.ProductssSystem, us *users.UsersSystem, cs *customers.CustomersSystem, rs *receipts.ReceiptsSystem, user model.User) {
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
		fmt.Println("(10) :> Edit Customer")
		fmt.Println("(11) :> Hapus Customer")
		fmt.Println("(12) :> Buat Nota Transaksi")
		fmt.Println("(13) :> Lihat Daftar Nota Transaksi")
		fmt.Println("(14) :> Hapus Nota Transaksi")
		fmt.Println("(15) :> Lihat Daftar Transaksi (COMING SOON)")
		fmt.Println("(16) :> Hapus Transaksi (COMING SOON)")
		fmt.Println("(99) :> Logout")
		fmt.Print("Masukkan Pilihan : ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			us.AddPegawai()
		case 2:
			us.ViewPegawai()
		case 3:
			us.DeletePegawai()
		case 4:
			ps.AddProduct(user)
		case 5:
			ps.ViewProduct()
		case 6:
			ps.UpdateProduct()
		case 7:
			ps.DeleteProduct()
		case 8:
			cs.AddCustomer()
		case 9:
			cs.ViewCustomers()
		case 10:
			cs.EditCustomer()
		case 11:
			cs.DeleteCustomer()
		case 12:
			rs.CreateReceipt(cs, ps, user)
		case 13:
			rs.ViewReceipt()
		case 14:
			rs.DeleteReceipt()

		case 15:
			fmt.Println("COMING SOON")
		case 16:
			fmt.Println("COMING SOON")
		case 99:
			fmt.Println("\nAdmin logout")
			return
		default:
			fmt.Println("\nPilihan tidak valid. Silakan pilih menu yang benar.")
		}
	}
}

func menuPegawai(auth *auth.AuthSystem, ps *products.ProductssSystem, cs *customers.CustomersSystem, rs *receipts.ReceiptsSystem, user model.User) {
	for {

		fmt.Println("\n    === MENU PEGAWAI  ===")
		fmt.Println("(1) :> Tambah Barang")
		fmt.Println("(2) :> Edit Barang")
		fmt.Println("(3) :> Lihat Daftar Barang")
		fmt.Println("(4) :> Tambah Customer")
		fmt.Println("(5) :> Lihat Daftar Customer")
		fmt.Println("(6) :> Buat Nota Transaksi")
		fmt.Println("(7) :> Lihat Daftar Nota Transaksi")
		fmt.Println("(8) :> Lihat Daftar Transaksi")
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
			cs.ViewCustomers()
		case 6:
			rs.CreateReceipt(cs, ps, user)
		case 7:
			rs.ViewReceipt()
		case 8:
			fmt.Println("COMING SOON")
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
