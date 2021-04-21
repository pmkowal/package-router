package builders

import (
	"packageRouter/internal/models"
	"packageRouter/internal/services"
)

func MakeRoutesResponseModel(requestModel *models.RoutesRequestModel) *models.RoutesResponseModel {
	responseModel := &models.RoutesResponseModel{}
	for _, destination := range requestModel.Destinations {
		route := services.GetRoute(&requestModel.Source, &destination)
		routeModel := &models.RouteResponseModel{}
		routeModel.DestinationRaw = destination.Description()
		routeModel.SetRoute(route)
		responseModel.Routes = append(responseModel.Routes, *routeModel)
	}
	return responseModel
}
