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
		fmt.Println("Welcome!")
		fmt.Println("1. Login")
		
		fmt.Println("9. Exit")
		fmt.Print("Masukkan input: ")
		fmt.Scanln(&inputMenu)
		switch inputMenu {
		case 1:
			// kode login
			if role == admin {
				// interface admin
				1. user {
					1. show 
					2. add
					3. edit
					4. delete
				}
				2. products {
					1. show 
					2. add
					3. edit
					4. delete
				}
				3. transactions
					1. show 
					2. add
					3. edit
					4. delete
				4. customers {
					1. show 
					2. add
					3. edit
					4. delete
				}
				5. receipts {
					1. show 
					2. add
					3. edit
					4. delete
				}
			} else if role == user {
				
				2. products {
					1. show 
					2. add
					3. edit
					
				}
		
				4. customers {
					1. show 
					2. add
					
				}
				5. receipts {
					1. show 
					2. add
					
				}
			}
		case 9:
			// exit
			
		}
		
	}
	fmt.Println("Thank you....")
}
