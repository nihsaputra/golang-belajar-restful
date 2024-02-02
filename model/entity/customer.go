package entity

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Name    string `gorm:"column:name"`
	Address string `gorm:"address"`
}
