package core

// Article represents a processed RSS feed article with its metadata and analysis.
// It serves as the core data model for the application.
type Article struct {
	// Title is the headline of the article
	Title string `json:"title"`

	// Link is the URL to the full article
	Link string `json:"link"`

	// Analyzed contains the AI-generated summary of the article content
	Analyzed string `json:"analyzed"`

	// Published stores the formatted publication date and time
	Published string `json:"published"`
}
