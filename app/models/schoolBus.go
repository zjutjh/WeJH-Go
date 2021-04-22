package models

import "time"

type SchoolBus struct {
	ID         int       `json:"id"`
	Line       int       `json:"line"`
	From       string    `json:"from"`
	To         string    `json:"to"`
	StartTime  time.Time `json:"startTime"`
	UpdateTime time.Time `json:"updateTime"`
}
