package models

type CustomizeHome struct {
	ID       int    `json:"ID"`
	Username string `json:"username"`
	Content  string `json:"content"`
}
