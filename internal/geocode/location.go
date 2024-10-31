package geocode

import "github.com/rylenko/guide/internal/globe"

// Location is an interface for storing data about a location on a map:
// point and address components.
type Location interface {
	// City of the location on the map.
	City() string

	// Country of the location on the map.
	Country() string

	// House number of the location on the map.
	HouseNumber() string

	// Point of the location on the map.
	Point() globe.Point

	// State of the location on the map.
	State() string

	// Street of the location on the map.
	Street() string
}
