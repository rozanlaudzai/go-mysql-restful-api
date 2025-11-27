package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/rozanlaudzai/go-mysql-restful-api/model/web"
	"github.com/rozanlaudzai/go-mysql-restful-api/service"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	// decode json to CategoryCreateRequest
	decoder := json.NewDecoder(request.Body)
	categoryCreateRequest := web.CategoryCreateRequest{}
	err := decoder.Decode(&categoryCreateRequest)
	if err != nil {
		panic(err)
	}

	categoryResponse, err := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	if err != nil {
		panic(err)
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	writer.Header().Set("Content-Type", "application/json")
	// encode webResponse to json
	if err := json.NewEncoder(writer).Encode(webResponse); err != nil {
		panic(err)
	}
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	// decode json to CategoryUpdateRequest
	decoder := json.NewDecoder(request.Body)
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	err := decoder.Decode(&categoryUpdateRequest)
	if err != nil {
		panic(err)
	}

	// get the category id
	categoryIdString := params.ByName("categoryId")
	categoryId, err := strconv.Atoi(categoryIdString)
	if err != nil {
		panic(err)
	}

	categoryUpdateRequest.Id = categoryId

	categoryResponse, err := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	if err != nil {
		panic(err)
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	writer.Header().Set("Content-Type", "application/json")
	// encode webResponse to json
	if err := json.NewEncoder(writer).Encode(webResponse); err != nil {
		panic(err)
	}
}

func (controller *CategoryControllerImpl) DeleteById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	// get the category id
	categoryIdString := params.ByName("categoryId")
	categoryId, err := strconv.Atoi(categoryIdString)
	if err != nil {
		panic(err)
	}

	err = controller.CategoryService.DeleteById(request.Context(), categoryId)
	if err != nil {
		panic(err)
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	writer.Header().Set("Content-Type", "application/json")
	// encode webResponse to json
	if err := json.NewEncoder(writer).Encode(webResponse); err != nil {
		panic(err)
	}
}

func (controller *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	// get the category id
	categoryIdString := params.ByName("categoryId")
	categoryId, err := strconv.Atoi(categoryIdString)
	if err != nil {
		panic(err)
	}

	categoryResponse, err := controller.CategoryService.FindById(request.Context(), categoryId)
	if err != nil {
		panic(err)
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponse,
	}

	writer.Header().Set("Content-Type", "application/json")
	// encode webResponse to json
	if err := json.NewEncoder(writer).Encode(webResponse); err != nil {
		panic(err)
	}
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryResponses, err := controller.CategoryService.FindAll(request.Context())
	if err != nil {
		panic(err)
	}
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categoryResponses,
	}

	writer.Header().Set("Content-Type", "application/json")
	// encode webResponse to json
	if err := json.NewEncoder(writer).Encode(webResponse); err != nil {
		panic(err)
	}
}
