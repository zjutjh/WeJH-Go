package userServices

import (
	"wejh-go/app/models"
	"wejh-go/service/database"
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
