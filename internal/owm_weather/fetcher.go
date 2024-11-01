package owm_weather

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/rylenko/guide/internal/globe"
	"github.com/rylenko/guide/internal/network"
	"github.com/rylenko/guide/internal/weather"
)

const (
	endpointURL string = "https://api.openweathermap.org/data/2.5/weather"
	unitsQueryParamValue string = "metric"
)

// Fetcher is implementation of fetcher interface via openweathermap API.
type Fetcher struct {
	apiKey string
	requester network.Requester
}

// Fetch sends a request to the openweathermap API and receives a response.
func (fetcher *Fetcher) Fetch(point globe.Point) (weather.Weather, error) {
	// Sends request to the API endpoint to receive weather response.
	response, err := fetcher.sendRequest(point)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	defer response.Body().Close()

	// Check that response contain errors.
	if err := response.Error(); err != nil {
		return nil, fmt.Errorf("response error: %w", err)
	}

	// Create new decoder to decode response's JSON.
	decoder := json.NewDecoder(response.Body())

	// Decode JSON response body into DTO.
	var weather Weather
	if err := decoder.Decode(&weather); err != nil {
		return nil, fmt.Errorf("json decode: %w", err)
	}

	return &weather, nil
}

// Builds the request URL with passed query.
func (fetcher *Fetcher) buildURL(point globe.Point) (string, error) {
	// Try to parse base endpoint URL.
	u, err := url.Parse(endpointURL)
	if err != nil {
		return "", fmt.Errorf("parse endpoint URL %s: %w", endpointURL, err)
	}

	// Add API key and query parameter values to the parsed URL.
	values := url.Values{}
	values.Add("appid", fetcher.apiKey)
	values.Add("lat", fmt.Sprint(point.Lat()))
	values.Add("lon", fmt.Sprint(point.Long()))
	values.Add("units", unitsQueryParamValue)
	u.RawQuery = values.Encode()

	return u.String(), nil
}

// Sends request to the API endpoint to receive weather response.
func (fetcher *Fetcher) sendRequest(
		point globe.Point) (network.Response, error) {
	// Build URL with accepted point.
	url, err := fetcher.buildURL(point)
	if err != nil {
		return nil, fmt.Errorf("build URL: %w", err)
	}

	// Send get request to the built URL and receive a response.
	response, err := fetcher.requester.Get(url)
	if err != nil {
		return nil, fmt.Errorf("send get request to %s: %w", url, err)
	}
	return response, nil
}

// Creates a new instance of openweathermap fetcher.
func NewFetcher(requester network.Requester, apiKey string) *Fetcher {
	return &Fetcher{
		apiKey: apiKey,
		requester: requester,
	}
}

// Ensure that openweathermap fetcher implements fetcher interface.
var _ weather.Fetcher = (*Fetcher)(nil)
