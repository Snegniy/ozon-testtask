package model

type UrlBaseStorage struct {
	ShortURL string `json:"url,omitempty" db:"short_url"`
}

type UrlShortStorage struct {
	BaseURL string `json:"url,omitempty" db:"base_url"`
}
