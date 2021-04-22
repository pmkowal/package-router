package requests

import (
	"net/url"
	"packageRouter/internal/models"
	"path"
	"strings"
)

const OSRMPath = "/route/v1/driving"
const OSRMQuery = "overview=false"

type OSRMRequestModel struct {
	Source models.LocationModel
	Destination models.LocationModel
}

func (m *OSRMRequestModel) urlPath() string {
	locations := strings.Join([]string{m.Source.Description(), m.Destination.Description()}, ";")
	return path.Join(OSRMPath, locations)
}

func (m *OSRMRequestModel) urlQuery() string {
	return OSRMQuery
}

func (m *OSRMRequestModel) URL(endpoint string) (string, error) {
	endpointURL, err := url.Parse(endpoint)
	if err != nil {
		return "", err
	}
	endpointURL.Path = m.urlPath()
	endpointURL.RawQuery = m.urlQuery()
	return endpointURL.String(), nil
}