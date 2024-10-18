package userServices

import (
	"crypto/sha256"
	"encoding/hex"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/userCenterServices"
)

func CheckUsername(username string) bool {
	user, _ := GetUserByUsername(username)
	return user != nil
}

func CheckWechatOpenID(wechatOpenID string) bool {
	user := GetUserByWechatOpenID(wechatOpenID)
	return user != nil
}

func CheckLogin(username, password string) error {
	return userCenterServices.Login(username, password)
}

func CheckLocalLogin(user *models.User, password string) error {
	h := sha256.New()
	h.Write([]byte(password))
	pass := hex.EncodeToString(h.Sum(nil))

	if user.JHPassword != pass {
		return apiException.NoThatPasswordOrWrong
	}
	return nil
}
