package weather

import "github.com/rylenko/guide/internal/globe"

// Fetcher is interface for weather fetching.
type Fetcher interface {
	// Fetch must fetch weather using passed map point.
	Fetch(point globe.Point) (Weather, error)
}
