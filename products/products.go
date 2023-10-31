package products

import (
	"bufio"
	"fmt"
	"os"
	"sqlgo/model"

	"gorm.io/gorm"
)

type ProductssSystem struct {
	DB *gorm.DB
}

func (ps *ProductssSystem) AddProduct(user model.User) {
	var Product model.Product
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Nama Produk : ")
	longString, _ := reader.ReadString('\n')
	Product.ProductName = longString
	fmt.Print("Masukkan Deskripsi Produk : ")

	longString, _ = reader.ReadString('\n')

	Product.Description = longString
	fmt.Print("Masukkan Stok Produk : ")
	fmt.Scanln(&Product.Stok)

	fmt.Print("Masukkan Harga Produk : ")
	fmt.Scan(&Product.Price)

	Product.UserID = user.UserID

	result := ps.DB.Create(&Product)

	if result.Error != nil {
		fmt.Println("Error Saat Menambahkan Produk", result.Error)
		return
	}

	fmt.Printf("\nProduk '%s' berhasil ditambahkan oleh %s!\n", Product.ProductName, user.Username)
}
