package main

import (
	"emansulaeman/belajar-golang-restful-api/app"
	"emansulaeman/belajar-golang-restful-api/controller"
	"emansulaeman/belajar-golang-restful-api/helper"
	"emansulaeman/belajar-golang-restful-api/middleware"
	"emansulaeman/belajar-golang-restful-api/repository"
	"emansulaeman/belajar-golang-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// router := httprouter.New()

	// router.GET("/api/categories", categoryController.FindAll)
	// router.GET("/api/categories/:categoryId", categoryController.FindById)
	// router.POST("/api/categories", categoryController.Create)
	// router.PUT("/api/categories/:categoryId", categoryController.Update)
	// router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	// router.PanicHandler = exception.ErrorHandler

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
