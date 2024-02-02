package main

import (
	"golang-belajar-restful/app"
	"golang-belajar-restful/controller"
	"golang-belajar-restful/repository"
	"golang-belajar-restful/service"
)

func main() {
	// Database
	mysqlDB := app.NewDatabaseMysql()

	// Migration
	app.NewMigration(mysqlDB)

	// Repository
	categoryRepository := repository.NewCategoryRepository(mysqlDB)

	// Service
	categoryService := service.NewCategoryService(categoryRepository)

	// Controller
	categoryController := controller.NewCategoryController(categoryService)

	// Router
	app.NewCategoryRouter(categoryController)

	router := app.Router
	router.Listen(":8080")
}
