package models

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	SKU            string         `json:"sku" gorm:"index"`
	Barcode        string         `json:"barcode"`
	Name           string         `json:"name"`
	Price          uint64         `json:"price"`
	PriceOriginal  uint64         `json:"price_original"`
	Image          string         `json:"image"`
	Description    string         `json:"description"`
	Stock          uint8          `json:"stock"`
	ItemCategoryID uint           `json:"item_category_id"`
	ItemCategory   ItemCategory   `json:"item_category"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
