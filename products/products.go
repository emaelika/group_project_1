package products

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sqlgo/model"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type ProductssSystem struct {
	DB *gorm.DB
}

func (ps *ProductssSystem) AddProduct(user model.User) {
	reader := bufio.NewReader(os.Stdin)

	for {
		var Product model.Product

		fmt.Print("\nMasukkan Nama Produk : ")
		longString, _ := reader.ReadString('\n')
		longString = strings.TrimSpace(longString)

		var existingProduct model.Product
		if err := ps.DB.Where("product_name = ?", longString).First(&existingProduct).Error; err != nil {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				fmt.Println("Error saat memeriksa nama produk:", err)
				return
			}
		} else {
			fmt.Println("\nNama produk sudah ada, harap masukan nama produk yang lain")
			continue
		}

		Product.ProductName = longString

		fmt.Print("Masukkan Deskripsi Produk : ")
		longString, _ = reader.ReadString('\n')
		Product.Description = strings.TrimSpace(longString)

		fmt.Print("Masukkan Stok Produk : ")
		stokStr, _ := reader.ReadString('\n')
		stokStr = strings.TrimSpace(stokStr)
		Product.Stok, _ = strconv.Atoi(stokStr)

		fmt.Print("Masukkan Harga Produk : ")
		priceStr, _ := reader.ReadString('\n')
		priceStr = strings.TrimSpace(priceStr)
		Product.Price, _ = strconv.ParseFloat(priceStr, 64)

		Product.UserID = user.UserID
		Product.Username = user.Username

		result := ps.DB.Create(&Product)

		if result.Error != nil {
			fmt.Println("Error Saat Menambahkan Produk", result.Error)
			return
		}

		fmt.Printf("\n       Produk '%s' \nBerhasil ditambahkan oleh %s!\n", Product.ProductName, user.Username)

		fmt.Print("\n\t(1) :> Tambah Barang Lagi \n\t(9) :> Kembali \nMasukkan Pilihan : ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, _ := strconv.Atoi(choiceStr)

		if choice == 9 {
			return
		} else if choice != 1 {
			fmt.Println("Coba lagi")
		}
	}
}

func (ps *ProductssSystem) ViewProduct() {
	var Products []model.Product
	var choice int

	result := ps.DB.Find(&Products)
	if result.Error != nil {
		fmt.Println("Error Saat Menampilkan Produk", result.Error)
		return
	}

	for _, product := range Products {
		fmt.Printf("\nNama Produk : %s\nDeskripsi : %s\nStok : %d\nHarga : %g\nDibuat oleh : %s\n", product.ProductName, product.Description, product.Stok, product.Price, product.Username)
	}

	fmt.Print("\n\t(9) :> Kembali\nMasukkan Pilihan : ")
	fmt.Scanln(&choice)
	if choice == 9 {
		return
	} else {
		fmt.Println("Coba lagi")
	}
}

func (ps *ProductssSystem) UpdateProduct() {
	reader := bufio.NewReader(os.Stdin)
	var choice int
	var choiceStr string

	for {
		var Product model.Product

		fmt.Print("Masukkan Nama Produk yang ingin diubah :")
		ProductName, _ := reader.ReadString('\n')
		ProductName = strings.TrimSpace(ProductName)

		result := ps.DB.First(&Product, "product_name = ?", ProductName)

		if result.Error != nil {
			fmt.Println("\nProduk Tidak Ditemukan", result.Error)
			fmt.Print("\n\t(1) :> Ubah Barang Lagi \n\t(9) :> Kembali \nMasukkan Pilihan : ")
			fmt.Scanln(&choice)
			if choice == 9 {
				return
			} else if choice != 1 {
				fmt.Println("Coba lagi")
				continue
			}
		} else {
			fmt.Printf("\n\t(1) :> Edit Informasi %s \n\t(2) :> Update Stok %s \nMasukkan Pilihan : ", ProductName, ProductName)

			choiceStr, _ = reader.ReadString('\n')
			choiceStr = strings.TrimSpace(choiceStr)
			choice, _ := strconv.Atoi(choiceStr)

			switch choice {
			case 1:
				fmt.Print("Masukkan Nama Produk baru : ")
				newName, _ := reader.ReadString('\n')
				Product.ProductName = strings.TrimSpace(newName)

				fmt.Print("Masukkan Deskripsi Produk Baru : ")
				newDescription, _ := reader.ReadString('\n')
				Product.Description = strings.TrimSpace(newDescription)

				fmt.Print("Masukkan Harga Produk baru: ")
				priceStr, _ := reader.ReadString('\n')
				priceStr = strings.TrimSpace(priceStr)
				Product.Price, _ = strconv.ParseFloat(priceStr, 64)
			case 2:
				fmt.Print("Masukkan Stok Produk Baru : ")
				stokStr, _ := reader.ReadString('\n')
				stokStr = strings.TrimSpace(stokStr)
				Product.Stok, _ = strconv.Atoi(stokStr)

			default:
				fmt.Println("Pilihan tidak valid, coba lagi")
				continue
			}

			result = ps.DB.Save(&Product)
			if result.Error != nil {
				fmt.Println("Error Saat Mengubah Produk", result.Error)
				return
			}

			fmt.Printf("\nProduk %s berhasil diubah!\n", Product.ProductName)

			fmt.Print("\n\t(1) :> Ubah Barang Lagi \n\t(9) :> Kembali \nMasukkan Pilihan : ")
			choiceStr, _ = reader.ReadString('\n')
			choiceStr = strings.TrimSpace(choiceStr)
			choice, _ = strconv.Atoi(choiceStr)

			if choice == 9 {
				return
			} else if choice != 1 {
				fmt.Println("Coba lagi")
			}
		}
	}
}

func (ps *ProductssSystem) DeleteProduct() {
	for {
		var Product model.Product
		var choice int
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Masukkan Nama Produk yang ingin dihapus : ")
		productName, _ := reader.ReadString('\n')

		result := ps.DB.Delete(&Product, "product_name = ?", productName)

		if result.Error != nil {
			fmt.Println("Produk tidak ditemukan.")
			fmt.Print("\n\t(1) :> Hapus Barang Lagi \n\t(9) :> Kembali \nMasukkan Pilihan : ")
			fmt.Scanln(&choice)
			if choice == 9 {
				return
			} else if choice != 1 {
				fmt.Println("Coba lagi")
				continue
			}
		} else {
			fmt.Printf("\nProduk %s berhasil dihapus!\n", productName)

			fmt.Print("\n\t(1) :> Hapus Barang Lagi \n\t(9) :> Kembali \nMasukkan Pilihan : ")
			fmt.Scanln(&choice)
			if choice == 9 {
				return
			} else if choice != 1 {
				fmt.Println("Coba lagi")
				continue
			}
		}
	}
}

func (ps *ProductssSystem) SelectProduct(ProdName string) (model.Product, error) {
	var produk model.Product
	err := ps.DB.Where("Product_Name = ?", ProdName).First(&produk).Error
	if err != nil {
		return produk, err
	}

	return produk, nil
}

func (ps *ProductssSystem) MinusStock(Sale int, Produk model.Product) {
	Produk.Stok -= Sale

	// Update the product in the database.
	ps.DB.Save(&Produk)
}
