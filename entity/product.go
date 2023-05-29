package entity

import (
	"fmt"
	"gorm.io/gorm"
)

// membuat struct product
type Product struct {
	*gorm.Model

	Name        string `db:"name"`
	Price       int64  `db:"price"`
	Description string `db:"decription"`
}

func (p *Product) PrintDetail() {
	fmt.Println("ID Produk : ", p.ID)
	fmt.Println("Nama : ", p.Name)
	fmt.Println("Deskripsi Produk : ", p.Description)
	fmt.Println("Harga Barang : ", p.Price)
}
