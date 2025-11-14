package userCenterApi

import "github.com/zjutjh/mygo/config"

func GetUserCenterHost()string{
	return config.Pick().GetString("user.host")
}

type UserCenterApi string

const (
	UCRegWithoutVerify UserCenterApi = "api/activation/notVerify"
	UCReg              UserCenterApi = "api/activation"
	VerifyEmail        UserCenterApi = "api/verify/email"
	ReSendEmail        UserCenterApi = "api/email"
	Auth               UserCenterApi = "api/auth"
	RePass             UserCenterApi = "api/changePwd"
	RePassWithoutEmail UserCenterApi = "api/repass"
	DelAccount         UserCenterApi = "api/del"
)
