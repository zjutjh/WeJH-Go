package models

import "time"

type Announcement struct {
	ID          int
	Title       string
	Content     string
	PublishTime time.Time `gorm:"comment:'发布时间';type:timestamp;"`
}
