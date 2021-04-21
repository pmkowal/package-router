package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type LocationModel struct {
	Latitude float64
	Longitude float64
}

func (m *LocationModel) Parse(location string) error {
	source := strings.Split(location, ",")
	if len(source) < 2 {
		return errors.New("the argument should contain two values separated by a comma")
	}
	m.Latitude, _ = strconv.ParseFloat(source[0], 64)
	m.Longitude, _ = strconv.ParseFloat(source[1], 64)
	return nil
}

func (m *LocationModel) Description() string {
	return fmt.Sprintf("%.6f,%.6f", m.Latitude, m.Longitude)
}

