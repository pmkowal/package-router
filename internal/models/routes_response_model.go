package models

type RoutesResponseModel struct {
	SourceRaw string `json:"source"`
	Routes []RouteModel `json:"routes"`
}

func (m *RoutesResponseModel) Make(requestModel RoutesRequestModel) {
	m.SourceRaw = requestModel.SourceRaw
	for index, destination := range requestModel.DestinationsRaw {
		m.Routes = append(m.Routes, RouteModel{})
		m.Routes[index].DestinationRaw = destination
	}
}