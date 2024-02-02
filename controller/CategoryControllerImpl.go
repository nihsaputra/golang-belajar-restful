package controller

import (
	"github.com/gofiber/fiber/v2"
	"golang-belajar-restful/halper"
	"golang-belajar-restful/model/request"
	"golang-belajar-restful/model/response"
	"golang-belajar-restful/service"
	"net/http"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func (c *CategoryControllerImpl) GetAll(ctx *fiber.Ctx) error {
	categoryResponses := c.CategoryService.GetAll()
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}
	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func (c *CategoryControllerImpl) GetById(ctx *fiber.Ctx) error {
	paramsInt, err := ctx.ParamsInt("id")
	halper.PanicIfError(err)
	categoryResponse := c.CategoryService.GetById(uint(paramsInt))

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func (c *CategoryControllerImpl) Create(ctx *fiber.Ctx) error {
	var categoryRequest request.CategoryRequest
	err := ctx.BodyParser(&categoryRequest)
	halper.PanicIfError(err)
	categoryResponse := c.CategoryService.Create(categoryRequest)

	webResponse := response.WebResponse{
		Code:   201,
		Status: "CREATE",
		Data:   categoryResponse,
	}

	return ctx.Status(http.StatusCreated).JSON(webResponse)
}

func (c *CategoryControllerImpl) Update(ctx *fiber.Ctx) error {
	var categoryRequest request.CategoryUpdateRequest
	ctx.BodyParser(&categoryRequest)
	categoryResponse := c.CategoryService.Update(categoryRequest)

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func (c *CategoryControllerImpl) Delete(ctx *fiber.Ctx) error {
	paramsInt, err := ctx.ParamsInt("id")
	halper.PanicIfError(err)
	resonse := c.CategoryService.Delete(uint(paramsInt))

	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   resonse,
	}

	return ctx.Status(http.StatusOK).JSON(webResponse)
}

func NewCategoryController(service service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: service,
	}
}
