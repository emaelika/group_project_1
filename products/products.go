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

	Product.Username = user.Username
	Product.UserID = user.UserID

	result := ps.DB.Create(&Product)

	if result.Error != nil {
		fmt.Println("Error Saat Menambahkan Produk", result.Error)
		return
	}

	fmt.Printf("\nProduk %s \nBerhasil ditambahkan oleh %s!\n", Product.ProductName, user.Username)
}


func (ps *ProductssSystem) ViewProduct() {
	var Products []model.Product

	result := ps.DB.Find(&Products)
	if result.Error != nil {
		fmt.Println("Error Saat Menampilkan Produk", result.Error)
		return
	}

	for _, product := range Products {
		fmt.Printf("\nNama Produk: %sDeskripsi: %sStok: %d\nHarga: %d\nDibuat oleh: %s\n", product.ProductName, product.Description, product.Stok, product.Price, product.Username)
	}
}


func (ps *ProductssSystem) UpdateProduct() {
	var Product model.Product
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Nama Produk yang ingin diubah :")
	ProductName, _ := reader.ReadString('\n')

	result := ps.DB.First(&Product, "product_name = ?", ProductName)

	if result.Error != nil {
		fmt.Println("Produk Tidak Ditemukan", result.Error)
		return
	}

	fmt.Print("Masukkan Nama Produk baru : ")
	newName, _ := reader.ReadString('\n')
	Product.ProductName = newName

	fmt.Print("Masukkan Deskripsi Produk Baru : ")
	newDescription, _ := reader.ReadString('\n')
	Product.Description = newDescription

	fmt.Print("Masukkan Stok Produk Baru : ")
	fmt.Scanln(&Product.Stok)

	fmt.Print("Masukkan Harga Produk baru: ")
	fmt.Scan(&Product.Price)

	result = ps.DB.Save(&Product)
	if result.Error != nil {
		fmt.Println("Error Saat Mengubah Produk", result.Error)
		return
	}

	fmt.Printf("\nProduk %s Berhasil Diubah!\n", Product.ProductName)
}

func (ps *ProductssSystem) DeleteProduct() {
	var Product model.Product
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan Nama Produk yang ingin dihapus : ")
	productName, _ := reader.ReadString('\n')

	result := ps.DB.Delete(&Product, "product_name = ?", productName)

	if result.Error != nil {
		fmt.Println("Error Saat Menghapus Produk", result.Error)
		return
	}
	fmt.Printf("\nProduk %s berhasil dihapus!\n", productName)
}