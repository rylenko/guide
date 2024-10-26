package globe

// Location is an interface for storing data about a location on a map:
// point and address components.
type Location interface {
	// Point of the location on the map.
	Point() Point

	// Country of the location.
	Country() string

	// City of the location.
	City() string

	// State of the location.
	State() string

	// Street of the location.
	Street() string

	// House number of the location.
	HouseNumber() string
}
