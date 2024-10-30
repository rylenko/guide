package main

import (
	"log"
	"os"

	"github.com/rylenko/guide/internal/gh_geocode"
	"github.com/rylenko/guide/internal/launch"
	"github.com/rylenko/guide/internal/std_network"
)

const APIKey string = "baf019de-382e-4910-b1c4-906233282200"

func main() {
	geocoder := gh_geocode.NewGeocoder(std_network.NewRequester(), APIKey)

	if err := launch.Launch(geocoder, os.Stdin, os.Stdout); err != nil {
		log.Fatalf("Failed to launch: %v", err)
	}
}
