package funnelServices

import (
	"net/url"
	"wejh-go/config/api/funnelApi"
)

func GetCanteenFlowRate() (interface{}, error) {
	form := url.Values{}
	return FetchHandle(form, funnelApi.ZFClassTable)
}
