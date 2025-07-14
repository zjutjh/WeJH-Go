package libraryController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/services/funnelServices"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/services/userServices"
	"wejh-go/app/utils"
)

func GetCurrent(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}
	list, err := funnelServices.GetCurrentBorrow(user)
	if err != nil {
		userServices.DelPassword(err, user, "Library")
		apiException.AbortWithError(c, err)
		return
	}
	utils.JsonSuccessResponse(c, list)
}

func GetHistory(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		apiException.AbortWithException(c, apiException.NotLogin, err)
		return
	}
	list, err := funnelServices.GetHistoryBorrow(user)
	if err != nil {
		userServices.DelPassword(err, user, "Library")
		apiException.AbortWithError(c, err)
		return
	}
	utils.JsonSuccessResponse(c, list)
}
