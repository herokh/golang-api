package models

import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"primaryKey;column:id"`
	Email    string `gorm:"column:email"`
	Password string `gorm:"column:password"`
	gorm.Model
}
