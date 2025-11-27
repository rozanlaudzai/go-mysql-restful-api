package exception

import (
	"encoding/json"
	"net/http"

	"github.com/rozanlaudzai/go-mysql-restful-api/model/web"
)

func WriteErrorResponse(writer http.ResponseWriter, statusCode int, status string, data string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)

	webResponse := web.WebResponse{
		Code:   statusCode,
		Status: status,
		Data:   data,
	}

	// encode webResponse to json
	if err := json.NewEncoder(writer).Encode(webResponse); err != nil {
		panic(err)
	}
}
