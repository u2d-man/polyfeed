package core

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func ParseAndFormatTime(raw string) (string, error) {
	t, err := time.Parse(InputFormat, raw)
	if err != nil {
		return "", err
	}
	return t.Format(TimeLayout), nil
}

func SaveToFile(data any, filename string) error {
	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, content, 0644)
}

func GetAPIKey() (string, error) {
	key := os.Getenv(EnvAPIKey)
	if key == "" {
		return "", fmt.Errorf("%s environment variable not set", EnvAPIKey)
	}
	return key, nil
}
