package location

import (
	"time"

	"gorm.io/gorm"
)

type Country struct {
	Code      string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Cities    []City
}

type City struct {
	gorm.Model
	Name        string
	CountryCode string
	Country     Country `gorm:"foreignKey:CountryCode"`
}

type Address struct {
	gorm.Model
	CityID      int64
	City        City `gorm:"foreignKey:CityID"`
	Street      string
	HouseNumber int
	PostalCode  string
	UserID      uint
}
