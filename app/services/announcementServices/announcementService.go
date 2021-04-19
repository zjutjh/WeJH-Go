package announcementServices

import (
	"time"
	"wejh-go/app/models"
	"wejh-go/config/database"
)

func GetAnnouncements(count int) ([]models.Announcement, error) {
	var announcements []models.Announcement
	result := database.DB.Limit(count).Find(&announcements)
	if result.Error != nil {
		return nil, result.Error
	}
	return announcements, nil
}

func CreateAnnouncement(announcement models.Announcement) error {
	announcement.PublishTime = time.Now()
	result := database.DB.Create(&announcement)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func UpdateAnnouncement(id int, announcement models.Announcement) error {
	result := database.DB.Model(models.Announcement{}).Where(
		&models.Announcement{
			ID: id,
		}).Updates(&announcement)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteAnnouncement(id int) error {
	result := database.DB.Delete(models.Announcement{
		ID: id,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetAnnouncementPagination(offset, pagesize int) []models.Announcement {
	var announcements []models.Announcement
	database.DB.Offset(offset).Limit(pagesize).Find(&announcements)
	return announcements
}
