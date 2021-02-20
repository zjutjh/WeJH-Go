package database

type User struct {
	ID       int    // 登陆编号
	Uno      string // 学号
	Pass     string // 密码
	LibPass  string // 图书馆密码
	CardPass string // 饭卡密码
	ZfPass   string // 正方教务系统密码
}
