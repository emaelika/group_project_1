package transactions

import (
	"fmt"
	"fmt"
	"sqlgo/model"

	"gorm.io/gorm"
)

type TransactionsSystem struct {
	DB *gorm.DB
}


// func (ts *TransactionsSystem) AddTransaction (ps *products.ProductssSystem, cs *customers.CustomersSystem, user model.User) {
	
// 	var custName string
// 	var pembeli model.Customer
	
// 	fmt.Print("Masukkan nama customer : ")
// 	fmt.Scanln(&custName)
// 	pembeli, err := cs.SelectCustomerByName(custName)
// 	if err !=  nil {
// 		fmt.Println("Customer Tidak Ditemukan")
// 		return
// 	}

// 	Transaksi.CustomerID = pembeli.CustomerID

// 	fmt.Print("Masukkan Produk yang dibeli: ")
// 	var prodName string
// 	fmt.Scanln(&prodName)
// 	var item model.Product
// 	item, err = ps.SelectProductByName(prodName)
// 	if err != nil {
// 		fmt.Println("Produk Tidak Ditemukan")
// 		return
// 	}
	
// 	fmt.Printf("%s, Harga: %v, Stok: %v\n", item.ProductName, item.Price, item.Stok)
// 	fmt.Print("Masukkan Jumlah pembelian: ")
// 	var jumlah int
// 	fmt.Scanln(&jumlah)
// 	if jumlah > item.Stok {
// 		fmt.Println("Anda beli terlalu banyak")
// 		return
// 	}
// 	var SubTotal float64
// 	SubTotal = float64(item.Price) * float64(jumlah)
// 	Transaksi.Total = SubTotal
// 	Transaksi.Quantity = jumlah
// 	Transaksi.ProductID = item.ProductID
// 	fmt.Println(Transaksi)
// 	var input int
// 	fmt.Println("9. Batal\nInput any key to continue")
// 	fmt.Scanln(&input)
// 	if input == 9 {
// 		return
// 	}
	
// 	result := ts.DB.Create(&Transaksi)

// 	if result.Error != nil {
// 		fmt.Println("Error Saat Menambahkan Transaksi", result.Error)
// 		return
// 	}

// 	fmt.Printf("\nTransaksi Berhasil ditambahkan oleh %s!\n",  user.Username)
// }


func (ts *TransactionsSystem) DeleteTransaction(transactionID uint) {
	var transaction model.Transaction

	result := ts.DB.First(&transaction, transactionID)
	if result.Error != nil {
		fmt.Println("Transaksi tidak ditemukan")
		return
	}

	result = ts.DB.Delete(&transaction)
	if result.Error != nil {
		fmt.Println("Error saat menghapus transaksi", result.Error)
		return
	}

	fmt.Println("Transaksi berhasil dihapus")
}

func (ts *TransactionsSystem) ViewTransactions() {
	var List []model.Transaction

	result := ts.DB.Find(&List)
	if result.Error != nil {
		fmt.Println("Error Saat Menampilkan Transaksi", result.Error)
		return
	}

	for _, Transaction := range List {
		fmt.Printf("\nNomor Nota: %v, Tanggal: %v, Customer: %v, Kasir: %v, Barang: %v, Harga: %v, Jumlah: %v, Total: %v\n", Transaction.TransactionID, Transaction.CreatedAt, Transaction.CustomerID, Transaction.UserID, Transaction.ProductID, Transaction.Price, Transaction.Quantity, Transaction.Total)
	}


// func (ts *TransactionsSystem) AddTransaction (ps *products.ProductssSystem, cs *customers.CustomersSystem, user model.User) {
// 	var Transaksi model.Transaction

// 	var custName string
// 	var pembeli model.Customer

// 	fmt.Print("Masukkan nama customer : ")
// 	fmt.Scanln(&custName)
// 	pembeli, err := cs.SelectCustomerByName(custName)
// 	if err !=  nil {
// 		fmt.Println("Customer Tidak Ditemukan")
// 		return
// 	}

// 	Transaksi.CustomerID = pembeli.CustomerID

// 	fmt.Print("Masukkan Produk yang dibeli: ")
// 	var prodName string
// 	fmt.Scanln(&prodName)
// 	var item model.Product
// 	item, err = ps.SelectProductByName(prodName)
// 	if err != nil {
// 		fmt.Println("Produk Tidak Ditemukan")
// 		return
// 	}

// 	fmt.Printf("%s, Harga: %v, Stok: %v\n", item.ProductName, item.Price, item.Stok)
// 	fmt.Print("Masukkan Jumlah pembelian: ")
// 	var jumlah int
// 	fmt.Scanln(&jumlah)
// 	if jumlah > item.Stok {
// 		fmt.Println("Anda beli terlalu banyak")
// 		return
// 	}
// 	var SubTotal float64
// 	SubTotal = float64(item.Price) * float64(jumlah)
// 	Transaksi.Total = SubTotal
// 	Transaksi.Quantity = jumlah
// 	Transaksi.ProductID = item.ProductID
// 	fmt.Println(Transaksi)
// 	var input int
// 	fmt.Println("9. Batal\nInput any key to continue")
// 	fmt.Scanln(&input)
// 	if input == 9 {
// 		return
// 	}

// 	result := ts.DB.Create(&Transaksi)

// 	if result.Error != nil {
// 		fmt.Println("Error Saat Menambahkan Transaksi", result.Error)
// 		return
// 	}

// 	fmt.Printf("\nTransaksi Berhasil ditambahkan oleh %s!\n",  user.Username)
// }

func (ts *TransactionsSystem) DeleteTransaction(transactionID uint) {
	var transaction model.Transaction

	result := ts.DB.First(&transaction, transactionID)
	if result.Error != nil {
		fmt.Println("Transaksi tidak ditemukan")
		return
	}

	result = ts.DB.Delete(&transaction)
	if result.Error != nil {
		fmt.Println("Error saat menghapus transaksi", result.Error)
		return
	}

	fmt.Println("Transaksi berhasil dihapus")
}

func (ts *TransactionsSystem) ViewTransactions() {
	var List []model.Transaction

	result := ts.DB.Find(&List)
	if result.Error != nil {
		fmt.Println("Error Saat Menampilkan Transaksi", result.Error)
		return
	}

	for _, Transaction := range List {
		fmt.Printf("\nNomor Nota: %v, Tanggal: %v, Customer: %v, Kasir: %v, Barang: %v, Harga: %v, Jumlah: %v, Total: %v\n", Transaction.TransactionID, Transaction.CreatedAt, Transaction.CustomerID, Transaction.UserID, Transaction.ProductID, Transaction.Price, Transaction.Quantity, Transaction.Total)
	}
}
