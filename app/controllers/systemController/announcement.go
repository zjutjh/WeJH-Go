package systemController

import (
	"wejh-go/app/apiException"
	"wejh-go/app/services/announcementServices"
	"wejh-go/app/utils"

	"github.com/gin-gonic/gin"
)

func GetAnnouncement(c *gin.Context) {
	announcements, err := announcementServices.GetAnnouncements(20)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
	} else {
		utils.JsonSuccessResponse(c, announcements)
	}

}
