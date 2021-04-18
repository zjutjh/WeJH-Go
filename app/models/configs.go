package models

import "time"

type Configs struct {
	ID         int
	Key        string
	Value      string
	CreateTime time.Time `gorm:"comment:'发布时间';type:timestamp;"`
}
