package announcementServices

import (
	"wejh-go/app/models"
	"wejh-go/config/database"
)

func GetAnnouncements(count int) []models.Announcement {
	var announcements []models.Announcement
	database.DB.Limit(count).Find(&announcements)
	return announcements
}

func GetAnnouncementPagination(offset, pagesize int) []models.Announcement {
	var announcements []models.Announcement
	database.DB.Offset(offset).Limit(pagesize).Find(&announcements)
	return announcements
}
