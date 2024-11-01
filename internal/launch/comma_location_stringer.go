package launch

import (
	"fmt"
	"strings"

	"github.com/rylenko/guide/internal/geocode"
)

const builderInitialCap int = 1024

// CommaLocationStringer implements LocationStringer interface, it strings
// location with comma separator between location components.
//
// The zero value is ready to use.
type CommaLocationStringer struct {
	builder strings.Builder
}

// String representation of the location.
func (stringer *CommaLocationStringer) String(
		location geocode.Location) string {
	// Grow to initial capacity to reduce allocations count.
	if stringer.builder.Cap() < builderInitialCap {
		stringer.builder.Grow(builderInitialCap - stringer.builder.Cap())
	}

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

// Ensure that comma location stringer implements location stringer interface.
var _ LocationStringer = (*CommaLocationStringer)(nil)
