package main

import (
	"awesomeProject/Challenge-H8-Stephanie/cli"
	"awesomeProject/Challenge-H8-Stephanie/config"
	"awesomeProject/Challenge-H8-Stephanie/helpers"
	"fmt"
	"time"
)

// clear screen terlebih dahulu

func main() {
	//define clear screen terlebih dahulu
	fmt.Println("Clearing screen ...")
	time.Sleep(1 * time.Second)
	helpers.ClearScreen()

	//inisialisasi database
	config.InitDB()

	//logic user login

	//db := config.DB //save config db to db var
	//print out menu after login section
	cli.ShowingMenu()

}
