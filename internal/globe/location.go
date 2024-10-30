package globe

import "fmt"

// Location is an interface for storing data about a location on a map:
// point and a string representation of address components.
type Location interface {
	fmt.Stringer

	// Point of the location on the map.
	Point() Point
}
