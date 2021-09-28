package models

import "time"

type CartItem struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	OptionID   uint      `json:"option_id"`
	OptionName string    `json:"option_name"`
	ItemID     uint      `json:"item_id"`
	Item       Item      `json:"item"`
	Qty        uint8     `json:"qty"`
	UserID     uint      `json:"user_id"`
	User       User      `json:"user"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
