package services

import (
	"encoding/json"
	"errors"
	"log"
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
	maxChan chan bool,
	src models.LocationModel,
	dst models.LocationModel,
	responseModel *models.RoutesResponseModel,
) {
	defer waitGroup.Done()
	defer func(maxChan chan bool) { <-maxChan }(maxChan)
	route, err := getRoute(src, dst)
	if err != nil {
		log.Println(err)
		return
	}
	routeModel := &models.RouteResponseModel{}
	routeModel.DestinationRaw = dst.Description()
	routeModel.SetRoute(route)
	mutex.Lock()
	responseModel.Routes = append(responseModel.Routes, *routeModel)
	mutex.Unlock()
}


func getRoute(src models.LocationModel, dst models.LocationModel) (*models.RouteModel, error) {
	endpointURL, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	locations := strings.Join([]string{src.Description(), dst.Description()}, ";")
	endpointURL.Path = path.Join(endpointURL.Path, locations)
	response, err := http.Get(endpointURL.String())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	responseModel := &models.OSRMResponseModel{}
	err = json.NewDecoder(response.Body).Decode(&responseModel)
	if err != nil {
		return nil, err
	}
	if !responseModel.IsValid() {
		return nil, errors.New("no routes found")
	}
	return &responseModel.Routes[0], nil
}