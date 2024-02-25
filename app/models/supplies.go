package models

import "gorm.io/gorm"

type Supplies struct {
	ID           int            `json:"id"`
	Name         string         `json:"name"`   // 物资名称
	Kind         string         `json:"kind"`   // 种类 正装
	Spec         string         `json:"spec"`   // 规格
	Campus       uint8          `json:"campus"` // 校区 1:朝晖 2:屏峰 3: 莫干山
	Organization string         `json:"organization"`
	Stock        uint           `json:"stock"`    // 库存
	Borrowed     uint           `json:"borrowed"` // 借出数量
	Img          string         `json:"img"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
