package models

type AppList struct {
	ID              int64
	Title           string // 名称
	Route           string // 路由
	BackgroundColor string // 颜色
	Icon            string // 图标
}
