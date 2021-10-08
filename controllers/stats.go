package controllers

import (
	"encoding/json"
	"net/http"

	"example.com/ml_tech/services"
)

func GetStats(w http.ResponseWriter, r *http.Request) {
	stat := services.GetStatsService()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stat)
}
