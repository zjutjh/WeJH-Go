package models

type LowBatteryAlertSubscription struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Campus    string `json:"campus"`
	Threshold int    `json:"threshhold"`
	Count     int    `json:"count"`
}
