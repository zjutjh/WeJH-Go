package userCenterApi

import "wejh-go/config/config"

var UserCenterHost = config.Config.GetString("user.host")
