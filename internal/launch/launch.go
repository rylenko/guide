package launch

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/rylenko/guide/internal/geocode"
	"github.com/rylenko/guide/internal/globe"
	"github.com/rylenko/guide/internal/weather"
)

const placeInputSeparator string = "------------------------------------------"

// TODO: split into small functions
func Launch(
		geocoder geocode.Geocoder,
		locationStringer LocationStringer,
		weatherFetcher weather.Fetcher,
		weatherStringer WeatherStringer,
		input io.Reader,
		output io.Writer) error {
	// Create standard input reader to read input efficiently.
	bufInput := bufio.NewReader(input)

	for {
		// Try to read place input.
		locations, err := readPlaceAndGeocode(bufInput, output, geocoder)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return fmt.Errorf("read place and geocode: %w", err)
		}
		// Prompt another place if no locations found.
		if len(locations) == 0 {
			if _, err := fmt.Fprintln(output, "Locations not found.\n"); err != nil {
				return fmt.Errorf("print locations not found: %w", err)
			}
			continue
		}

		// Try to suggest received locations to the user.
		selectedLocation, err := suggestLocations(
			locations, locationStringer, output, bufInput)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return fmt.Errorf("suggest locations: %w", err)
		}

		// Try to fetch and print weather in selected location.
		err = fetchWeather(
			weatherFetcher, selectedLocation.Point(), weatherStringer, output)
		if err != nil {
			return fmt.Errorf(
				"fetch weather for %s: %w", locationStringer.String(selectedLocation), err)
		}

		// Print separator between places for readability.
		if _, err := fmt.Printf("\n%s\n\n", placeInputSeparator); err != nil {
			return fmt.Errorf("print place input separator: %w", err)
		}
	}

	return nil
}

// Fetches and prints weather in the passed point.
func fetchWeather(
		fetcher weather.Fetcher,
		point globe.Point,
		stringer WeatherStringer,
		output io.Writer) error {
	// Try to fetch weather using accepted point.
	weather, err := fetcher.Fetch(point)
	if err != nil {
		return fmt.Errorf("fetch weather at %v: %w", point, err)
	}

	// Try to print weather string to the output.
	weatherString := stringer.String(weather)
	_, err = fmt.Fprintf(output, "Weather: %s.\n", weatherString)
	if err != nil {
		return fmt.Errorf("print weather %s: %w", weatherString, err)
	}

	return nil
}

// Prompts for and reads place input until a non-blank place is entered.
func readPlaceAndGeocode(
		reader *bufio.Reader,
		output io.Writer,
		geocoder geocode.Geocoder) ([]geocode.Location, error) {
	for {
		// Print the place prompt to the user.
		if _, err := fmt.Fprint(output, ">>> Enter place to guide: "); err != nil {
			return nil, fmt.Errorf("print place prompt: %w", err)
		}

		// Try to read place to guide from input.
		place, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("read string until '\n': %w", err)
		}
		// ReadString reads leaves newline at the end so we need to trim it.
		place = strings.TrimSuffix(place, "\n")
		if place == "" {
			continue
		}

		// Try to geocode input place.
		locations, err := geocoder.Geocode(place)
		if err != nil {
			return nil, fmt.Errorf("geocode \"%s\": %w", place, err)
		}
		return locations, nil
	}
}

// Prints all locations to the output to further select the desired location.
func suggestLocations(
		locations []geocode.Location,
		stringer LocationStringer,
		output io.Writer,
		input *bufio.Reader) (geocode.Location, error) {
	// Suggest locations to select.
	if _, err := fmt.Fprintln(output, "\nSuggestions:"); err != nil {
		return nil, fmt.Errorf("print suggestions header: %w", err)
	}
	for i, location := range locations {
		locationString := stringer.String(location)

		_, err := fmt.Fprintf(output, "[%d] %s.\n", i, locationString)
		if err != nil {
			return nil, fmt.Errorf("print location %s: %w", locationString, err)
		}
	}
	if _, err := fmt.Fprintln(output); err != nil {
		return nil, fmt.Errorf("print newline after suggestions: %w", err)
	}

	for {
		// Prompt location input.
		_, err := fmt.Fprint(output, ">>> Select location using its index: ")
		if err != nil {
			return nil, fmt.Errorf("print location selection prompt: %w", err)
		}

		// Read selected location index as string.
		locationIndexStr, err := input.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("read string until '\n': %w", err)
		}
		// Trim newline character from index input.
		locationIndexStr = strings.TrimSuffix(locationIndexStr, "\n")

		// Try to convert input line to location index integer.
		locationIndex, err := strconv.Atoi(locationIndexStr)
		if err != nil || locationIndex < 0 || len(locations) <= locationIndex {
			if _, err := fmt.Fprintln(output, "Invalid location index."); err != nil {
				return nil, fmt.Errorf("print invalid location index: %w", err)
			}

			continue
		}

		selectedLocation := locations[locationIndex]

		// Print selected location to the user.
		selectedLocationString := stringer.String(selectedLocation)
		_, err = fmt.Fprintf(
			output, "\nSelected location: %s.\n", selectedLocationString)
		if err != nil {
			return nil, fmt.Errorf(
				"print selected location %s: %w", selectedLocationString, err)
		}

		return selectedLocation, nil
	}
}
