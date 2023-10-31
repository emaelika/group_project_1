package products

import (
	"fmt"
	"sqlgo/model"

	"gorm.io/gorm"
)

type ProductssSystem struct {
	DB *gorm.DB
}

func (ps *ProductssSystem) AddProduct(user model.User) {
	var Product model.Product

	fmt.Print("Masukkan Nama Produk : ")
	fmt.Scanln(&Product.ProductName)

	fmt.Print("Masukkan Deskripsi Produk : ")
	fmt.Scanln(&Product.Description)

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
