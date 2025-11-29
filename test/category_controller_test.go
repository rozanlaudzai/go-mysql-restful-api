package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rozanlaudzai/go-mysql-restful-api/app"
	"github.com/rozanlaudzai/go-mysql-restful-api/controller"
	"github.com/rozanlaudzai/go-mysql-restful-api/middleware"
	"github.com/rozanlaudzai/go-mysql-restful-api/model/domain"
	"github.com/rozanlaudzai/go-mysql-restful-api/repository"
	"github.com/rozanlaudzai/go-mysql-restful-api/service"
	"github.com/stretchr/testify/assert"
)

func truncateCategory(db *sql.DB) error {
	_, err := db.Exec("TRUNCATE category")
	return err
}

func newDBTester() (*sql.DB, error) {
	// set up db tester
	db, err := app.NewDB()
	if err != nil {
		return db, err
	}
	err = truncateCategory(db)
	return db, err
}

func newRouterTester(db *sql.DB) (http.Handler, error) {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)
	// set auth middleware
	apiKey := os.Getenv("API_KEY")
	authMiddleware := middleware.NewAuthMiddleware(router, apiKey)
	return authMiddleware, nil
}

func TestCreateCategorySuccess(t *testing.T) {
	if err := godotenv.Load("../.env.test"); err != nil {
		panic(err)
	}
	db, err := newDBTester()
	if err != nil {
		panic(err)
	}
	router, err := newRouterTester(db)
	if err != nil {
		panic(err)
	}

	requestBody := strings.NewReader(`{"name": "Electronics"}`)
	url := fmt.Sprintf(
		"http://localhost:%v/api/categories",
		os.Getenv("SERVER_PORT"),
	)
	request := httptest.NewRequest(http.MethodPost, url, requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", os.Getenv("API_KEY"))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Electronics", responseBody["data"].(map[string]any)["name"])
}

func TestCreateCategoryBadRequest(t *testing.T) {
	if err := godotenv.Load("../.env.test"); err != nil {
		panic(err)
	}
	db, err := newDBTester()
	if err != nil {
		panic(err)
	}
	router, err := newRouterTester(db)
	if err != nil {
		panic(err)
	}

	requestBody := strings.NewReader(`{"name": ""}`)
	url := fmt.Sprintf(
		"http://localhost:%v/api/categories",
		os.Getenv("SERVER_PORT"),
	)
	request := httptest.NewRequest(http.MethodPost, url, requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", os.Getenv("API_KEY"))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	if err := godotenv.Load("../.env.test"); err != nil {
		panic(err)
	}

	db, err := newDBTester()
	if err != nil {
		panic(err)
	}

	// add a category to db
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	categoryRepository := repository.NewCategoryRepository()
	category, err := categoryRepository.Create(context.Background(), tx, domain.Category{
		Name: "Electronics",
	})
	if err != nil {
		panic(err)
	}
	if err = tx.Commit(); err != nil {
		panic(err)
	}

	router, err := newRouterTester(db)
	if err != nil {
		panic(err)
	}

	// update to from "Electronics" to "Pork"
	updatedName := "Pork"
	requestBody := strings.NewReader(fmt.Sprintf(`{"name": "%v"}`, updatedName))
	url := fmt.Sprintf(
		"http://localhost:%v/api/categories/%v",
		os.Getenv("SERVER_PORT"),
		category.Id,
	)
	request := httptest.NewRequest(http.MethodPut, url, requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", os.Getenv("API_KEY"))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(
		t,
		category.Id,
		int(responseBody["data"].(map[string]any)["id"].(float64)),
	)
	assert.Equal(
		t,
		updatedName,
		responseBody["data"].(map[string]any)["name"],
	)
}

