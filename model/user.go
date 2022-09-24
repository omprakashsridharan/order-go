package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"not null" validate:"required,email"`
	Password string `gorm:"not null" validate:"required"`
	Phone    string
	Role     string `gorm:"not null" validate:"required"`
}
