package exception

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err any) {

	switch errAssert := err.(type) {
	case NotFoundError:
		WriteErrorResponse(writer, http.StatusNotFound, "NOT FOUND", errAssert.Error()) // data message is always safe because it is my creation
	case validator.ValidationErrors:
		WriteErrorResponse(writer, http.StatusBadRequest, "BAD REQUEST", "invalid fields")
	default:
		WriteErrorResponse(writer, http.StatusInternalServerError, "INTERNAL SERVER ERROR", "internal server error")
	}

}
