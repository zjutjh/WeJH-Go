package models

import "time"

type QA struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	PublishTime time.Time `json:"publish_time" gorm:"type:timestamp;"`
	Publisher   string    `json:"publisher"`
}
