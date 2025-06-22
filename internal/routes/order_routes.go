package routes

import (
	"go_ecommerce/internal/handlers"
	"go_ecommerce/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func OrderRouters(r chi.Router) {
	r.With(middleware.JWTAuth).Get("/", handlers.GetOrders)
	r.With(middleware.JWTAuth).Post("/", handlers.CreateOrder)

}
