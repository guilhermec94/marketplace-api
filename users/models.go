package users

import (
	"marketplace-api/location"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Username  string
	Email     string
	Password  string
	Addresses []location.Address `gorm:"foreignKey:UserID"`
}
