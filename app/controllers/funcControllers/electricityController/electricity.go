package electricityController

/*
#cgo CFLAGS: -I${SRCDIR}/lib
#cgo LDFLAGS: -L${SRCDIR}/lib -lyxy
#include <stdlib.h>
#include "../../../../include/yxy.h"
*/
import "C"

import (
	"github.com/gin-gonic/gin"
	"wejh-go/app/apiException"
	"wejh-go/app/services/sessionServices"
	"wejh-go/app/utils"
)

func GetElectricity(c *gin.Context) {
	user, err := sessionServices.GetUserSession(c)
	if err != nil {
		_ = c.AbortWithError(200, apiException.NotLogin)
		return
	}
	var eleResult *C.ele_result
	session := C.auth(C.CString(user.YXYUid))
	errCode := C.query_ele(session, &eleResult)
	if errCode != 0 {
		if errCode != 202 {
			_ = c.AbortWithError(200, apiException.ServerError)
		} else {
			_ = c.AbortWithError(200, apiException.NotBindYxy)
		}
	}
	utils.JsonSuccessResponse(c, gin.H{
		"total_surplus":     eleResult.total_surplus,
		"total_amount":      eleResult.total_amount,
		"surplus":           eleResult.surplus,
		"surplus_amount":    eleResult.surplus_amount,
		"subsidy":           eleResult.subsidy,
		"subsidy_amount":    eleResult.subsidy_amount,
		"display_room_name": C.GoString(eleResult.display_room_name),
		"room_status":       C.GoString(eleResult.room_status),
	})
	C.free_c_string(session)
	C.free_ele_result(eleResult)
}
