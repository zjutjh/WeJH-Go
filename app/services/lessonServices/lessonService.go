package lessonServices

import (
	"wejh-go/app/models"
	"wejh-go/config/database"
)

func CreateLesson(lesson models.Lesson) error {
	result := database.DB.Create(&lesson)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetLesson(username, term, year string) ([]models.Lesson, error) {
	var lessons []models.Lesson
	result := database.DB.Where(models.Lesson{
		Username: username,
		Term:     term,
		Year:     year,
	}).Find(&lessons)
	if result.Error != nil {
		return nil, result.Error
	}
	return lessons, nil
}

func UpdateLesson(id int, lesson models.Lesson) error {
	result := database.DB.Model(models.Lesson{}).Where(
		&models.Lesson{
			ID: id,
		}).Updates(&lesson)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteLesson(id int) error {
	result := database.DB.Delete(models.Lesson{
		ID: id,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
