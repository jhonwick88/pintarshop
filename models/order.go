package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	OrderNumber  string         `json:"order_number" gorm:"type:varchar(200)"`
	NoResi       string         `json:"no_resi" gorm:"type:varchar(150)"`
	OrderQty     uint8          `json:"order_qty"`
	SubTotal     uint64         `json:"sub_total"`
	Discount     uint64         `json:"discount"`
	Tax          uint64         `json:"tax"`
	Total        uint64         `json:"total"`
	Status       string         `json:"status" gorm:"type:varchar(100);default:lunas"`
	UserID       uint           `json:"user_id"`
	User         User           `json:"user"`
	CustomerID   uint           `json:"customer_id"`
	Customer     Customer       `json:"customer"`
	OrderDetails []OrderDetail  `json:"order_details" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

//

type OrderDetail struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	SKU        string `json:"sku"`
	Name       string `json:"name"`
	Price      uint64 `json:"price"`
	Qty        uint8  `json:"qty"`
	TotalPrice uint64 `json:"total_price"`
	OrderID    uint
	//Order      Order
	ItemID uint
	//Item       Item
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
