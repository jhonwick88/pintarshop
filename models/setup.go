package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:@tcp(127.0.0.1:3306)/pintarshop?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database broh")
	}

	db.AutoMigrate(&Item{})
	//db.Create(&Item{Name: "Buku Sidu 38", Price: 3000, PriceOriginal: 2500, Stock: 100, Description: "Buku Sinar Dunia"})

	DB = db

}
