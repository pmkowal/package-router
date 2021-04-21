package models

type OSRMResponseModel struct {
	Code   string       `json:"code"`
	Routes []RouteModel `json:"routes"`
}

func (m *OSRMResponseModel) IsValid() bool {
	return m.Code == "Ok" && len(m.Routes) > 0
}