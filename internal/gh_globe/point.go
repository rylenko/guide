package gh_globe

import (
	"fmt"

	"github.com/rylenko/guide/internal/globe"
)

// Data transfer object of graphhopper's API point representation, implements
// globe's point interface.
type Point struct {
	Latitude float64  `json:"lat"`
	Longitude float64 `json:"lng"`
}

// Latitude of the point DTO.
func (point *Point) Lat() float64 {
	return point.Latitude
}

// Longitude of the point DTO.
func (point *Point) Long() float64 {
	return point.Longitude
}

// String representation of the point.
func (point *Point) String() string {
	return fmt.Sprint(point.Latitude, ", ", point.Longitude)
}

// Ensure that point DTO implements point interface.
var _ globe.Point = (*Point)(nil)