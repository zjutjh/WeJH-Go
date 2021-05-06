package userServices

func CheckUsername(username string) bool {
	user := GetUserByUsername(username)
	if user != nil {
		return false
	}
	return true
}

func CheckWechatOpenID(wechatOpenID string) bool {
	user := GetUserByWechatOpenID(wechatOpenID)
	if user != nil {
		return false
	}
	return true
}
