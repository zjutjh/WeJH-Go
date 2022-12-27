package yxyServices

import (
	"github.com/mitchellh/mapstructure"
	"net/url"
	"wejh-go/app/apiException"
	"wejh-go/config/api/yxyApi"
)

type securityToken struct {
	Level int    `json:"level"`
	Token string `json:"token"`
}

type captcha struct {
	Img string `json:"img"`
}

type exists struct {
	UserExists bool `json:"user_exists"`
}

func GetSecurityToken(deviceId string) (*securityToken, error) {
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
	var data securityToken
	err = mapstructure.Decode(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetCaptchaImage(deviceId, token string) (*string, error) {
	params := url.Values{}
	Url, err := url.Parse(string(yxyApi.CaptchaImage))
	if err != nil {
		return nil, err
	}
	params.Set("device_id", deviceId)
	params.Set("security_token", token)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := FetchHandleOfGet(yxyApi.YxyApi(urlPath))
	if err != nil {
		return nil, err
	}
	var data captcha
	err = mapstructure.Decode(resp, &data)
	if err != nil {
		return nil, err
	}
	return &data.Img, nil
}

func SendVerificationCode(securityToken, deviceId, captcha, phoneNum string) error {
	form := url.Values{}
	form.Set("phone_num", phoneNum)
	form.Set("security_token", securityToken)
	form.Set("captcha", captcha)
	form.Set("device_id", deviceId)
	resp, err := FetchHandleOfPost(form, yxyApi.SendVerificationCode)
	if err != nil {
		return err
	}
	var data exists
	err = mapstructure.Decode(resp, &data)
	if err != nil {
		return err
	}
	if !data.UserExists {
		return apiException.NotBindYxy
	}
	return nil
}

func LoginByCode(code, deviceId, phoneNum string) {

}
