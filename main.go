package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/rozanlaudzai/go-mysql-restful-api/app"
	"github.com/rozanlaudzai/go-mysql-restful-api/controller"
	"github.com/rozanlaudzai/go-mysql-restful-api/exception"
	"github.com/rozanlaudzai/go-mysql-restful-api/middleware"
	"github.com/rozanlaudzai/go-mysql-restful-api/repository"
	"github.com/rozanlaudzai/go-mysql-restful-api/service"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	db, err := app.NewDB()
	if err != nil {
		panic(err)
	}
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	// setup endpoints
	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.DeleteById)
	router.PanicHandler = exception.ErrorHandler

	// setup address
	serverPort := os.Getenv("SERVER_PORT")
	address := fmt.Sprintf("localhost:%v", serverPort)

	// setup auth middleware
	apiKey := os.Getenv("API_KEY")
	authMiddleware := middleware.NewAuthMiddleware(router, apiKey)

	server := http.Server{
		Addr:    address,
		Handler: authMiddleware,
	}
	fmt.Printf("Listening to http://%v\n", address)
	if err = server.ListenAndServe(); err != nil {
		panic(err)
	}

}
