package main

import (
	"fmt"
	"sqlgo/auth"
	"sqlgo/config"
	"sqlgo/customers"
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
}
