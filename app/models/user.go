package models

type User struct {
	ID           int    // 登录编号
	StudentID    string // 学号
	Password     string // 密码
	OpenID       string // 小程序 openID
	LibPassword  string // 图书馆密码
	CardPassword string // 饭卡密码
	ZFPassword   string // 正方教务系统密码
}
