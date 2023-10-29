package main

import (
	"fmt"
	"sqlgo/config"
	"sqlgo/customers"
<<<<<<< HEAD
	"sqlgo/model"
=======
	"sqlgo/products"
	"sqlgo/receipts"
	"sqlgo/transactions"
	"sqlgo/users"
>>>>>>> 542e56a07d9e9108743285783f53c27e4af0dc69
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		fmt.Println("Ada Masalah", err.Error())
		return
	}

<<<<<<< HEAD
	err = db.AutoMigrate(&model.User{}, &model.Customer{}, &model.Product{}, &model.Receipt{}, &model.Transaction{})
	if err != nil {
		fmt.Println("Ada Masalah", err.Error())
		return
	}
	var auth = auth.AuthSystem{DB: db}
	var customers = customers.CustomersSystem{DB: db}
	for {
		var inputMenu int
		fmt.Println("Selamat datang di Tokoku App!\nMenu:\n1. Log in\n99. Exit\nMasukkan input: ")
		fmt.Scanln(&inputMenu)
		switch inputMenu {
		case 1:
			var menuLogin int
			result, permit := auth.Login()
			if permit && result.Username == "ADMIN" {
				fmt.Println("\nSelamat datang ", result.Username)
				for permit {
					fmt.Println("\n1.")
					fmt.Println("2. Customer")
					fmt.Println("3. ")
					fmt.Println("4. ")
					fmt.Println("0. Logout")
					fmt.Println("99. Exit")
					fmt.Print("Masukkan pilihan: ")
					fmt.Scanln(&menuLogin)
					switch menuLogin {
					case 2:
						for {
							fmt.Println("\n1. Customers List")
							fmt.Println("2. ")
							fmt.Println("3. ")
							fmt.Println("4. ")
							fmt.Println("0. Logout")
							fmt.Println("99. Exit")
							fmt.Print("Masukkan pilihan: ")
							fmt.Scanln(&menuLogin)
							switch menuLogin {
							case 1:
								cust, err := customers.ListCustomers()
								if err != nil {
									fmt.Println("Error getting customer:", err.Error())
									return
								}

								for _, cust := range cust {
									fmt.Println(cust.CustomerID, cust.CustomerName, cust.Address)
								}
							case 99:
								fmt.Println("\nTerima kasih!")
								return

							}
						}
					case 99:
						fmt.Println("\nTerima kasih!")
						return
					}
				}
			}

		case 99:
			fmt.Println("\nTerima kasih!")
			return
		}
	}
=======
	err = db.AutoMigrate(&users.User{}, &customers.Customer{}, &products.Product{}, &receipts.Receipt{}, &transactions.Transaction{})
	if err != nil {
		fmt.Println("Ada Masalah", err.Error())
		return
	}

	fmt.Println("Tabel berhasil dibuat")
>>>>>>> 542e56a07d9e9108743285783f53c27e4af0dc69
}
