package network

import "io"

// Response after network request.
type Response interface {
	// Body stream of the response.
	Body() io.ReadCloser

	// Check that response containt error.
	Error() error
}
