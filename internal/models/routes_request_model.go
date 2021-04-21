package models

import "errors"

type RoutesRequestModel struct {
	SourceRaw       string
	Source          LocationModel
	DestinationsRaw []string
	Destinations    []LocationModel
}

func (m *RoutesRequestModel) Parse(values map[string][]string) error {
	src := values["src"]
	dst := values["dst"]
	if len(src) < 1 {
		return errors.New("`src` key should not be empty")
	}
	if len(dst) < 1 {
		return errors.New("`dst` key should have at least one value")
	}
	m.SourceRaw = src[0]
	m.DestinationsRaw = dst
	m.Source.Parse(m.SourceRaw)
	for index, destination := range m.DestinationsRaw {
		m.Destinations = append(m.Destinations, LocationModel{})
		m.Destinations[index].Parse(destination)
	}
	return nil
}