package app

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rozanlaudzai/go-mysql-restful-api/controller"
	"github.com/rozanlaudzai/go-mysql-restful-api/exception"
)

func NewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()

	// setup endpoints
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.DeleteById)

	// setup panic handler
	router.PanicHandler = exception.ErrorHandler

	return router
}
