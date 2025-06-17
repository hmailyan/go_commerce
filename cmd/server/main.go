package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hmailyan/go-commerce/internal/handlers"
)

func main() {
	r := chi.NewRouter()

	r.Get("/health", handlers.HealthCheck)
	r.Get("/products", handlers.GetProducts)
	r.Get("/products/{id}", handlers.GetProductByID)
	r.Post("/products", handlers.CreateProduct)

	fmt.Println("🚀 Server started at :8080")
	http.ListenAndServe(":8080", r)
}
