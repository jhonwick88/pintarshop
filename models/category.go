package models

import (
	"time"

	"gorm.io/gorm"
)

type ItemCategory struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Image       string         `json:"image"`
	Description string         `json:"description"`
	ParentId    *uint          `json:"parent_id"`
	Parent      *ItemCategory  `json:"parent"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
