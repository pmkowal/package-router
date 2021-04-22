package builders

import (
	"packageRouter/internal/models/requests"
	"packageRouter/internal/models/responses"
	"packageRouter/internal/services"
	"sync"
)

const maxConcurrentRequests = 10

func MakeRoutesResponseModel(requestModel *requests.RoutesRequestModel) *responses.RoutesResponseModel {
	waitGroup := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	maxChan := make(chan bool, maxConcurrentRequests)
	responseModel := &responses.RoutesResponseModel{}
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
