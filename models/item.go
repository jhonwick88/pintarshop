package models

import "gorm.io/gorm"

type Item struct {
	Name          string
	Price         uint64
	PriceOriginal uint64
	Description   string
	Stock         uint8
	gorm.Model
}
