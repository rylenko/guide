package launch

import "github.com/rylenko/guide/internal/weather"

// WeatherStringer is an interface to get weather's string representation.
type WeatherStringer interface {
	// String must build string representation of passed weather.
	String(weather weather.Weather) string
}
