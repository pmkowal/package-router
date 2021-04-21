package handlers

import (
	"encoding/json"
	"net/http"
	"packageRouter/internal/models"
)

func RoutesHandler(w http.ResponseWriter, r *http.Request) {
	responseData := models.RoutesResponseModel{}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseData)
}