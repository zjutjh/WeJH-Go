package announcementServices

import (
	"time"
	"wejh-go/app/models"

	"github.com/zjutjh/mygo/ndb"
)

func GetAnnouncements(count int) ([]models.Announcement, error) {
	var announcements []models.Announcement
	result := ndb.Pick().Limit(count).Find(&announcements)
	if result.Error != nil {
		return nil, result.Error
	}
	return announcements, nil
}

func CreateAnnouncement(announcement models.Announcement) error {
	announcement.PublishTime = time.Now()
	result := ndb.Pick().Create(&announcement)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func UpdateAnnouncement(id int, announcement models.Announcement) error {
	result := ndb.Pick().Model(models.Announcement{}).Where(
		&models.Announcement{
			ID: id,
		}).Updates(&announcement)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func DeleteAnnouncement(id int) error {
	result := ndb.Pick().Delete(models.Announcement{
		ID: id,
	})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetAnnouncementPagination(offset, pagesize int) []models.Announcement {
	var announcements []models.Announcement
	ndb.Pick().Offset(offset).Limit(pagesize).Find(&announcements)
	return announcements
}
