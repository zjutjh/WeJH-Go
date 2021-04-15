package userServices

import (
	"crypto/sha256"
	"encoding/hex"
	"wejh-go/app/models"
	"wejh-go/config/database"
)

func GetUserByOpenID(openid string) (*models.User, error) {
	user := models.User{}
	result := database.DB.Where(
		&models.User{
			OpenID: openid,
		},
	).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
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
	return &user, nil
}

func GetUserByUnionID(unionID string) (*models.User, error) {
	user := models.User{}
	result := database.DB.Where(
		&models.User{
			UnionID: unionID,
		},
	).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
