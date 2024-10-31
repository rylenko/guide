package std_network

import (
	"errors"
	"io"
	"net/http"

	"github.com/rylenko/guide/internal/network"
)

// Implementation of response interface that can be instantiated from a
// standard HTTP response.
//
// Please note that the body is a stream and needs to be closed.
type Response struct {
	body io.ReadCloser
	statusCode int
}

// Returns body stream of response.
func (response *Response) Body() io.ReadCloser {
	return response.body
}

// Checks that response contains HTTP error code.
func (response *Response) Error() error {
	if response.statusCode < 400 {
		return nil
	}
	return errors.New(http.StatusText(response.statusCode))
}

// Creates a new instance of response using passed standard HTTP response.
func NewResponse(response *http.Response) *Response {
	return &Response{
		body: response.Body,
		statusCode: response.StatusCode,
	}
}

// Ensure that standard response implements response interface.
var _ network.Response = (*Response)(nil)
