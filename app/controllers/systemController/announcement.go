package systemController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/services/announcementServices"
	"wejh-go/app/utils"
)

func GetAnnouncement(c *gin.Context) {
	announcements, err := announcementServices.GetAnnouncements(20)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
	} else {
		utils.JsonSuccessResponse(c, announcements)
	}

}
