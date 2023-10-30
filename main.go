package main

import (
	"fmt"
	"sqlgo/auth"
	"sqlgo/config"
	"sqlgo/model"
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
	for {
		fmt.Println("\nSelamat Datang Di Tokoku App!")

		for {
		fmt.Println("\n     === MENU ===")
		fmt.Println("(1) :> Login Admin")
		fmt.Println("(2) :> Login Pegawai")
		fmt.Println("(0) :> Exit")
		fmt.Print("Masukkan input : ")

		var inputMenu int
		fmt.Scanln(&inputMenu)

		switch inputMenu {
		case 1:
			auth.LoginAdmin()
		case 2:
			auth.LoginPegawai()
		case 0:
			fmt.Println("\nTerima kasih!")
			return
		default:
			fmt.Println("\nPilihan tidak valid. Silakan pilih menu yang benar.")
		}
	}
}
}
