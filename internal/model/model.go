package model

type UrlStorage struct {
	BaseURL  string `json:"baseurl,omitempty" db:"base_url"`
	ShortURL string `json:"shorturl,omitempty" db:"short_url"`
}
