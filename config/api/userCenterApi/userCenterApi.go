package userCenterApi

import "github.com/zjutjh/mygo/config"

func GetUserCenterHost() string {
	return config.Pick().GetString("user.host")
}

type UserCenterApi string

const (
	Register           UserCenterApi = "api/register"
	Auth               UserCenterApi = "api/auth"
	RePassWithoutEmail UserCenterApi = "api/repass"
	DelAccount         UserCenterApi = "api/del"
)
