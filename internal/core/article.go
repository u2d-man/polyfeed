package core

type Article struct {
	Title     string `json:"title"`
	Link      string `json:"link"`
	Analyzed  string `json:"analyzed"`
	Published string `json:"published"`
}
