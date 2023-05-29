package entity

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	*gorm.Model // call gorm

	ProductId    uint      `db:"product_id"`
	BuyerEmail   string    `db:"buyer_email"`
	BuyerAddress string    `db:"buyer_address"`
	OrderDate    time.Time `db:"order_date"`
	Product      Product   `gorm:"foreignkey:ProductId"`
}

func (o *Order) PrintDetail() {
	fmt.Println("Buyer email : ", o.BuyerAddress)
	fmt.Println("Buyer address : ", o.BuyerAddress)
	fmt.Println("Produk yang dibeli : ", o.Product)
	fmt.Println("Order Date : ", o.OrderDate)
}
