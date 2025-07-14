package noticeController

import (
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"strings"
	"time"
	"wejh-go/app/apiException"
	"wejh-go/app/config"
	"wejh-go/app/models"
	"wejh-go/app/services/noticeServices"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"
)

type Publisher struct {
	Name               string `json:"name"`
	BackgroundImageUrl string `json:"backgroundImageUrl"`
}

type NoticeForm struct {
	ID      int         `json:"id"`
	Title   string      `json:"title"`
	Img1    interface{} `json:"img1"`
	Img2    interface{} `json:"img2"`
	Img3    interface{} `json:"img3"`
	Content string      `json:"content"`
}

func InsertNotice(c *gin.Context) {
	var postForm NoticeForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	publisher := getPublisher(c)
	notice := models.Notice{
		Title:       postForm.Title,
		PublishTime: time.Now(),
		Publisher:   *publisher,
		Content:     postForm.Content,
	}
	if postForm.Img1 != nil {
		notice.Img1 = postForm.Img1.(string)
	}
	if postForm.Img2 != nil {
		notice.Img2 = postForm.Img2.(string)
	}
	if postForm.Img3 != nil {
		notice.Img3 = postForm.Img3.(string)
	}
	err = noticeServices.CreateRecord(notice)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

func UpdateNotice(c *gin.Context) {
	var postForm NoticeForm
	err := c.ShouldBindJSON(&postForm)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	publisher := getPublisher(c)
	record, err := noticeServices.GetNoticeById(postForm.ID)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	if *publisher != record.Publisher {
		apiException.AbortWithException(c, apiException.NotAdmin, nil)
		return
	}
	err = noticeServices.UpdateNotice(postForm.ID, models.Notice{
		Title:     postForm.Title,
		Publisher: *publisher,
		Content:   postForm.Content,
	})
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, nil)
}

func DeleteNotice(c *gin.Context) {
	noticeId, err := strconv.Atoi(c.Query("notice_id"))
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	publisher := getPublisher(c)
	record, err := noticeServices.GetNoticeById(noticeId)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	if *publisher != record.Publisher {
		apiException.AbortWithException(c, apiException.NotAdmin, nil)
		return
	}
	err = noticeServices.DeleteNotice(noticeId)
	_ = os.Remove("./img/" + strings.TrimPrefix(record.Img1, config.GetWebpUrlKey()))
	_ = os.Remove("./img/" + strings.TrimPrefix(record.Img2, config.GetWebpUrlKey()))
	_ = os.Remove("./img/" + strings.TrimPrefix(record.Img3, config.GetWebpUrlKey()))
	utils.JsonSuccessResponse(c, nil)
}

func GetNoticeByAdmin(c *gin.Context) {
	var notice []models.Notice
	publisher := getPublisher(c)
	if *publisher == "Admin" {
		notice_, err := noticeServices.GetNoticeBySuperAdmin()
		notice = notice_
		if err != nil {
			apiException.AbortWithException(c, apiException.ServerError, err)
			return
		}
	} else {
		notice_, err := noticeServices.GetRecordByAdmin(*publisher)
		notice = notice_
		if err != nil {
			apiException.AbortWithException(c, apiException.ServerError, err)
			return
		}
	}
	utils.JsonSuccessResponse(c, notice)
}

func GetNotice(c *gin.Context) {
	notice, err := noticeServices.GetNoticeBySuperAdmin()
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, notice)
}

func getPublisher(c *gin.Context) *string {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return nil
	}
	var publisher string
	if user.Username == "zhforu" || user.Username == "pfforu" || user.Username == "mgsforu" {
		publisher = "\"For You\" 工程"
	} else if user.Type == models.Admin {
		publisher = "Admin"
	}
	return &publisher
}
