package funnelServices

import (
	"wejh-go/config/api/funnelApi"
)

func GetCanteenFlowRate() (interface{}, error) {
	return FetchHandleOfGet(funnelApi.CanteenFlow)
}
