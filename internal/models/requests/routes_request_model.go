package requests

import (
	"errors"
	"fmt"
	"packageRouter/internal/models"
)

type RoutesRequestModel struct {
	SourceRaw       string
	Source          models.LocationModel
	DestinationsRaw []string
	Destinations    []models.LocationModel
}

func (m *RoutesRequestModel) Parse(values map[string][]string) error {
	src := values["src"]
	dst := values["dst"]
	if len(src) == 0 {
		return errors.New("`src` - source location need to be provided")
	}
	if len(src) > 1 {
		return errors.New("`src` - only one source location can be provided")
	}
	if len(dst) < 1 {
		return errors.New("`dst` - at least one destination location need to be provided")
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
