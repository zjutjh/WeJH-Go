package canteenController

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/services/funnelServices"
	"wejh-go/app/utils"
)

func GetCanteenFlowRate(c *gin.Context) {
	flowRate, err := funnelServices.GetCanteenFlowRate()
	if err != nil {
		utils.JsonErrorResponse(c, err)
	} else {
		utils.JsonSuccessResponse(c, flowRate)
	}
}
