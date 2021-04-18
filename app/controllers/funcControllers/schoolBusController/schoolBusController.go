package schoolBusController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/services/schoolBusServices"
	"wejh-go/app/utils"
)

func GetBusList(c *gin.Context) {
	buslist, err := schoolBusServices.GetSchoolBusList()
	if err != nil {
		utils.JsonErrorResponse(c, err)
	} else {
		utils.JsonSuccessResponse(c, buslist)
	}
}
