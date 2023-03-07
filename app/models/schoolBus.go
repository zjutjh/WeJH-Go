package models

import (
	"time"
)

type SchoolBus struct {
	ID          int           `json:"id"`
	Line        string        `json:"line"`
	Departure   string        `json:"departure"`
	Destination string        `json:"destination"`
	Type        SchoolBusType `json:"type"`
	StartTime   string        `json:"startTime"`
	UpdatedAt   time.Time     `json:"-" gorm:"type:timestamp;"`
}

type SchoolBusType int

const (
	Weekday SchoolBusType = 1
	Weekend SchoolBusType = 2
)
