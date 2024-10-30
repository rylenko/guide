package gh_globe

// Data transfer object of graphhopper's API locations representation.
type Locations struct {
	Slice []Location `json:"hits"`
}
