package busController

import (
	"net/http"
	"wejh-go/app/apiException"
	"wejh-go/app/config"
	"wejh-go/app/services/yxyServices"
	"wejh-go/app/utils"

	"github.com/gin-gonic/gin"
)

type getBusInfoRequest struct {
	Search string `form:"search"`
}

func GetBusInfo(c *gin.Context) {
	var req getBusInfoRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	busInfo, err := yxyServices.GetBusInfo(req.Search)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, busInfo)
}

type getBusAnnouncementRequest struct {
	Page     int `form:"page" default:"1"`
	PageSize int `form:"page_size" default:"10"`
}

func GetBusAnnouncement(c *gin.Context) {
	var req getBusAnnouncementRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		apiException.AbortWithException(c, apiException.ParamError, err)
		return
	}
	page := req.Page
	pageSize := req.PageSize
	if page < 1 {
		page = 1
	}
	if pageSize > 10 || pageSize < 1 {
		pageSize = 10
	}
	busInfo, err := yxyServices.GetAnnouncement(page, pageSize)
	if err != nil {
		apiException.AbortWithException(c, apiException.ServerError, err)
		return
	}
	utils.JsonSuccessResponse(c, busInfo)
}

func GetBusConfig(c *gin.Context) {
	url := config.GetBusConfigUrl()
	if url == "" {
		apiException.AbortWithException(c, apiException.NotFound, nil)
		return
	}
	c.Redirect(http.StatusFound, url)
}
