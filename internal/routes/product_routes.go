package routes

import (
	"go_ecommerce/internal/handlers"
	"go_ecommerce/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func ProductRoutes(r chi.Router) {
	r.With(middleware.JWTAuth).Get("/", handlers.GetProducts)
	r.With(middleware.JWTAuth).Post("/", handlers.CreateProduct)
	r.With(middleware.JWTAuth).Get("/{id}", handlers.GetProduct)
	r.With(middleware.JWTAuth).Put("/{id}", handlers.UpdateProduct)
	r.With(middleware.JWTAuth).Delete("/{id}", handlers.DeleteProduct)
}
