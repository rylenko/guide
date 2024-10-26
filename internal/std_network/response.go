package network

import (
	"io"

	"github.com/rylenko/guide/internal/network"
)

// Implementation of response interface that can be instantiated from a
// standard HTTP response.
//
// Please note that the body is a stream and needs to be closed.
type Response struct {
	statusCode int
	body io.ReadCloser
}

// Returns body stream of response.
func (response *Response) Body() io.ReadCloser {
	return response.body
}

// Returns the status code of response.
func (response *Response) StatusCode() int {
	return response.statusCode
}

// Creates a new instance of response using passed standard HTTP response.
func NewResponse(response: *http.Response) *Response {
	return &Response{
		statusCode: response.StatusCode,
		body: response.Body,
	}
}

// Ensure that standard response implements response interface.
var _ network.Response = (*Response)(nil)
