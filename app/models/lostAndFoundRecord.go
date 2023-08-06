package models

import "time"

type LostAndFoundRecord struct {
	ID               int       `json:"id"`
	Type             bool      `json:"type"` // 1-失物 0-寻物
	Campus           string    `json:"campus"`
	Kind             string    `json:"kind"`
	PublishTime      time.Time `json:"publish_time" gorm:"type:timestamp;"`
	IsProcessed      bool      `json:"-"`
	Img1             string    `json:"img1"`
	Img2             string    `json:"img2"`
	Img3             string    `json:"img3"`
	Publisher        string    `json:"publisher"`
	ItemName         string    `json:"item_name"`           // 物品名称
	LostOrFoundPlace string    `json:"lost_or_found_place"` // 丢失或拾得地点
	LostOrFoundTime  string    `json:"lost_or_found_time"`  // 丢失或拾得时间
	PickupPlace      string    `json:"pickup_place"`        // 失物领取地点
	Contact          string    `json:"contact"`             // 寻物联系方式
	Introduction     string    `json:"introduction"`        // 物品介绍
}
