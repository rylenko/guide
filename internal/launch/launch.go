package launch

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/rylenko/guide/internal/geocode"
	"github.com/rylenko/guide/internal/weather"
)

func Launch(
		geocoder geocode.Geocoder,
		locationStringer LocationStringer,
		weatherFetcher weather.Fetcher,
		weatherStringer WeatherStringer,
		input io.Reader,
		output io.Writer) error {
	// Create standard input reader.
	bufInput := bufio.NewReader(input)

	for {
		// Try to read place input.
		place, err := readPlace(bufInput, output)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return fmt.Errorf("read place input: %w", err)
		}

		// Try to geocode input place.
		locations, err := geocoder.Geocode(place)
		if err != nil {
			return fmt.Errorf("geocode \"%s\": %w", place, err)
		}
		// Prompt another place if no locations found.
		if len(locations) == 0 {
			fmt.Fprintln(output, "Locations not found.\n")
			continue
		}

		// Try to suggest received locations to the user.
		selectedLocation, err := suggestLocations(
			locations, locationStringer, output, bufInput)
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return fmt.Errorf("suggest locations of %s: %w", place, err)
		}

		// Try to get location weather.
		weather, err := weatherFetcher.Fetch(selectedLocation.Point())
		if err != nil {
			return fmt.Errorf("fetch weather at %v: %w", selectedLocation.Point(), err)
		}
		fmt.Fprintln(output, "\nWeather: ", weatherStringer.String(weather), ".\n")

	}

	return nil
}

// Prompts for and reads place input until a non-blank place is entered.
func readPlace(reader *bufio.Reader, output io.Writer) (string, error) {
	for {
		// Print the place prompt to the user.
		fmt.Fprint(output, "Enter place to guide: ")

		// Try to read place to guide from input.
		place, err := reader.ReadString('\n')
		if err != nil {
			return "", fmt.Errorf("read string until '\n': %w", err)
		}

		// Trim newline character from input place and return it if non-blank.
		place = strings.TrimSuffix(place, "\n")
		if place != "" {
			return place, nil
		}
	}
}

// Prints all locations to the output to further select the desired location.
func suggestLocations(
		locations []geocode.Location,
		stringer LocationStringer,
		output io.Writer,
		input *bufio.Reader) (geocode.Location, error) {
	// Suggest locations to select.
	for i, location := range locations {
		fmt.Fprintf(output, "[%d] %s.\n", i, stringer.String(location))
	}
	fmt.Fprintln(output)

	for {
		// Prompt location input.
		fmt.Fprint(output, "Select location using its index: ")

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
			fmt.Fprintln(output, "Invalid index.")
			continue
		}

		return locations[locationIndex], nil
	}
}
