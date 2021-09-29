package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"golang.org/x/crypto/bcrypt"
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

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
func VerifyPassword(hashedPasword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPasword), []byte(password))
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// func (u *User) BeforeSave() error {
// 	hashedPassword, err := Hash(u.Password)
// 	if err != nil {
// 		return err
// 	}
// 	u.Password = string(hashedPassword)
// 	return nil
// }
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if u.Password == "" {
			return errors.New("Password required")
		}
		if u.Email == "" {
			return errors.New("Email required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Email invalid")
		}
		return nil
	default:
		if u.Username == "" {
			return errors.New("Username required")
		}
		if u.Password == "" {
			return errors.New("Password required")
		}
		if u.Email == "" {
			return errors.New("Email required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Email invalid")
		}
		return nil

	}
}
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}
