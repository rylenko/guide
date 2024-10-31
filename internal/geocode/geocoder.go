package geocode

// Geocoder is interface for direct geocoding of raw text into coordinates and
// address components.
type Geocoder interface {
	// Geocode must geocode query string into coordinates and address components.
	Geocode(query string) ([]Location, error)
}
