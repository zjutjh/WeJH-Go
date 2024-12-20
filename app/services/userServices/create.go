package userServices

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/userCenterServices"
	"wejh-go/config/database"
)

func CreateStudentUser(username, password, studentID, IDCardNumber, email string, userType uint) (*models.User, error) {
	if CheckUsername(username) {
		return nil, apiException.UserAlreadyExisted
	}
	err := userCenterServices.RegWithoutVerify(studentID, password, IDCardNumber, email, userType)
	if err != nil && err != apiException.ReactiveError {
		return nil, err
	}

	h := sha256.New()
	h.Write([]byte(password))
	pass := hex.EncodeToString(h.Sum(nil))

	user := &models.User{
		JHPassword:  pass,
		Username:    username,
		Type:        models.UserType(userType),
		StudentID:   studentID,
		LibPassword: "",
		PhoneNum:    "",
		YxyUid:      "",
		DeviceID:    "",
		CreateTime:  time.Now(),
	}

	EncryptUserKeyInfo(user)
	res := database.DB.Create(&user)

	return user, res.Error
}

func CreateStudentUserWechat(username, password, studentID, IDCardNumber, email, wechatOpenID string, userType uint) (*models.User, error) {
	if CheckWechatOpenID(wechatOpenID) {
		return nil, apiException.OpenIDError
	}
	user, err := CreateStudentUser(username, password, studentID, IDCardNumber, email, userType)
	if err != nil && err != apiException.ReactiveError {
		return nil, err
	}
	user.WechatOpenID = wechatOpenID
	database.DB.Save(user)
	return user, nil
}

func CreateAdmin(userName, password string, adminType int) error {
	h := sha256.New()
	h.Write([]byte(password))
	pass := hex.EncodeToString(h.Sum(nil))
	admin := &models.User{
		JHPassword:  pass,
		Username:    userName,
		Type:        models.UserType(adminType),
		StudentID:   userName,
		LibPassword: "",
		PhoneNum:    "",
		YxyUid:      "",
		DeviceID:    "",
		CreateTime:  time.Now(),
	}
	res := database.DB.Create(&admin)
	return res.Error
}
