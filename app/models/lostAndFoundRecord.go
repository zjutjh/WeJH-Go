package models

import "time"

type LostAndFoundRecord struct {
	ID          int       `json:"id"`
	Type        bool      `json:"type"`
	Campus      string    `json:"campus"`
	Kind        string    `json:"kind"`
	PublishTime time.Time `json:"publish_time" gorm:"type:timestamp;"`
	IsProcessed bool      `json:"-"`
	Img1        string    `json:"img1"`
	Img2        string    `json:"img2"`
	Img3        string    `json:"img3"`
	Publisher   string    `json:"publisher"`
	Content     string    `json:"content"`
}
