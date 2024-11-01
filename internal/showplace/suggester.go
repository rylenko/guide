package showplace

import "github.com/rylenko/guide/internal/globe"

// Suggester is an interface that provides showplace suggestions.
type Suggester interface {
	// Suggest must provide showplace suggestions located near a passed point.
	Suggest(point globe.Point) []Showplace
}
