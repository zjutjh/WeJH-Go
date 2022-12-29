package yxyServices

import (
	"github.com/mitchellh/mapstructure"
	"net/url"
	"wejh-go/app/apiException"
	"wejh-go/config/api/yxyApi"
)

const ZJUTSchoolCode = "10337"

type balance struct {
	Balance string `json:"balance"`
}

func GetCardBalance(deviceId, uid string) (*string, error) {
	params := url.Values{}
	Url, err := url.Parse(string(yxyApi.CardBalance))
	if err != nil {
		return nil, err
	}
	params.Set("device_id", deviceId)
	params.Set("uid", uid)
	params.Set("school_code", ZJUTSchoolCode)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := FetchHandleOfGet(yxyApi.YxyApi(urlPath))
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, apiException.YxySessionExpired
	}
	var data balance
	err = mapstructure.Decode(resp.Data, &data)
	if err != nil {
		return nil, err
	}
	return &data.Balance, nil
}
