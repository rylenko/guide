package network

import "io"

// Response after network request.
type Response interface {
	// Status code of the response.
	StatusCode() int

	// Body stream of the response.
	Body() io.ReadCloser
}
