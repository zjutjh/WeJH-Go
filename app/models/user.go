package models

import "time"

type User struct {
	ID           int    // 登录编号
	Username     string // 登录编号
	Type         UserType
	StudentID    string // 学号
	JHPassword   string // 密码
	WechatOpenID string // 小程序 openID
	UnionID      string
	LibPassword  string // 图书馆密码
	ZFPassword   string // 正方教务系统密码
	DeviceID     string // 生成的设备 uuid
	PhoneNum     string // 易校园绑定的手机号
	YXYUid       string // 易校园uid
	CreateTime   time.Time
}

type UserType int

const (
	Undergraduate UserType = 0
	Postgraduate  UserType = 1
	Teacher       UserType = 2
	Admin         UserType = 3
)
