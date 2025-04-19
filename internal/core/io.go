package core

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// ParseAndFormatTime converts a time string from RSS feed format to the application's
// standard time format. It supports both RFC1123 and RFC3339 formats.
//
// Parameters:
//   - raw: The time string to parse, typically from RSS feed's published date
//
// Returns:
//   - The formatted time string in the application's standard format (YYYY-MM-DD HH:MM:SS)
//   - An error if the time string cannot be parsed in any supported format
func ParseAndFormatTime(raw string) (string, error) {
	t, err := time.Parse(InputFormat, raw)
	if err != nil {
		t, err = time.Parse(InputFormatISO, raw)
		if err != nil {
			return "", fmt.Errorf("invalid time format (tried RFC3339): %w", err)
		}
	}
	return t.Format(TimeLayout), nil
}

// SaveToFile serializes the provided data as JSON and writes it to the specified file.
// The file will be created if it does not exist, or overwritten if it does.
//
// Parameters:
//   - data: The data structure to serialize to JSON
//   - filename: The name of the file to write the JSON data to
//
// Returns:
//   - An error if JSON serialization fails or if the file cannot be written
func SaveToFile(data any, filename string) error {
	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, content, 0644)
}

// GetAPIKey retrieves the OpenAI API key from the environment.
//
// Returns:
//   - The API key as a string
//   - An error if the API key environment variable is not set
func GetAPIKey() (string, error) {
	key := os.Getenv(EnvAPIKey)
	if key == "" {
		return "", fmt.Errorf("%s environment variable not set", EnvAPIKey)
	}
	return key, nil
}
