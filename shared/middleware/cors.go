package middleware

import (
	"net/http"

	"github.com/dungvan2512/soccer-social-network-api/infrastructure"
	"github.com/go-chi/cors"
)

// CORS middleware
func CORS(logger *infrastructure.Logger) func(http.Handler) http.Handler {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Use this to allow specific origin hosts
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
	})
	return cors.Handler
}
