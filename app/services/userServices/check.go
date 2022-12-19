package userServices

func CheckUsername(username string) bool {
	user, _ := GetUserByUsername(username)
	if user != nil {
		return true
	}
	return false
}

func CheckWechatOpenID(wechatOpenID string) bool {
	user := GetUserByWechatOpenID(wechatOpenID)
	if user != nil {
		return false
	}
	return true
}
