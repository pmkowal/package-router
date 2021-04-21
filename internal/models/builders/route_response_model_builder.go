package builders

import (
	"packageRouter/internal/models"
	"packageRouter/internal/services"
	"sync"
)

func MakeRoutesResponseModel(requestModel *models.RoutesRequestModel) *models.RoutesResponseModel {
	waitGroup := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	responseModel := &models.RoutesResponseModel{}
	for _, destination := range requestModel.Destinations {
		waitGroup.Add(1)
		go services.GetRouteWorker(waitGroup, mutex, requestModel.Source, destination, responseModel)
	}
	waitGroup.Wait()
	return responseModel
}
