package systemController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/services/announcementServices"
	"wejh-go/app/utils"
)

func GetAnnouncement(c *gin.Context) {
	announcements, err := announcementServices.GetAnnouncements(10)
	if err != nil {
		utils.JsonErrorResponse(c, err)
	} else {
		utils.JsonSuccessResponse(c, announcements)
	}

}
