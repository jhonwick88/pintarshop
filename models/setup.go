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

	//db.AutoMigrate(&Item{})
	//db.AutoMigrate(&ItemCategory{})
	//db.AutoMigrate(&Customer{})
	//db.AutoMigrate(&User{})
	//db.AutoMigrate(&Order{})
	//db.AutoMigrate(&OrderDetail{})
	//db.AutoMigrate(&CartItem{})
	// db.Create(&Customer{
	// 	Name:    "Pelanggan",
	// 	Address: "Default Address",
	// 	Phone:   "N/A",
	// })
	// db.Create(&User{
	// 	Username: "Admin",
	// 	Password: "password",
	// 	NickName: "Admin",
	// 	Email:    "admin@pintarmedia.net",
	// 	Role:     1,
	// })
	// db.Create(&ItemCategory{
	// 	Name:        "Default",
	// 	Description: "Default Item Category",
	// })
	// db.Create(&Order{
	// 	OrderNumber: "1646176816",
	// 	OrderQty:    2,
	// 	Discount:    2000,
	// 	SubTotal:    6500,
	// 	Tax:         0,
	// 	Total:       4500,
	// 	UserID:      1,
	// 	CustomerID:  1,
	// })
	// db.Create(&OrderDetail{
	// 	SKU:        "BKSD38",
	// 	Name:       "Buku Sidu 38",
	// 	Price:      3500,
	// 	Qty:        1,
	// 	TotalPrice: 3500,
	// 	OrderID:    1,
	// 	ItemID:     1,
	// })
	// db.Create(&OrderDetail{
	// 	SKU:        "BKVS38",
	// 	Name:       "Buku Vision 38",
	// 	Price:      3000,
	// 	Qty:        1,
	// 	TotalPrice: 3000,
	// 	OrderID:    1,
	// 	ItemID:     3,
	// })
	//db.Create(&Item{Name: "Buku Sidu 38", Price: 3000, PriceOriginal: 2500, Stock: 100, Description: "Buku Sinar Dunia"})
	DB = db

}
