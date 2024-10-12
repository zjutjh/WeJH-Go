package themeServices

import (
	"wejh-go/app/models"
	"wejh-go/config/database"
)

func GetAllStudentIDs() ([]string, error) {
	var studentIDs []string
	var users []models.User
	err := database.DB.Select("student_id").Find(&users).Error
	if err != nil {
		return nil, err
	}
	for _, user := range users {
		studentIDs = append(studentIDs, user.StudentID)
	}
	return studentIDs, nil
}
