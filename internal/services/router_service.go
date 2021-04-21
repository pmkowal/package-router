package services

import (
	"encoding/json"
	"net/http"
	"net/url"
	"packageRouter/internal/models"
	"path"
	"strings"
	"sync"
)

const endpoint = "https://router.project-osrm.org/route/v1/driving?overview=false"

func GetRouteWorker(
	waitGroup *sync.WaitGroup,
	mutex *sync.Mutex,
	src models.LocationModel,
	dst models.LocationModel,
	responseModel *models.RoutesResponseModel,
) {
	defer waitGroup.Done()
	route := getRoute(src, dst)
	routeModel := &models.RouteResponseModel{}
	routeModel.DestinationRaw = dst.Description()
	routeModel.SetRoute(route)
	mutex.Lock()
	responseModel.Routes = append(responseModel.Routes, *routeModel)
	mutex.Unlock()
}


func getRoute(src models.LocationModel, dst models.LocationModel) *models.RouteModel {
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