func TestUpdateCategoryBadRequest(t *testing.T) {
	if err := godotenv.Load("../.env.test"); err != nil {
		panic(err)
	}

	db, err := newDBTester()
	if err != nil {
		panic(err)
	}

	// add a category to db
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	categoryRepository := repository.NewCategoryRepository()
	category, err := categoryRepository.Create(context.Background(), tx, domain.Category{
		Name: "Electronics",
	})
	if err != nil {
		panic(err)
	}
	if err = tx.Commit(); err != nil {
		panic(err)
	}

	router, err := newRouterTester(db)
	if err != nil {
		panic(err)
	}

	requestBody := strings.NewReader(`{"name": ""}`)
	url := fmt.Sprintf(
		"http://localhost:%v/api/categories/%v",
		os.Getenv("SERVER_PORT"),
		category.Id,
	)
	request := httptest.NewRequest(http.MethodPut, url, requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Key", os.Getenv("API_KEY"))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusBadRequest, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestFindCategoryByIdSuccess(t *testing.T) {
	if err := godotenv.Load("../.env.test"); err != nil {
		panic(err)
	}

	db, err := newDBTester()
	if err != nil {
		panic(err)
	}

	// add a category to db
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	categoryRepository := repository.NewCategoryRepository()
	category, err := categoryRepository.Create(context.Background(), tx, domain.Category{
		Name: "Electronics",
	})
	if err != nil {
		panic(err)
	}
	if err = tx.Commit(); err != nil {
		panic(err)
	}

	router, err := newRouterTester(db)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf(
		"http://localhost:%v/api/categories/%v",
		os.Getenv("SERVER_PORT"),
		category.Id,
	)
	request := httptest.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("X-API-Key", os.Getenv("API_KEY"))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(
		t,
		category.Id,
		int(responseBody["data"].(map[string]any)["id"].(float64)),
	)
	assert.Equal(t, "Electronics", responseBody["data"].(map[string]any)["name"])
}

func TestCategoryByIdNotFound(t *testing.T) {
	if err := godotenv.Load("../.env.test"); err != nil {
		panic(err)
	}

	db, err := newDBTester()
	if err != nil {
		panic(err)
	}

	// add a category to db
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	categoryRepository := repository.NewCategoryRepository()
	_, err = categoryRepository.Create(context.Background(), tx, domain.Category{
		Name: "Electronics",
	})
	if err != nil {
		panic(err)
	}
	if err = tx.Commit(); err != nil {
		panic(err)
	}

	router, err := newRouterTester(db)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf(
		"http://localhost:%v/api/categories/%v",
		os.Getenv("SERVER_PORT"),
		-1,
	)
	request := httptest.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("X-API-Key", os.Getenv("API_KEY"))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	if err := godotenv.Load("../.env.test"); err != nil {
		panic(err)
	}

	db, err := newDBTester()
	if err != nil {
		panic(err)
	}

	// add a category to db
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	categoryRepository := repository.NewCategoryRepository()
	category, err := categoryRepository.Create(context.Background(), tx, domain.Category{
		Name: "Electronics",
	})
	if err != nil {
		panic(err)
	}
	if err = tx.Commit(); err != nil {
		panic(err)
	}

	router, err := newRouterTester(db)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf(
		"http://localhost:%v/api/categories/%v",
		os.Getenv("SERVER_PORT"),
		category.Id,
	)
	request := httptest.NewRequest(http.MethodDelete, url, nil)
	request.Header.Set("X-API-Key", os.Getenv("API_KEY"))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteCategoryNotFound(t *testing.T) {
	if err := godotenv.Load("../.env.test"); err != nil {
		panic(err)
	}

	db, err := newDBTester()
	if err != nil {
		panic(err)
	}

	// add a category to db
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	categoryRepository := repository.NewCategoryRepository()
	_, err = categoryRepository.Create(context.Background(), tx, domain.Category{
		Name: "Electronics",
	})
	if err != nil {
		panic(err)
	}
	if err = tx.Commit(); err != nil {
		panic(err)
	}

	router, err := newRouterTester(db)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf(
		"http://localhost:%v/api/categories/%v",
		os.Getenv("SERVER_PORT"),
		-1,
	)
	request := httptest.NewRequest(http.MethodDelete, url, nil)
	request.Header.Set("X-API-Key", os.Getenv("API_KEY"))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusNotFound, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestGetAllCategories(t *testing.T) {
	if err := godotenv.Load("../.env.test"); err != nil {
		panic(err)
	}

	db, err := newDBTester()
	if err != nil {
		panic(err)
	}

	// add 2 categories to db
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	categoryRepository := repository.NewCategoryRepository()
	category1, err := categoryRepository.Create(context.Background(), tx, domain.Category{
		Name: "Electronics",
	})
	if err != nil {
		panic(err)
	}
	category2, err := categoryRepository.Create(context.Background(), tx, domain.Category{
		Name: "Fashion",
	})
	if err != nil {
		panic(err)
	}
	if err = tx.Commit(); err != nil {
		panic(err)
	}

	router, err := newRouterTester(db)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf(
		"http://localhost:%v/api/categories",
		os.Getenv("SERVER_PORT"),
	)
	request := httptest.NewRequest(http.MethodGet, url, nil)
	request.Header.Set("X-API-Key", os.Getenv("API_KEY"))

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusOK, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	// check all categories
	categories := responseBody["data"].([]any)

	categoryResponse1 := categories[0].(map[string]any)
	categoryResponse2 := categories[1].(map[string]any)

	assert.Equal(t, category1.Id, int(categoryResponse1["id"].(float64)))
	assert.Equal(t, category1.Name, categoryResponse1["name"])
	assert.Equal(t, category2.Id, int(categoryResponse2["id"].(float64)))
	assert.Equal(t, category2.Name, categoryResponse2["name"])

}

func TestUnauthorized(t *testing.T) {
	if err := godotenv.Load("../.env.test"); err != nil {
		panic(err)
	}
	db, err := newDBTester()
	if err != nil {
		panic(err)
	}
	router, err := newRouterTester(db)
	if err != nil {
		panic(err)
	}

	requestBody := strings.NewReader(`{"name": "Electronics"}`)
	url := fmt.Sprintf(
		"http://localhost:%v/api/categories",
		os.Getenv("SERVER_PORT"),
	)
	request := httptest.NewRequest(http.MethodPost, url, requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-Keya", "wrong api key")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusUnauthorized, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, http.StatusUnauthorized, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}
