package receipts

import (
	"fmt"
	"sqlgo/customers"
	"sqlgo/model"
	"sqlgo/products"
	"time"

	"gorm.io/gorm"
)

type ReceiptsSystem struct {
	DB *gorm.DB
}

func (rs *ReceiptsSystem) CreateReceipt(cs *customers.CustomersSystem, ps *products.ProductssSystem, user model.User) {
	var nota model.Receipt
	// cari customer , pake nama, disimpan nama dan id
	var pembeli model.Customer
	fmt.Print("Masukkan Nama Customer: ")
	var paijo string
	fmt.Scanln(&paijo)
	pembeli, err := cs.SelectCustomer(paijo)
	if err != nil {
		fmt.Println("Customer tidak ditemukan")
	}
	nota.CustomerID = pembeli.CustomerID
	nota.CustomerName = pembeli.CustomerName

	// cari barang, pake nama disimpan nama dan id
	var keranjang []model.Product
	var jumlahisi []int
	var Total float64
	for {
		var item model.Product
		fmt.Print("Masukkan nama barang: \n")
		fmt.Print("9. Untuk selesai menambahkan Barang")
		fmt.Print("\n0. Untuk keluar\n")
		fmt.Scanln(&paijo)
		if paijo == "9" {
			now := time.Now()
			//tampilkan nota
			fmt.Printf("Nota Pembelian:\nCustomer: %v\tKasir: %v\n Tanggal : %v ", nota.CustomerName, user.Username, now)
			for j := 0; j <= len(keranjang); j++ {

				SubTotal := keranjang[j].Price * float64(jumlahisi[j])
				fmt.Printf("%v\t%v\t%v", keranjang[j].ProductName, jumlahisi[j], SubTotal)
			}
			fmt.Println("Total: ", Total)
			fmt.Print("\nPress any key to continue: ")
			fmt.Scanln(&paijo)
			if paijo != "" {
				continue
			}
			nota.ProductID = keranjang[0].ProductID
			nota.Price = keranjang[0].Price
			nota.Quantity = jumlahisi[0]
			nota.Total = float64(jumlahisi[0]) * keranjang[0].Price
			result := rs.DB.Create(&nota)
			if result.Error != nil {
				fmt.Println("Error Saat Menambahkan Nota", result.Error)
				return
			}
			ps.MinusStock(jumlahisi[0], keranjang[0])

			uid := nota.TransactionID

			for i := 1; i <= len(keranjang); i++ {
				nota.ProductID = keranjang[i].ProductID
				nota.Price = keranjang[i].Price
				nota.Quantity = jumlahisi[i]
				nota.Total = keranjang[i].Price * float64(jumlahisi[i])
				nota.TransactionID = uid
				result := rs.DB.Create(&nota)
				if result.Error != nil {
					fmt.Println("Error Saat Menambahkan Nota", result.Error)
					return
				}
				ps.MinusStock(jumlahisi[i], keranjang[i])

			}

			return
		} else if paijo == "0" {
			return
		}
		item, err = ps.SelectProduct(paijo)
		if err != nil {
			fmt.Println("Barang tidak ditemukan")
		}

		fmt.Print("Masukkan Jumlah pembelian: ")
		var count int
		fmt.Scanln(&count)
		if count > item.Stok {
			fmt.Println("Jumlah terlalu banyak")
			return
		} else if count < 1 {
			fmt.Println("Error")
			return
		}
		Total += (item.Price * float64(count))
		fmt.Printf("Subtotal saat ini: %v", Total)
		fmt.Print("\nTekan 9 untuk membatalkan nota, tekan yang lain untuk melanjutkan input barang: ")
		fmt.Scanln(&paijo)
		if paijo == "9" {
			return
		}
		keranjang = append(keranjang, item)
		jumlahisi = append(jumlahisi, count)
	}

	fmt.Printf("a")

	// total barang a,  b atau c
	// total harga per barang
	// total harga receipt
	// save ke receipt dan transaksi

}
