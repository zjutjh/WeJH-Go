package models

type User struct {
	ID           int    // 登录编号
	Username     string // 登录编号
	Type         UserType
	StudentID    string // 学号
	JHPassword   string // 密码
	OpenID       string // 小程序 openID
	UnionID      string
	LibPassword  string // 图书馆密码
	CardPassword string // 饭卡密码
	ZFPassword   string // 正方教务系统密码
}

type UserType int

const (
	Undergraduate UserType = 0
	Postgraduate  UserType = 1
	Teacher       UserType = 2
	Admin         UserType = 3
)
