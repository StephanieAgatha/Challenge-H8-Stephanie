package cli

import (
	"awesomeProject/Challenge-H8-Stephanie/config"
	"awesomeProject/Challenge-H8-Stephanie/entity"
	"awesomeProject/Challenge-H8-Stephanie/helpers"
	"bufio"
	"fmt"
	"os"
)

func listProduct() {
	helpers.ClearScreen()

	consoleReader := bufio.NewReader(os.Stdin)
	var products []entity.Product
	//inisiasliasi error
	err := config.DB.Find(&products).Error

	if err != nil {
		HandleError(err)
		return
	}

	fmt.Println("------ List Product ------")
	//ambil dari range products
	for _, product := range products {
		product.PrintDetail()
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
