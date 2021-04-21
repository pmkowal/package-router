package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const locationError = "location should be a pair of two numbers separated by a comma"

type LocationModel struct {
	Latitude float64
	Longitude float64
}

func (m *LocationModel) Parse(location string) error {
	source := strings.Split(location, ",")
	if len(source) < 2 {
		return errors.New(locationError)
	}
	var err error
	m.Latitude, err = strconv.ParseFloat(source[0], 64)
	if err != nil {
		return errors.New(locationError)
	}
	m.Longitude, err = strconv.ParseFloat(source[1], 64)
	if err != nil {
		return errors.New(locationError)
	}
	return nil
}

func (m *LocationModel) Description() string {
	return fmt.Sprintf("%.6f,%.6f", m.Latitude, m.Longitude)
}

