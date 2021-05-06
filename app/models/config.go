package models

import "time"

type Config struct {
	ID         int
	Key        string
	Value      string
	UpdateTime time.Time `gorm:"comment:'设置时间';type:timestamp;"`
}
