package launch

import (
	"fmt"

	"github.com/rylenko/guide/internal/weather"
)

// CommaWeatherStringer implements WeatherStringer interface, it strings
// weather with comma separator between weather components.
//
// The zero value is ready to use.
type CommaWeatherStringer struct {}

// String representation of the location.
func (stringer *CommaWeatherStringer) String(weather weather.Weather) string {
	return fmt.Sprintf(
		"%s, %f℃, humidity %d%, wind %f m/s",
		weather.Type(),
		weather.Temp(),
		weather.Humidity(),
		weather.WindSpeed())
}

// Ensure that comma weather stringer implements weather stringer interface.
var _ WeatherStringer = (*CommaWeatherStringer)(nil)
