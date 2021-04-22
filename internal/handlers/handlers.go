package handlers

import (
	"encoding/json"
	"net/http"
	"packageRouter/internal/models/builders"
	"packageRouter/internal/models/requests"
)

func RoutesHandler(w http.ResponseWriter, r *http.Request) {
	requestModel := &requests.RoutesRequestModel{}
	err := requestModel.Parse(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseModel := builders.MakeRoutesResponseModel(requestModel)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(responseModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
