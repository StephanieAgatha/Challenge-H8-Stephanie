package cli

import (
	"awesomeProject/Challenge-H8-Stephanie/config"
	"awesomeProject/Challenge-H8-Stephanie/entity"
	"awesomeProject/Challenge-H8-Stephanie/helpers"
	"bufio"
	"fmt"
	"os"
	"time"
)

func listOrder() {
	helpers.ClearScreen()

	consoleReader := bufio.NewReader(os.Stdin)
	var orders []entity.Order
	//inisiasliasi error
	err := config.DB.Find(&orders).Error

	if err != nil {
		HandleError(err)
		return
	}

	fmt.Println("------ List Product ------")
	//ambil dari range products
	for _, order := range orders {
		order.PrintDetail()
	}
	var input string
	fmt.Println("Input order ID : ")
	fmt.Println("Press M for back to menu ")
	fmt.Println("Press q for quit application ")

	input, _ = consoleReader.ReadString('\n')

	switch input {
	case "m\n":
		ShowingMenu()
	case "q\n":
		fmt.Println("Sayonara ")
		os.Exit(1)
	default:
		OrderProduct(input)
	}
}
func OrderProduct(id string) {

	helpers.ClearScreen()

	consoleReader := bufio.NewReader(os.Stdin)

	var product entity.Product
	err := config.DB.Where("ID = ?", id).First(&product).Error
	//deklarasi
	if err != nil {
		fmt.Println("Got error")
		HandleError(err)
		return
	}

	product.PrintDetail()

	var input string
	fmt.Println("Press y for continue order")
	fmt.Println("Press m for back to menu")
	fmt.Println("Press q for quit app")

	//define user input
	input, _ = consoleReader.ReadString('\n')

	switch input {
	case "y\n":
		CreateOrder(product)
	case "m\n":
		listProduct()
	case "q\n":
		fmt.Println("Sayonara")
		os.Exit(1)
	default:
		OrderProduct(id)
	}
}

func CreateOrder(product entity.Product) {

	//clear screen
	helpers.ClearScreen()

	//read
	consolReader := bufio.NewReader(os.Stdin)
	//create variable email
	var email string
	//create variable address for user
	var address string

	fmt.Println("For continue the order,please fill below ")

	//input email
	fmt.Print("Your Email : ")
	//save to var email
	email, _ = consolReader.ReadString('\n')

	//input address
	fmt.Print("Your Address : ")
	address, _ = consolReader.ReadString('\n')

	//define custome order
	order := entity.Order{
		ProductId:    (product.ID),
		BuyerEmail:   email,
		BuyerAddress: address,
		OrderDate:    time.Now(),
	}
	//handle error
	err := config.DB.Create(&order).Error
	if err != nil {
		HandleError(err)
		return
	}
	//if order was successfully
	fmt.Println("Successfully Checkout")

	var input string
	fmt.Println("Press any key for return to main menu")
	fmt.Println("Press q for exit app")
	_, err = fmt.Scanln(&input)
	if err != nil {
		ShowingMenu()
	}

	switch input {
	case "q\n":
		fmt.Println("Bye Bye !")
		os.Exit(1)
	default:
		ShowingMenu()
	}
}
