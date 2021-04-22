package requests

import (
	"errors"
	"fmt"
	"packageRouter/internal/models"
)

const srcKey = "src"
const dstKey = "dst"
const noSourceLocationError = "`src` - source location need to be provided"
const toManySourceLocationsError = "`src` - only one source location can be provided"
const noDestinationLocationError = "`dst` - at least one destination location need to be provided"

type RoutesRequestModel struct {
	SourceRaw       string
	Source          models.LocationModel
	DestinationsRaw []string
	Destinations    []models.LocationModel
}

func (m *RoutesRequestModel) Parse(values map[string][]string) error {
	src := values[srcKey]
	dst := values[dstKey]
	if len(src) == 0 {
		return errors.New(noSourceLocationError)
	}
	if len(src) > 1 {
		return errors.New(toManySourceLocationsError)
	}
	if len(dst) < 1 {
		return errors.New(noDestinationLocationError)
	}
	m.SourceRaw = src[0]
	m.DestinationsRaw = dst
	err := m.Source.Parse(m.SourceRaw)
	if err != nil {
		message := fmt.Sprintf("`src` - %s", err.Error())
		return errors.New(message)
	}
	for index, destination := range m.DestinationsRaw {
		m.Destinations = append(m.Destinations, models.LocationModel{})
		err := m.Destinations[index].Parse(destination)
		if err != nil {
			message := fmt.Sprintf("`dst` - %s", err.Error())
			return errors.New(message)
		}
	}
	return nil
}
