package tests

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/u2d-man/polyfeed/internal/core"
)

func TestSaveToFile_WritesCorrectJson(t *testing.T) {
	tmpFile := "test_file.json"
	defer os.Remove(tmpFile)

	data := []core.Article{
		{
			Title:     "This is article title.",
			Link:      "https://example.com",
			Analyzed:  "This is article analyzed.",
			Published: "2025-03-23 12:34:56",
		},
	}

	err := core.SaveToFile(data, tmpFile)
	if err != nil {
		t.Fatalf("SaveToFile failed: %v", err)
	}

	// Read and check the written file.
	content, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	var result []core.Article
	err = json.Unmarshal(content, &result)
	if err != nil {
		t.Fatalf("Invalid JSON content: %v", err)
	}

	if len(result) != 1 {
		t.Errorf("Expected 1 article, got %d", len(result))
	}
	if result[0].Title != "This is article title." {
		t.Errorf("Unexpected title: %s", result[0].Title)
	}
}
