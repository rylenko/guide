package globe

// Point is an interface for storing the latitude and longitude of a given
// point.
type Point interface {
	// Latitude of the point.
	Lat() float64

	// Longitude of the point.
	Long() float64
}
