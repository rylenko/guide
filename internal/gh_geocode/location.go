package gh_geocode

import (
	"github.com/rylenko/guide/internal/geocode"
	"github.com/rylenko/guide/internal/gh_globe"
	"github.com/rylenko/guide/internal/globe"
)

// Data transfer object of graphhopper's API location representation,
// implements geocode location interface.
type Location struct {
	CityValue string          `json:"city"`
	CountryValue string       `json:"country"`
	HouseNumberValue string   `json:"housenumber"`
	StateValue string         `json:"state"`
	StreetValue string        `json:"street"`
	PointValue gh_globe.Point `json:"point"`
}


// City of the location.
func (location *Location) City() string {
	return location.CityValue
}

// City of the location.
func (location *Location) Country() string {
	return location.CountryValue
}

// House number of the location.
func (location *Location) HouseNumber() string {
	return location.HouseNumberValue
}

// Point of the location.
func (location *Location) Point() globe.Point {
	return &location.PointValue
}

// State of the location.
func (location *Location) State() string {
	return location.StateValue
}

// State of the location.
func (location *Location) Street() string {
	return location.StreetValue
}

// Ensure that location DTO implements location interface.
var _ geocode.Location = (*Location)(nil)
