package std_network

import (
	"fmt"
	"net/http"

	"github.com/rylenko/guide/internal/network"
)

// Requester interface implementation, which uses standard HTTP library.
type Requester struct {}

// Sends get request to passed URL using standard HTTP library and returns
// response from it.
func (requester *Requester) Get(url string) (network.Response, error) {
	// Send GET request using standard HTTP library.
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get(\"%s\"): %w", err)
	}

	return NewResponse(response), nil
}

// Creates a new instance of standard requester.
func NewRequester() *Requester {
	return &Requester{}
}

// Ensure that standard requester implements requester interface.
var _ network.Requester = (*Requester)(nil)
