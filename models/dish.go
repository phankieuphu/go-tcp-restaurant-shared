package models

import (
	"gorm.io/gorm"
)

type Dishes struct {
	gorm.Model
	Name         string  `gorm:"type:varchar(100);not null" json:"name"`
	Description  string  `gorm:"type:varchar(100);not null" json:"description"`
	Price        float64 `gorm:"type:decimal(10,2); not null;default:0.0 " json:"price"`
	Category     string  `gorm:"type:varchar(50)"`
	RestaurantID uint    `gorm:"not null"`
}

func (c *Dishes) TableName() string {
	return "dishes"
}
