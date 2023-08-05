package suppliesBorrowController

import (
	"time"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/suppliesBorrowService"
	"wejh-go/app/utils"

	"github.com/gin-gonic/gin"
)

type CreateQAData struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// 创建 QA
func CreateQA(c *gin.Context) {
	var data CreateQAData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		_ = c.AbortWithError(200, apiException.ParamError)
		return
	}

	publisher := suppliesBorrowService.GetIdentity(c)
	err = suppliesBorrowService.CreateQA(models.QA{
		Title:       data.Title,
		Content:     data.Content,
		PublishTime: time.Now(),
		Publisher:   *publisher,
	})

	if err != nil {
		_ = c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

type GetQAListData struct {
	Publisher string `form:"publisher" binding:"required"`
}

// 用户获取对应发布方的 QA
func GetQAList(c *gin.Context) {
	var data GetQAListData
	err := c.ShouldBindQuery(&data)
	if err != nil {
		c.AbortWithError(200, apiException.ParamError)
		return
	}

	_, err = sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}

	var QAList []models.QA
	QAList, err = suppliesBorrowService.GetQAListByPublisher(data.Publisher)
	if err != nil {
		c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, QAList)
}

// 管理员获取自己发布的 QA
func GetQAListByAdmin(c *gin.Context) {
	admin := suppliesBorrowService.GetIdentity(c)

	var QAList []models.QA
	var err error
	if *admin == "Admin" {
		QAList, err = suppliesBorrowService.GetQAListBySuperAdmin()
	} else {
		QAList, err = suppliesBorrowService.GetQAListByPublisher(*admin)
	}

	if err != nil {
		c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, QAList)
}

type UpdateQAData struct {
	ID      int    `json:"id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// 更新 QA
func UpdateQA(c *gin.Context) {
	var data UpdateQAData
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.AbortWithError(200, apiException.ParamError)
		return
	}

	pubisher := suppliesBorrowService.GetIdentity(c)

	var QA *models.QA
	QA, err = suppliesBorrowService.GetQAbyID(data.ID)
	if err != nil {
		c.AbortWithError(200, apiException.ServerError)
		return
	}

	if QA.Publisher != *pubisher && *pubisher != "Admin" {
		c.AbortWithError(200, apiException.ServerError)
		return
	}

	err = suppliesBorrowService.UpdateQA(models.QA{
		ID:          data.ID,
		Title:       data.Title,
		Content:     data.Content,
		PublishTime: time.Now(),
	})
	if err != nil {
		c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}

type DeleteQAData struct {
	ID int `form:"id" binding:"required"`
}

// 删除 QA
func DeleteQA(c *gin.Context) {
	var data DeleteQAData
	err := c.ShouldBindQuery(&data)
	if err != nil {
		c.AbortWithError(200, apiException.ParamError)
		return
	}

	pubisher := suppliesBorrowService.GetIdentity(c)

	var QA *models.QA
	QA, err = suppliesBorrowService.GetQAbyID(data.ID)
	if err != nil {
		c.AbortWithError(200, apiException.ServerError)
		return
	}

	if QA.Publisher != *pubisher && *pubisher != "Admin" {
		c.AbortWithError(200, apiException.ServerError)
		return
	}

	err = suppliesBorrowService.DeleteQA(data.ID)
	if err != nil {
		c.AbortWithError(200, apiException.ServerError)
		return
	}

	utils.JsonSuccessResponse(c, nil)
}
