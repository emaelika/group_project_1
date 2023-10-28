package main

import (
	"fmt"
	"sqlgo/config"
	"sqlgo/customers"
	"sqlgo/products"
	"sqlgo/receipts"
	"sqlgo/transactions"
	"sqlgo/users"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		fmt.Println("Ada Masalah", err.Error())
		return
	}

	err = db.AutoMigrate(&users.User{}, &customers.Customer{}, &products.Product{}, &receipts.Receipt{}, &transactions.Transaction{})
	if err != nil {
		fmt.Println("Ada Masalah", err.Error())
		return
	}

	fmt.Println("Tabel berhasil dibuat")
}
