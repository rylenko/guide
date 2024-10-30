package gh_globe

import (
	"fmt"
	"strings"

	"github.com/rylenko/guide/internal/globe"
)

// Data transfer object graphhopper's API location representation, implements
// globe's location interface.
type Location struct {
	City string        `json:"city"`
	Country string     `json:"country"`
	HouseNumber string `json:"housenumber"`
	State string       `json:"state"`
	Street string      `json:"street"`
	PointValue Point   `json:"point"`
}

// Point of the location.
func (location *Location) Point() globe.Point {
	return &location.PointValue
}

// String representation of the location.
func (location *Location) String() string {
	var stringBuilder strings.Builder

	// Address components to join with comma.
	components := [...]string{
		location.Country,
		location.State,
		location.City,
		location.Street,
		location.HouseNumber}
	// Append address components to the string builder.
	for _, component := range components {
		if component != "" {
			fmt.Fprint(&stringBuilder, component, ", ")
		}
	}

	// Append address point to the string builder.
	fmt.Fprintf(&stringBuilder, "[%s]", location.PointValue.String())

	return stringBuilder.String()
}

// Ensure that location DTO implements location interface.
var _ globe.Location = (*Location)(nil)
