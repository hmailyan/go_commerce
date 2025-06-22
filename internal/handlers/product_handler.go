package handlers

import (
	"encoding/json"
	"net/http"

	"go_ecommerce/internal/db"
	"go_ecommerce/internal/models"

	"github.com/go-chi/chi/v5"
)

// Get all products
func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	db.DB.Find(&products)
	json.NewEncoder(w).Encode(products)
}

// Get single product by ID
func GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(product)
}

// Create product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.DB.Create(&product)
	json.NewEncoder(w).Encode(product)
}

// Update product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	var updated models.Product
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product.Name = updated.Name
	product.Description = updated.Description
	product.Price = updated.Price

	db.DB.Save(&product)
	json.NewEncoder(w).Encode(product)
}

// Delete product
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var product models.Product
	if err := db.DB.First(&product, id).Error; err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}
	db.DB.Delete(&product)
	w.WriteHeader(http.StatusNoContent)
}
