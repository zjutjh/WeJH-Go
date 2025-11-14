package lessonServices

import (
	"wejh-go/app/models"

	"github.com/zjutjh/mygo/ndb"
)

func CreateLesson(lesson models.Lesson) error {
	result := ndb.Pick().Create(&lesson)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetLesson(username, term, year string) ([]models.Lesson, error) {
	var lessons []models.Lesson
	result := ndb.Pick().Where(models.Lesson{
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
	result := ndb.Pick().Model(models.Lesson{}).Where(
		&models.Lesson{
			ID: id,
		}).Updates(&lesson)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteLesson(id int) error {
	result := ndb.Pick().Delete(models.Lesson{
		ID: id,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
