package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/hmailyan/go-commerce/internal/models"
)

var products = []models.Product{
	{ID: 1, Name: "Laptop", Price: 999.99},
	{ID: 2, Name: "Mouse", Price: 19.99},
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idStr)

	for _, p := range products {
		if p.ID == id {
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	product.ID = len(products) + 1
	products = append(products, product)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}
