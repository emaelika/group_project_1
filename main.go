package main

import (
	"fmt"
	"sqlgo/auth"
	"sqlgo/config"
	"sqlgo/model"
	"sqlgo/users"
)

type Category struct {
	ID           int    `gorm:"primaryKey"`
	NamaCategory string `gorm:"type:varchar(255)"`
}

func main() {
	var inputMenu = 0
	cfg, err := config.InitDB()

	if err != nil {
		fmt.Println("Cannot start program, database issue", err.Error())
	}

	var authSystem = auth.AuthSystem{DB: cfg}
	var usersSystem = users.UsersSystem{DB: cfg}

	for {
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("9. Exit")
		fmt.Scanln(&inputMenu)
		if inputMenu == 9 {
			break
		} else if inputMenu == 2 {
			data, permit := authSystem.Register()
			if permit {
				fmt.Println(data)
			}
		} else if inputMenu == 1 {
			var inputInside int
			data, permit := authSystem.Login()
			fmt.Println("Hello", data.Nama)
			for permit {
				fmt.Println("1. Insert Buku")
				fmt.Println("9. Logout")
				fmt.Scanln(&inputInside)
				if inputInside == 9 {
					permit = false
					data = model.User{}
				} else if inputInside == 1 {
					list, err := usersSystem.ListUsers()

					if err != nil {
						fmt.Println("terjadi error", err.Error())

					} else {
						fmt.Println(list)
					}

				}
			}
		}
	}
	fmt.Println("Thank you....")
}
