package main

import (
	"log"
	"os"

	"github.com/rylenko/guide/internal/gh_geocode"
	"github.com/rylenko/guide/internal/launch"
	"github.com/rylenko/guide/internal/std_network"
)

const GHAPIKey string = "baf019de-382e-4910-b1c4-906233282200"

func main() {
	// Create a new instance of standard network requester.
	requester := std_network.NewRequester()

	// Create a new instance of graphhopper geocoder.
	//
	// TODO: Support language parameter.
	geocoder := gh_geocode.NewGeocoder(requester, GHAPIKey)

	// Create a new instance of geocoder location stringer.
	locationStringer := launch.NewCommaLocationStringer()

	// Create a new instance of openweathermap fetcher.
	//
	// TODO: Support language parameter.
	// weatherFetcher := owm_weather.NewFetcher(requester, OWMAPIKey)

	// Launch application using interface instances and standard IO.
	err := launch.Launch(geocoder, locationStringer, os.Stdin, os.Stdout)
	if err != nil {
		log.Fatal("Failed to launch: ", err)
	}
}
