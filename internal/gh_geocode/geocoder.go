package gh_geocode

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/rylenko/guide/internal/geocode"
	"github.com/rylenko/guide/internal/gh_globe"
	"github.com/rylenko/guide/internal/globe"
	"github.com/rylenko/guide/internal/network"
)

const endpointURL string = "https://graphhopper.com/api/1/geocode"

// Geocoder is implementation of geocoder interface via graphhopper.com API.
type Geocoder struct {
	apiKey string
	requester network.Requester
}

// Geocode sends a request to the graphhopper geocoder and receives a response.
func (geocoder *Geocoder) Geocode(query string) ([]globe.Location, error) {
	// Sends request to the API endpoint to receive geocoder response.
	response, err := geocoder.sendRequest(query)
	if err != nil {
		return nil, fmt.Errorf("sendRequest(\"%s\"): %w", query, err)
	}
	defer response.Body().Close()

	// Check that response contain errors.
	if err := response.Error(); err != nil {
		return nil, fmt.Errorf("response error with query \"%s\": %w", query, err)
	}

	// Create new decoder to decode response's JSON.
	decoder := json.NewDecoder(response.Body())

	// Decode JSON response body into DTO.
	var locations gh_globe.Locations
	if err := decoder.Decode(&locations); err != nil {
		return nil, fmt.Errorf("json Decode(): %w", err)
	}

	// Convert slice of location DTO to location interface.
	//
	// TODO: optimization.
	locationInterfaces := make([]globe.Location, len(locations.Slice))
	for i, location := range locations.Slice {
		locationInterfaces[i] = &location
	}

	return locationInterfaces, nil
}

// Builds the request URL with passed query.
func (geocoder *Geocoder) buildURL(query string) (string, error) {
	// Try to parse base endpoint URL.
	u, err := url.Parse(endpointURL)
	if err != nil {
		return "", fmt.Errorf("parse endpoint URL %s: %w", endpointURL, err)
	}

	// Add API key and query parameter values to the parsed URL.
	values := url.Values{}
	values.Add("key", geocoder.apiKey)
	values.Add("q", query)
	u.RawQuery = values.Encode()

	return u.String(), nil
}

// Sends request to the API endpoint to receive geocoder response.
func (geocoder *Geocoder) sendRequest(
		query string) (network.Response, error) {
	// Build URL with accepted query.
	url, err := geocoder.buildURL(query)
	if err != nil {
		return nil, fmt.Errorf("buildURL(\"%s\"): %w", query, err)
	}

	// Send get request to built URL and receive a response.
	response, err := geocoder.requester.Get(url)
	if err != nil {
		return nil, fmt.Errorf("requester.Get(\"%s\"): %w", url, err)
	}
	return response, nil
}

// Creates a new instance of graphhopper geocoder using passed apikey.
func NewGeocoder(requester network.Requester, apiKey string) *Geocoder {
	return &Geocoder{
		apiKey: apiKey,
		requester: requester,
	}
}

// Ensure that graphhopper geocoder implements geocoder interface.
var _ geocode.Geocoder = (*Geocoder)(nil)
