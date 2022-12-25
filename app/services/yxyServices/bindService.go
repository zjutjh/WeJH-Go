package yxyServices

import (
	"github.com/mitchellh/mapstructure"
	"net/url"
	"wejh-go/config/api/yxyApi"
)

type securityTokenForm struct {
	Level int    `json:"level"`
	Token string `json:"token"`
}

func GetSecurityToken(deviceId string) (*securityTokenForm, error) {
	params := url.Values{}
	Url, err := url.Parse(string(yxyApi.SecurityToken))
	if err != nil {
		return nil, err
	}
	params.Set("device_id", deviceId)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := FetchHandleOfGet(yxyApi.YxyApi(urlPath))
	if err != nil {
		return nil, err
	}
	var data securityTokenForm
	err = mapstructure.Decode(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetCaptchaImage(deviceId, token string) error {
	params := url.Values{}
	Url, err := url.Parse(string(yxyApi.CaptchaImage))
	if err != nil {
		return err
	}
	params.Set("device_id", deviceId)
	params.Set("security_token", token)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	_, err = FetchHandleOfGet(yxyApi.YxyApi(urlPath))
	if err != nil {
		return err
	}
	return nil
}
