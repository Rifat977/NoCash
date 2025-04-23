package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique"`
	Phone    string `gorm:"unique"`
	Password string
	Balance  float64
}

type Session struct {
	gorm.Model
	UserID uint
	Token  string `gorm:"unique"`
	Expiry int64
}

func GetAll() []interface{} {
	return []interface{}{
		&User{},
		&Session{},
	}
}
