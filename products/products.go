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
    for {
        var Product model.Product
        var choice int
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

        fmt.Printf("\n       Produk %s \nBerhasil ditambahkan oleh %s!\n", Product.ProductName, user.Username)

        fmt.Print("\n\t(1) :> Tambah Barang Lagi \n\t(9) :> Kembali \nMasukkan Pilihan : ")
        fmt.Scanln(&choice)
        if choice == 9 {
            return
        } else if choice != 1 {
            fmt.Println("Coba lagi")
            continue
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
		fmt.Printf("\nNama Produk: %sDeskripsi: %sStok: %d\nHarga: %d\nDibuat oleh: %s\n", product.ProductName, product.Description, product.Stok, product.Price, product.Username)
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
    for {
        var Product model.Product
        var choice int
        reader := bufio.NewReader(os.Stdin)

        fmt.Print("Masukkan Nama Produk yang ingin diubah :")
        ProductName, _ := reader.ReadString('\n')

        result := ps.DB.First(&Product, "product_name = ?", ProductName)

        if result.Error != nil {
            fmt.Println("Produk Tidak Ditemukan", result.Error)
            fmt.Print("\n\t(1) :> Ubah Barang Lagi \n\t(9) :> Kembali \nMasukkan Pilihan : ")
            fmt.Scanln(&choice)
            if choice == 9 {
                return
            } else if choice != 1 {
                fmt.Println("Coba lagi")
                continue
            }
        } else {
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

            fmt.Printf("\nProduk %s berhasil diubah!\n", newName)

            fmt.Print("\n\t(1) :> Ubah Barang Lagi \n\t(9) :> Kembali \nMasukkan Pilihan : ")
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
