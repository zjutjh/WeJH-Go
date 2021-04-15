package systemController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/services/announcementServices"
	"wejh-go/app/utils"
)

func GetAnnouncement(c *gin.Context) {
	announcements := announcementServices.GetAnnouncements(10)

	utils.JsonSuccessResponse(c, announcements)
}
