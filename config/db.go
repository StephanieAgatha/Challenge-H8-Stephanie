package config

import (
	"awesomeProject/Challenge-H8-Stephanie/entity"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	//create conf db
	dsn := "host=localhost user=postgres password=123456 dbname=db_steph1 port=5432 TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}) //buka koneksi postgres

	//handle err
	if err != nil {
		fmt.Println("Failed connect to database :( ")
		panic(err)
	} else {
		fmt.Println("Yeay...connected to database")
		DB = db
	}

	//migrate db
	db.AutoMigrate(&entity.Product{}, &entity.Order{}, &entity.User{})

}
