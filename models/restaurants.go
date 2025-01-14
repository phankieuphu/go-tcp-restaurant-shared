package models

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	Name           string `gorm:"type:varchar(100)" json:"name"`
	Description    string `gorm:"type:varchar(100)" json:"description"`
	LocationString string `gorm:"type:varchar(100)" json:"location_string"`

	Latitude  float64 `gorm:"not null" json:"latitude"`
	Longitude float64 `gorm:"not null" json:"longitude"`

	City    string `gorm:"type:varchar(100)"`
	State   string `gorm:"type:varchar(100)"`
	Country string `gorm:"type:varchar(100)"`
	ZipCode string `gorm:"type:varchar(20)"`

	Dishes []Dishes `gorm:"foreignKey:RestaurantID"`
}

func (r Restaurant) TableName() string {
	return "restaurant"
}
