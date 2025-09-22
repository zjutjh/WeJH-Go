package adminController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/adminServices"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"
)

type GetUserBindStatusData struct {
	StudentID string `form:"student_id"`
}

func CheckAdmin(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}
	if user.Type != models.Admin {
		apiException.AbortWithException(c, apiException.NotAdmin, nil)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

func GetUserBindStatus(c *gin.Context) {
	var data GetUserBindStatusData
	err := c.ShouldBindQuery(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	zfStatus, ouathStatus, err := adminServices.GetBindStatus(data.StudentID)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	response := gin.H{
		"zf_status":    zfStatus,
		"oauth_status": ouathStatus,
	}
	utils.JsonSuccessResponse(c, response)
}
