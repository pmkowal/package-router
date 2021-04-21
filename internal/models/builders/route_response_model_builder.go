package builders

import (
	"packageRouter/internal/models"
	"packageRouter/internal/services"
	"sync"
)

const maxConcurrentRequests = 10

func MakeRoutesResponseModel(requestModel *models.RoutesRequestModel) *models.RoutesResponseModel {
	waitGroup := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	maxChan := make(chan bool, maxConcurrentRequests)
	responseModel := &models.RoutesResponseModel{}
	responseModel.SourceRaw = requestModel.SourceRaw
	for _, destination := range requestModel.Destinations {
		maxChan <- true
		waitGroup.Add(1)
		go services.GetRouteWorker(waitGroup, mutex, maxChan, requestModel.Source, destination, responseModel)
	}
	waitGroup.Wait()
	responseModel.SortRoutesByDurationAndDistance()
	return responseModel
}

