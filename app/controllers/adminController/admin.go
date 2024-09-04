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
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	if user.Type != models.Admin {
		_ = c.AbortWithError(200, apiException.NotAdmin)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

func GetUserBindStatus(c *gin.Context) {
	var data GetUserBindStatusData
	err := c.ShouldBindQuery(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}
	zfStatus, ouathStatus, err := adminServices.GetBindStatus(data.StudentID)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}
	response := gin.H{
		"zf_status":    zfStatus,
		"oauth_status": ouathStatus,
	}
	utils.JsonSuccessResponse(c, response)
}
