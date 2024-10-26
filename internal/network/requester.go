package network

// Request interface for sending requests over the network.
type Requester interface {
	// Sends get request to passed URL and returns response from it.
	Get(url string) (Response, error)
}
