package gh_geocode

import (
	"fmt"

	"github.com/rylenko/guide/internal/geocode"
	"github.com/rylenko/guide/internal/globe"
	"github.com/rylenko/guide/internal/network"
)

const endpointURL string = "https://graphhopper.com/api/1/geocode"

// GHGeocoder is implementation of geocoder interface via graphhopper.com API.
type GHGeocoder struct {
	baseURL string
	requester network.Requester
}

// Geocode sends a request to the graphhopper geocoder and receives a response.
func (geocoder *GHGeocoder) Geocode(query string) ([]globe.Location, error) {
	// Sends request to the API endpoint to receive geocoder response.
	response, err := geocoder.sendRequest(query)
	if err != nil {
		return nil, fmt.Errorf("sendRequest(\"%s\"): %v", query, err)
	}

	// TODO: Read N bytes and deserialize to struct. then deserialized struct convert to locations: slice.
	response.Body().Close()
}

// Builds the request URL with passed query.
func (geocoder *GHGeocoder) buildURL(query string) string {
	return fmt.Sprintf("%s?key=api_key&q=%s", endpointURL, geocoder.apikey, query)
}

// Sends request to the API endpoint to receive geocoder response.
func (geocoder *GHGeocoder) sendRequest(
		query string) (network.Response, error) {
	// Build URL with accepted query.
	url := geocoder.buildURL(query)

	// Send get request to built URL and receive a response.
	response, err := geocoder.requester.Get(url)
	if err != nil {
		return nil, fmt.Errorf("requester.Get(\"%s\"): %v", url, err)
	}
	return response, nil
}

// Creates a new instance of graphhopper geocoder using passed apikey.
func NewGHGeocoder(requester network.Requester, apikey string) *GHGeocoder {
	// Format base URL for future requests.
	baseURL := fmt.Sprintf("%s?key=%s", endpointURL, apikey)

	return &GHGeocoder{
		baseURL: baseURL,
		requester: requester,
	}
}

// Ensure that graphhopper geocoder implements geocoder interface.
var _ geocode.Geocoder = (*GHGeocoder)(nil)
