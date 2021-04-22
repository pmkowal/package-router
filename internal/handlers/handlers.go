package handlers

import (
	"encoding/json"
	"net/http"
	"packageRouter/internal/models/builders"
	"packageRouter/internal/models/requests"
)

const contentTypeKey = "Content-Type"
const contentTypeValue = "application/json"

func RoutesHandler(w http.ResponseWriter, r *http.Request) {
	requestModel := &requests.RoutesRequestModel{}
	err := requestModel.Parse(r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	responseModel := builders.MakeRoutesResponseModel(requestModel)
	w.Header().Set(contentTypeKey, contentTypeValue)
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(responseModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
