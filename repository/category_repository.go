package repository

import "golang-belajar-restful/model/entity"

type CategoryRepository interface {
	FindAll() ([]entity.Category, error)
	FindById(id uint) (entity.Category, error)
	Save(category entity.Category) (entity.Category, error)
	Delete(category entity.Category) error
}
