package models

import "time"

type Announcement struct {
	ID          int       `json:"id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Content     string    `json:"content"`
	PublishTime time.Time `json:"publishTime" gorm:"comment:'发布时间';type:timestamp;"`
}
