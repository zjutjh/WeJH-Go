package models

import (
	"time"
)

type SchoolBus struct {
	ID        int           `json:"id"`
	Line      string        `json:"line"`
	From      string        `json:"from"`
	To        string        `json:"to"`
	Type      SchoolBusType `json:"-"`
	StartTime time.Time     `json:"startTime" gorm:"type:timestamp;"`
	UpdatedAt time.Time     `json:"updateTime" gorm:"type:timestamp;"`
}

type SchoolBusType int

const (
	Weekday SchoolBusType = 0
	Weekend SchoolBusType = 1
)
