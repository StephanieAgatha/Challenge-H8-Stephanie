package cli

import (
	"fmt"
	"os"
)

type User struct {
	ID       int
	Username string
	Password string
}

func ShowingMenu() {

	var input int
	//run login section first

	fmt.Println("Silahkan Pilih:")
	fmt.Println("1. Register")
	fmt.Println("2. Login")
	fmt.Println("3. Exit")

	//user input
	fmt.Print("Enter your choice : ")
	fmt.Scanln(&input)

	if input == 1 {
		Register()
	}

	if input == 2 {
		Login()
	}

	if input == 3 {
		fmt.Println("Exiting program")
		os.Exit(1)
	}

	for {

		fmt.Println("Doing Hacktiv8 challenge ^_^ ")
		fmt.Println("Mini E-commerce by Stephanie")
		fmt.Println("Choose your menu below ")

		//create logic untuk showmenu setelah login berhasil
		fmt.Println("1. Show Available Product")
		fmt.Println("2. Crate status order")
		fmt.Println("3. Exit")

		//user input
		fmt.Print("Enter your choice : ")
		fmt.Scanln(&input)

		if input == 1 {
			listProduct()
		}
		if input == 2 {
			listOrder()
		}
		if input == 3 {
			fmt.Println("Sayonara :)")
			os.Exit(1)
		} else {
			fmt.Println("Incorrect input")
		}
	}

}
