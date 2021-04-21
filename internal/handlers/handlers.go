package handlers

import (
	"encoding/json"
	"net/http"
	"packageRouter/internal/models"
)

func RoutesHandler(w http.ResponseWriter, r *http.Request) {
	requestModel := models.RoutesRequestModel{}
	err := requestModel.Parse(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseModel := models.RoutesResponseModel{}
	responseModel.Make(requestModel)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseModel)
}