package routes

import (
	"go_ecommerce/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func AuthRoutes(r chi.Router) {
	r.Post("/register", handlers.Register)
	r.Post("/login", handlers.Login)
}
