package main

// Story represents a Choose Your Own Adventure story
// It consists of a collection of chapters that contain links
// to other chapters. The story starts at the chapter called "intro"
type Story map[string]Chapter

// Chapter represents a chapter within a Choose Your Own Adventure story
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text    string `json:"text"`
		Chapter string `json:"arc"`
	} `json:"options"`
}
