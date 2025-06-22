package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"go_ecommerce/internal/db"
	"go_ecommerce/internal/handlers"
	"go_ecommerce/internal/routes"
)

func main() {
	db.Init() // <-- podklyuchenie k baze

	r := chi.NewRouter()
	r.Route("/auth", routes.AuthRoutes)
	r.Route("/products", routes.ProductRoutes)

	r.Get("/health", handlers.HealthCheck)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Go E-Commerce API"))
	})

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", r)
}
