package launch

import (
	"fmt"
	"strings"

	"github.com/rylenko/guide/internal/geocode"
)

const builderInitialCapacity int = 1024

// CommaLocationStringer implements LocationStringer interface, it strings
// location with comma separator between location components.
type CommaLocationStringer struct {
	builder strings.Builder
}

// String representation of the location.
func (stringer *CommaLocationStringer) String(
		location geocode.Location) string {
	// Reset previous location string.
	stringer.builder.Reset()

	// Address components to join with comma.
	components := [...]string{
		location.Country(),
		location.State(),
		location.City(),
		location.Street(),
		location.HouseNumber()}
	// Append address components to the string builder.
	for _, component := range components {
		if component != "" {
			fmt.Fprint(&stringer.builder, component, ", ")
		}
	}

	// Append address point to the string builder.
	point := location.Point()
	fmt.Fprintf(&stringer.builder, "[%f, %f]", point.Lat(), point.Long())

	return stringer.builder.String()
}

// Creates a new instance of comma location stringer.
func NewCommaLocationStringer() *CommaLocationStringer {
	// Create a new instance of string builder and grow its capacity.
	var builder strings.Builder
	builder.Grow(builderInitialCapacity)

	return &CommaLocationStringer{builder: builder}
}

// Ensure that comma location stringer implements location stringer interface.
var _ LocationStringer = (*CommaLocationStringer)(nil)
