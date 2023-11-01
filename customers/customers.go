package customers

import (
	"bufio"
	"fmt"
	"os"
	"sqlgo/model"

	"gorm.io/gorm"
)

type CustomersSystem struct {
	DB *gorm.DB
}

func (cs *CustomersSystem) AddCustomer() {
	var Customer model.Customer
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Nama Customer : ")
	longString, _ := reader.ReadString('\n')
	Customer.CustomerName = longString

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
}

func (cs *CustomersSystem) ListCustomers() ([]model.Customer, error) {
	var result = make([]model.Customer, 0)
	var qry = cs.DB.Table("customers").Find(&result)
	var err = qry.Error
	if err != nil {
		return nil, err
	}

	fmt.Println("Daftar Pelanggan:")
	for _, customer := range result {
		fmt.Printf("\nNama Pelanggan: %sAlamat: %sNomor Telepon: %s\nDibuat pada: %s\nDiperbarui pada: %s\n\n",
			customer.CustomerName, customer.Address, customer.Phone, customer.CreatedAt, customer.UpdatedAt)
	}

	return result, nil
}
