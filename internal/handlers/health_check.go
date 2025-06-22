package handlers

import (
	"go_ecommerce/internal/db"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	sqlDB, err := db.DB.DB()
	if err != nil || sqlDB.Ping() != nil {
		http.Error(w, "Database connection failed", http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
