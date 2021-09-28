package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Username  string         `json:"username" gorm:"comment:Nama login pengguna"`
	Password  string         `json:"-"  gorm:"comment:Kata sandi masuk pengguna"`
	NickName  string         `json:"nickname" gorm:"default:sysUser;comment:User's Nickname"`
	Email     string         `json:"email" gorm:"type:varchar(100);unique_index"`
	Role      uint           `json:"role"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
