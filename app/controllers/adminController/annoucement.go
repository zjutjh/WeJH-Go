package adminController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/models"
	"wejh-go/app/services/announcementServices"
	"wejh-go/app/utils"
)

func CreateAnnouncement(c *gin.Context) {
	var postForm models.Announcement
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}

	err = announcementServices.CreateAnnouncement(postForm)
	if err != nil {
		utils.JsonErrorResponse(c, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}
func UpdateAnnouncement(c *gin.Context) {

}
func DeleteAnnouncement(c *gin.Context) {

}
