package announcementServices

import (
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
	result := database.DB.Create(announcement)
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
