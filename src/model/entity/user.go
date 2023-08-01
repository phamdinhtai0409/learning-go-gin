package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email       string
	UserName    string
	Password    string
	Token       *string
	LastLoginAt *time.Time
}
