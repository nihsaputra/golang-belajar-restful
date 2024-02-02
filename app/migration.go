package app

import (
	"golang-belajar-restful/halper"
	"golang-belajar-restful/model/entity"
	"gorm.io/gorm"
)

func NewMigration(db *gorm.DB) {
	err := db.AutoMigrate(entity.Category{}, entity.Product{}, entity.Customer{})
	halper.PanicIfError(err)
}
