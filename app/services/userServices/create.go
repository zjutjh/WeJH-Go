package userServices

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"
	"wejh-go/app/models"
	"wejh-go/app/services/userCenterServices"
	"wejh-go/config/database"
)

func CreateStudentUser(username, password, studentID, IDCardNumber string) (*models.User, error) {
	if !CheckUsername(username) {
		return nil, errors.New("USERNAME")
	}

	err := userCenterServices.AuthStudent(studentID, IDCardNumber)
	if err != nil {
		return nil, err
	}

	h := sha256.New()
	h.Write([]byte(password))
	pass := hex.EncodeToString(h.Sum(nil))
	user := &models.User{
		JHPassword:  pass,
		Username:    username,
		Type:        models.Undergraduate,
		StudentID:   studentID,
		LibPassword: studentID,
		CreateTime:  time.Now(),
	}

	res := database.DB.Create(&user)

	return user, res.Error
}

func CreateStudentUserWechat(username, password, studentID, IDCardNumber, wechatOpenID string) (*models.User, error) {
	if !CheckWechatOpenID(wechatOpenID) {
		return nil, errors.New("WECHAT OpenID")
	}
	user, err := CreateStudentUser(username, password, studentID, IDCardNumber)
	if err != nil {
		return nil, err
	}
	user.WechatOpenID = wechatOpenID
	database.DB.Updates(user)
	database.DB.Save(user)
	return user, nil
}
