package main

import (
	"fmt"
	"sqlgo/config"
	"sqlgo/users"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		fmt.Println("Ada Masalah", err.Error())
		return
	}

	err = db.AutoMigrate(&users.User{})
	if err != nil {
		fmt.Println("Ada Masalah", err.Error())
		return
	}

	fmt.Println("Tabel berhasil dibuat")
}
