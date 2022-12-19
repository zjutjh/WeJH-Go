package userServices

import (
	"crypto/sha256"
	"encoding/hex"
	"wejh-go/app/models"
	"wejh-go/config/database"
)

func GetUserByWechatOpenID(openid string) *models.User {
	user := models.User{}
	result := database.DB.Where(
		&models.User{
			WechatOpenID: openid,
		},
	).First(&user)
	if result.Error != nil {
		return nil
	}

	DecryptUserKeyInfo(&user)
	return &user
}

func GetUserID(id int) (*models.User, error) {
	user := models.User{}
	result := database.DB.Where(
		&models.User{
			ID: id,
		},
	).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	DecryptUserKeyInfo(&user)
	return &user, nil
}

func GetUserByUsernameAndPassword(username, password string) (*models.User, error) {
	h := sha256.New()
	h.Write([]byte(password))
	user := models.User{}
	pass := hex.EncodeToString(h.Sum(nil))
	result := database.DB.Where(
		&models.User{
			Username:   username,
			JHPassword: pass,
		},
	).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	DecryptUserKeyInfo(&user)
	return &user, nil
}

func GetUserByUsername(username string) (*models.User, error) {
	user := models.User{}
	result := database.DB.Where(
		&models.User{
			Username: username,
		},
	).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	DecryptUserKeyInfo(&user)
	return &user, nil
}

func GetUserByStudentID(studentID string) (*models.User, error) {
	user := models.User{}
	result := database.DB.Where(
		&models.User{
			StudentID: studentID,
		},
	).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	DecryptUserKeyInfo(&user)
	return &user, nil
}
