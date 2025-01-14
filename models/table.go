package models

import "gorm.io/gorm"

type Tables struct {
	gorm.Model
	// number of table
	Number      int     `gorm:"not  null" json:"number"`
	Status      string  `gorm:"varchar(100); default:'active' " json:"status"`
	IsAvailable bool    `gorm:"default:true" json:"is_available"`
	Debt        float64 `gorm:"default:0" json:"debt"`
}
