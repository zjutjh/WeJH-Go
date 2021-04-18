package models

import "time"

type SchoolBus struct {
	ID         int
	Line       int
	From       string
	To         string
	StartTime  time.Time
	UpdateTime time.Time
}
