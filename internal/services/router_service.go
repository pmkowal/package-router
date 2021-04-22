package services

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"packageRouter/internal/models"
	"packageRouter/internal/models/requests"
	"packageRouter/internal/models/responses"
	"sync"
)

const OSRMEndpoint = "https://router.project-osrm.org"

func GetRouteWorker(
	waitGroup *sync.WaitGroup,
	mutex *sync.Mutex,
	maxChan chan bool,
	src models.LocationModel,
	dst models.LocationModel,
	responseModel *responses.RoutesResponseModel,
) {
	defer waitGroup.Done()
	defer func(maxChan chan bool) { <-maxChan }(maxChan)
	route, err := getRouteFromOSRM(src, dst)
	if err != nil {
		log.Println(err)
		return
	}
	routeModel := &responses.RouteResponseModel{}
	routeModel.DestinationRaw = dst.Description()
	routeModel.SetRoute(route)
	mutex.Lock()
	responseModel.Routes = append(responseModel.Routes, *routeModel)
	mutex.Unlock()
}


func getRouteFromOSRM(src models.LocationModel, dst models.LocationModel) (*models.RouteModel, error) {
	requestModel := &requests.OSRMRequestModel{
		Source: src,
		Destination: dst,
	}
	url, err := requestModel.URL(OSRMEndpoint)
	if err != nil {
		return nil, err
	}
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	responseModel := &responses.OSRMResponseModel{}
	err = json.NewDecoder(response.Body).Decode(&responseModel)
	if err != nil {
		return nil, err
	}
	if !responseModel.IsValid() {
		return nil, errors.New("no routes found")
	}
	return &responseModel.Routes[0], nil
}
