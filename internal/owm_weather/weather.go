package owm_weather

import (
	"github.com/rylenko/guide/internal/weather"
)

// Data transfer object of openweathermap's API weather representation,
// implements weather interface.
type Weather struct {
	Headers []struct {
		Type string `json:"main"`
	} `json:"weather"`
	Main struct {
		Humidity uint8 `json:"humidity"`
		Temp float64   `json:"temp"`
	} `json:"main"`
	Wind struct {
		Speed float64 `json:"speed"`
	} `json:"wind"`
}

// Humidity in percents of the weather.
func (weather *Weather) Humidity() uint8 {
	return weather.Main.Humidity
}

// Temperature of the weather.
func (weather *Weather) Temp() float64 {
	return weather.Main.Temp
}

// Type of the weather.
func (weather *Weather) Type() string {
	if len(weather.Headers) == 0 {
		return ""
	}
	return weather.Headers[0].Type
}

// Wind speed of the weather.
func (weather *Weather) WindSpeed() float64 {
	return weather.Wind.Speed
}

// Ensure that weather DTO implements weather interface.
var _ weather.Weather = (*Weather)(nil)
