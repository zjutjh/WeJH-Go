package models

import "time"

type SchoolBus struct {
	ID        int64
	Line      int // 名称
	From      string
	To        string
	StartTime time.Time
}
