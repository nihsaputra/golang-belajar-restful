package app

import (
	"github.com/gofiber/fiber/v2"
	"golang-belajar-restful/controller"
)

var Router = *fiber.New()

func NewCategoryRouter(controller controller.CategoryController) {
	Router.Get("/category", controller.GetAll)
	Router.Get("/category/:id", controller.GetById)
	Router.Post("/category", controller.Create)
	Router.Put("/category", controller.Update)
	Router.Delete("/category/:id", controller.Delete)
}
