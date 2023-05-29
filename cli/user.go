package cli

import (
	"awesomeProject/Challenge-H8-Stephanie/config"
	"awesomeProject/Challenge-H8-Stephanie/entity"
	"awesomeProject/Challenge-H8-Stephanie/helpers"
	"fmt"
	"time"
)

func Register() {
	var username string
	var password string

	fmt.Print("Input your username : ")
	fmt.Scan(&username)
	fmt.Print("Input yout password : ")
	fmt.Scan(&password)

	user := entity.User{Username: username, Password: password}
	err := config.DB.Create(&user).Error

	if err != nil {
		HandleError(err)
	}

	//print out success register
	fmt.Println("Successfully Register ")
	//sleep
	time.Sleep(3 * time.Second)
	fmt.Println("Loggin in ...")
	Login()
}

func Login() {
	helpers.ClearScreen()
	fmt.Println("Welcome to ecommerce app")
	var username string
	var password string

	fmt.Print("Input your username : ")
	fmt.Scan(&username)
	fmt.Print("Input yout password : ")
	fmt.Scan(&password)

	var users entity.User

	err := config.DB.Model(&entity.User{}).Where("username = ? AND password = ?", username, password).Find(&users).Error
	if err != nil {
		HandleError(err)
	}

	//logic untuk invalid username / password
	if users.Username == "" {
		fmt.Println("Invalid username / password, Try Again")
		fmt.Println("Returning you to login page")
		time.Sleep(3 * time.Second)
		Login()
	}

	users.PrintSuccess()
}
