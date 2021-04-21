package services

import (
	"encoding/json"
	"net/http"
	"net/url"
	"packageRouter/internal/models"
	"path"
	"strings"
)

const endpoint = "https://router.project-osrm.org/route/v1/driving?overview=false"

func GetRoute(src *models.LocationModel, dst *models.LocationModel) *models.RouteModel {
	route := &models.RouteModel{}
	endpointURL, err := url.Parse(endpoint)
	if err != nil {
		return route
	}
	locations := strings.Join([]string{src.Description(), dst.Description()}, ";")
	endpointURL.Path = path.Join(endpointURL.Path, locations)
	response, err := http.Get(endpointURL.String())
	if err != nil {
		return route
	}
	defer response.Body.Close()
	responseModel := &models.OSRMResponseModel{}
	err = json.NewDecoder(response.Body).Decode(&responseModel)
	if err != nil {
		return route
	}
	if len(responseModel.Routes) < 1 {
		return route
	}
	return &responseModel.Routes[0]
}