package models

type AppList struct {
	ID              int64  `json:"id"`
	Title           string `json:"title"`           // 名称
	Route           string `json:"route"`           // 路由
	BackgroundColor string `json:"backgroundColor"` // 颜色
	Icon            string `json:"icon"`            // 图标
	Require         string `json:"require"`         // 所需绑定
}
