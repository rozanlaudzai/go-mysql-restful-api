package middleware

import (
	"net/http"

	"github.com/rozanlaudzai/go-mysql-restful-api/exception"
)

type AuthMiddleware struct {
	Handler       http.Handler
	CorrectAPIKey string
}

func NewAuthMiddleware(handler http.Handler, correctAPIkey string) *AuthMiddleware {
	return &AuthMiddleware{
		Handler:       handler,
		CorrectAPIKey: correctAPIkey,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-API-Key") == middleware.CorrectAPIKey {
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		exception.WriteErrorResponse(writer, http.StatusUnauthorized, "UNAUTHORIZED", "")
	}
}
