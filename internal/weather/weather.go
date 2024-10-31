package weather

// Weather is an interface for storing weather data.
type Weather interface {
	// Humidity must return humidity in percents.
	Humidity() uint8

	// Temp must return current temperature.
	Temp() float64

	// Type must return weather type: rain, snow, clouds, etc.
	Type() string

	// WindSpeed must return wind speed per unit of time.
	WindSpeed() float64
}
