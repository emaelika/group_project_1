package receipts

import (
	"fmt"
	"sqlgo/customers"
	"sqlgo/model"
	"sqlgo/products"

	"gorm.io/gorm"
)

type ReceiptsSystem struct {
	DB *gorm.DB
}

func (rs *ReceiptsSystem) CreateReceipt(cs *customers.CustomersSystem, ps *products.ProductssSystem, user model.User) {
	for {
		var nota model.Receipt

		// Cari customer berdasarkan nama dan simpan nama dan ID
		fmt.Print("Masukkan Nama Customer: ")
		var customerName string
		fmt.Scanln(&customerName)
		pembeli, err := cs.SelectCustomer(customerName)
		if err != nil {
			fmt.Println("Customer tidak ditemukan")
			return
		}

		// Membuat keranjang untuk menyimpan produk yang akan dibeli
		var keranjang []model.Product
		var jumlahisi []int
		var total float64

		for {
			var item model.Product
			fmt.Print("Masukkan nama barang: ")
			var productName string
			fmt.Scanln(&productName)

			if productName == "9" {
				break
			} else if productName == "0" {
				return
			}

			// Cari produk berdasarkan nama
			item, err = ps.SelectProduct(productName)
			if err != nil {
				fmt.Println("Barang tidak ditemukan")
				return
			}

			fmt.Print("Masukkan Jumlah pembelian: ")
			var count int
			fmt.Scanln(&count)
			if count > item.Stok {
				fmt.Println("Jumlah terlalu banyak")
				continue
			} else if count < 1 {
				fmt.Println("Error")
				continue
			}
			total += (item.Price * float64(count))
			fmt.Printf("Subtotal saat ini: %v\n", total)
			keranjang = append(keranjang, item)
			jumlahisi = append(jumlahisi, count)

			// Tampilkan detail keranjang
			fmt.Println("Detail Keranjang:")
			for j, product := range keranjang {
				subTotal := product.Price * float64(jumlahisi[j])
				fmt.Printf("%v\t%v\t%v\n", product.ProductName, jumlahisi[j], subTotal)
			}
			fmt.Printf("Total: %v\n", total)

			// Tampilkan menu
			fmt.Println("1. Simpan Nota")
			fmt.Println("9. Batal")
			fmt.Print("Masukkan pilihan: ")
			var input string
			fmt.Scanln(&input)
			if input == "1" {
				break
			} else if input == "9" {
				return
			}
		}

		if len(keranjang) == 0 {
			fmt.Println("Keranjang kosong. Transaksi dibatalkan.")
			return
		}

		// Simpan nota ke database
		uid := nota.TransactionID
		for i, product := range keranjang {
			nota.ProductID = product.ProductID
			nota.ProductName = product.ProductName
			nota.Price = product.Price
			nota.Quantity = jumlahisi[i]
			nota.Total = product.Price * float64(jumlahisi[i])
			nota.TransactionID = uid
			nota.CustomerName = pembeli.CustomerName
			nota.CustomerID = pembeli.CustomerID
			result := rs.DB.Create(&nota)
			if result.Error != nil {
				fmt.Println("Error saat menambahkan nota:", result.Error)
				return
			}
			ps.MinusStock(jumlahisi[i], product)
		}

		// Tampilkan menu utama
		fmt.Println("1. Buat Nota Lagi")
		fmt.Println("9. Kembali")
		fmt.Print("Masukkan pilihan: ")
		var input string
		fmt.Scanln(&input)
		if input == "1" {
			continue
		} else if input == "9" {
			break
		}
	}
}
func (rs *ReceiptsSystem) ViewReceipt() {
	var receipts []model.Receipt
	
	result := rs.DB.Find(&receipts)
	if result.Error != nil {
		fmt.Println("Error saat mencari nota:", result.Error)
		return
	}
	for _, item := range receipts {
		fmt.Printf("TransactionID: %v, CustomerID: %v, ProductID: %v, CustomerName: %v, ProductName: %v, Total: %v, Quantity: %v, Price: %v, CreatedAt: %v, UpdatedAt: %v\n",
			item.TransactionID, item.CustomerID, item.ProductID, item.CustomerName, item.ProductName, item.Total, item.Quantity, item.Price, item.CreatedAt, item.UpdatedAt)
	}
}

func (rs *ReceiptsSystem) DeleteReceipt() {
	fmt.Print("Masukkan Transaction ID: ")
	var transactionID uint
	fmt.Scanln(&transactionID)

	var receipt model.Receipt
	result := rs.DB.Where("transaction_id = ?", transactionID).Delete(&receipt)
	if result.Error != nil {
		fmt.Println("Error saat menghapus nota:", result.Error)
		return
	}
	fmt.Println("Nota berhasil dihapus.")
}
