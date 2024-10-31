package launch

import "github.com/rylenko/guide/internal/geocode"

// LocationStringer is an interface to get location string representation.
type LocationStringer interface {
	// String must build string representation of passed location.
	String(location geocode.Location) string
}
