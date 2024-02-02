package service

import (
	"golang-belajar-restful/halper"
	"golang-belajar-restful/model/entity"
	"golang-belajar-restful/model/request"
	"golang-belajar-restful/model/response"
	"golang-belajar-restful/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
}

func (c *CategoryServiceImpl) GetAll() []response.CategoryResponse {
	var categoryResponses []response.CategoryResponse
	categories, err := c.CategoryRepository.FindAll()
	halper.PanicIfError(err)

	for _, category := range categories {
		categoryResponse := response.CategoryResponse{
			ID:        category.ID,
			Name:      category.Name,
			CreatedAt: category.CreatedAt,
			UpdatedAt: category.UpdatedAt,
		}
		categoryResponses = append(categoryResponses, categoryResponse)
	}
	return categoryResponses
}

func (c *CategoryServiceImpl) GetById(id uint) response.CategoryResponse {
	category, err := c.CategoryRepository.FindById(id)
	halper.PanicIfError(err)

	categoryResponse := response.CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}

	return categoryResponse
}

func (c *CategoryServiceImpl) Create(request request.CategoryRequest) response.CategoryResponse {
	category := entity.Category{
		Name: request.Name,
	}

	saveCategory, err := c.CategoryRepository.Save(category)
	halper.PanicIfError(err)

	categoryResponse := response.CategoryResponse{
		ID:        saveCategory.ID,
		Name:      saveCategory.Name,
		CreatedAt: saveCategory.CreatedAt,
		UpdatedAt: saveCategory.UpdatedAt,
	}

	return categoryResponse
}

func (c *CategoryServiceImpl) Update(request request.CategoryUpdateRequest) response.CategoryResponse {
	category, err := c.CategoryRepository.FindById(request.ID)
	halper.PanicIfError(err)
	category.Name = request.Name

	updateCategory, err := c.CategoryRepository.Save(category)
	halper.PanicIfError(err)

	categoryResponse := response.CategoryResponse{
		ID:        updateCategory.ID,
		Name:      updateCategory.Name,
		CreatedAt: updateCategory.CreatedAt,
		UpdatedAt: updateCategory.UpdatedAt,
	}

	return categoryResponse
}

func (c *CategoryServiceImpl) Delete(id uint) string {
	category, err := c.CategoryRepository.FindById(id)
	halper.PanicIfError(err)

	err = c.CategoryRepository.Delete(category)
	halper.PanicIfError(err)

	return "sukses"
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
	}
}
