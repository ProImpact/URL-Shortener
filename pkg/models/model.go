package models

type URL struct {
	Original  string `json:"original"`
	Shortened string `json:"shortened"`
	Clicks    int64  `json:"clicks"`
}
