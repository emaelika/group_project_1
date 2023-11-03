package customers

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sqlgo/model"

	"gorm.io/gorm"
)

type CustomersSystem struct {
	DB *gorm.DB
}

func (cs *CustomersSystem) AddCustomer() {
    for {
        var Customer model.Customer
        var choice int
        reader := bufio.NewReader(os.Stdin)

        fmt.Print("\nMasukkan Nama Customer : ")
        longString, _ := reader.ReadString('\n')
        Customer.CustomerName = longString

        var existingCustomer model.Customer
        if err := cs.DB.Where("customer_name = ?", Customer.CustomerName).First(&existingCustomer).Error; err != nil {
        if !errors.Is(err, gorm.ErrRecordNotFound) {
        fmt.Println("Error saat memeriksa nama customer:", err)
        return
     }
        } else {
            fmt.Println("\nNama customer sudah ada, harap masukan nama customer yang lain")
            continue
        }

        fmt.Print("Masukkan Alamat Customer : ")
        longString, _ = reader.ReadString('\n')
        Customer.Address = longString

        fmt.Print("Masukkan Nomor Telepon Customer : ")
        longString, _ = reader.ReadString('\n')
        Customer.Phone = longString

        result := cs.DB.Create(&Customer)

        if result.Error != nil {
            fmt.Println("Error Saat Menambahkan Customer", result.Error)
            return
        }

        fmt.Printf("\nCustomer %s \nBerhasil ditambahkan!\n", Customer.CustomerName)

        fmt.Print("\n\t(1) :> Tambah Customer Lagi \n\t(9) :> Kembali \nMasukkan Pilihan : ")
        fmt.Scanln(&choice)
        if choice == 9 {
            return
        } else if choice != 1 {
            fmt.Println("Coba lagi")
            continue
        }
    }
}

func (cs *CustomersSystem) ViewCustomers() {
    for {
        var Customers []model.Customer
        var choice int

        result := cs.DB.Find(&Customers)
        
        if result.Error != nil {
            fmt.Println("Error Saat Menampilkan Customers", result.Error)
            return
        }

        for _, customer := range Customers {
            fmt.Printf("\nNama Pelanggan: %sAlamat: %sNomor Telepon: %s\nDibuat pada: %s\nDiperbarui pada: %s\n\n",
                customer.CustomerName, customer.Address, customer.Phone, customer.CreatedAt, customer.UpdatedAt)
        }


		fmt.Print("\n\t(9) :> Kembali\nMasukkan Pilihan : ")
	fmt.Scanln(&choice)
	if choice == 9 {
		return
	} else {
		fmt.Println("Coba lagi")
		}
	}
}

func (cs *CustomersSystem) DeleteCustomer() {
    for {
        var Customer model.Customer
        var choice int
        reader := bufio.NewReader(os.Stdin)

        fmt.Print("\nMasukkan Nama Customer yang akan dihapus : ")
        customerName, _ := reader.ReadString('\n')

        result := cs.DB.Where("customer_name = ?", customerName).First(&Customer)
        
        if result.Error != nil {
            fmt.Println("Error saat mencari customer:", result.Error)
            
            fmt.Print("\n\t(1) :> Hapus Customer Lagi \n\t(9) :> Kembali \nMasukkan Pilihan : ")
            fmt.Scanln(&choice)
            
            if choice == 9 {
                return
            } else if choice != 1 {
                fmt.Println("Coba lagi")
                continue
            }
            
            return
        }

        result = cs.DB.Delete(&Customer)
        
        if result.Error != nil {
            fmt.Println("Customer Tidak Ada", result.Error)
            
            return
        }

        fmt.Printf("\nCustomer '%s' berhasil dihapus!\n", customerName)

        
         fmt.Print("\n\t(1) :> Hapus Customer Lagi \n\t(9) :> Kembali \nMasukkan Pilihan : ")
         fmt.Scanln(&choice)
         
         if choice == 9 {
             return
         } else if choice != 1 {
             fmt.Println("Coba lagi")
             continue
         }
    }
}

func (cs *CustomersSystem) SelectCustomer(CustName string) (model.Customer, error) {
    var cust model.Customer
    result := cs.DB.Where("Customer_Name = ?", CustName).Find(&cust)
    if result.Error != nil {
        return cust, result.Error
    }

    return cust, nil
}
