package suppliesController

import (
	"errors"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/suppliesServices"
	"wejh-go/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 获取填写的个人信息
func GetPersonalInfo(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}

	var personalInfo *models.PersonalInfo
	personalInfo, err = suppliesServices.GetPersonalInfoByStudentID(user.StudentID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, personalInfo)
}

type PersonalInfoData struct {
	Name      string `json:"name" binding:"required"`
	Gender    string `json:"gender" binding:"required"`
	College   string `json:"college" binding:"required"`
	Dormitory string `json:"dormitory" binding:"required"`
	Contact   string `json:"contact" binding:"required"`
}

// 创建或更新个人信息
func SavePersonalInfo(c *gin.Context) {
	var data PersonalInfoData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	var user *models.User
	user, err = sessionServices.GetUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}

	err = suppliesServices.SavePersonalInfo(models.PersonalInfo{
		Name:      data.Name,
		Gender:    data.Gender,
		StudentID: user.StudentID,
		College:   data.College,
		Dormitory: data.Dormitory,
		Contact:   data.Contact,
	})
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

// 获取学生个人信息
type GetPersonalInfoData struct {
	StudentID string `form:"student_id" binding:"required"`
}

func GetPersonalInfoByAdmin(c *gin.Context) {
	var data GetPersonalInfoData
	err := c.ShouldBindQuery(&data)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}

	admin := getIdentity(c)
	if *admin != "学生事务大厅" && *admin != "Admin" {
		apiException.AbortWithException(c, apiException.NotAdmin, nil)
		return
	}

	var personalInfo *models.PersonalInfo
	personalInfo, err = suppliesServices.GetPersonalInfoByStudentID(data.StudentID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}

	utils.JsonSuccessResponse(c, personalInfo)
}
