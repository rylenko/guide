package gh_geocode

// Data transfer object of graphhopper's API locations representation.
type Locations struct {
	Slice []Location `json:"hits"`
}
