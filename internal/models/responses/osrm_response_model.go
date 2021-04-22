package responses

import "packageRouter/internal/models"

const validResponseCode = "Ok"

type OSRMResponseModel struct {
	Code   string              `json:"code"`
	Routes []models.RouteModel `json:"routes"`
}

func (m *OSRMResponseModel) IsValid() bool {
	return m.Code == validResponseCode && len(m.Routes) > 0
}
