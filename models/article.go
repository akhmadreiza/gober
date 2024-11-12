package models

type Article struct {
	URL       string `json:"url"`
	Title     string `json:"title"`
	ShortDesc string `json:"description"`
	Author    string `json:"author"`
	Date      string `json:"timestamp"`
	SourceUrl string `json:"source_url"`
	Content   string `json:"content"`
}
