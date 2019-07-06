package utils

import (
	"net/http"

	"github.com/gorilla/handlers"
)

func Cors() func(http.Handler) http.Handler {
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "TRACE", "HEAD"})
	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "X-Requested-With"})
	return handlers.CORS(headers, methods, origins)
}
