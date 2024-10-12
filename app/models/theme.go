package models

type Theme struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	ThemeConfig string `json:"theme_config"`
}
