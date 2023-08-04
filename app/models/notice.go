package models

import "time"

type Notice struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	PublishTime time.Time `json:"publish_time" gorm:"type:timestamp;"`
	Img1        string    `json:"img1"`
	Img2        string    `json:"img2"`
	Img3        string    `json:"img3"`
	Publisher   string    `json:"publisher"`
	Content     string    `json:"content"`
}
