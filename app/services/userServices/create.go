package userServices

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
	"wejh-go/app/apiExpection"
	"wejh-go/app/models"
	"wejh-go/app/services/userCenterServices"
	"wejh-go/config/database"
)

func CreateStudentUser(username, password, studentID, IDCardNumber, email string) (*models.User, error) {
	if !CheckUsername(username) {
		return nil, apiExpection.UserAlreadyExisted
	}

	err := userCenterServices.OldActiveStudent(studentID, password, IDCardNumber, email)
	if err != nil {
		return nil, err
	}

	h := sha256.New()
	h.Write([]byte(password))
	pass := hex.EncodeToString(h.Sum(nil))
	cardDefPass := ""
	if len(studentID) > 6 {
		cardDefPass = studentID[len(studentID)-6 : len(studentID)]
	}

	user := &models.User{
		JHPassword:   pass,
		Username:     username,
		Type:         models.Undergraduate,
		StudentID:    studentID,
		LibPassword:  studentID,
		CardPassword: cardDefPass,
		CreateTime:   time.Now(),
	}

	EncryptUserKeyInfo(user)
	res := database.DB.Create(&user)

	return user, res.Error
}

func CreateStudentUserWechat(username, password, studentID, IDCardNumber, email, wechatOpenID string) (*models.User, error) {
	if !CheckWechatOpenID(wechatOpenID) {
		return nil, apiExpection.OpenIDError
	}
	user, err := CreateStudentUser(username, password, studentID, IDCardNumber, email)
	if err != nil {
		return nil, err
	}
	user.WechatOpenID = wechatOpenID
	database.DB.Updates(user)
	database.DB.Save(user)
	return user, nil
}
