package yxyServices

import (
	"net/url"
	"wejh-go/app/apiException"
	"wejh-go/config/api/yxyApi"

	"github.com/mitchellh/mapstructure"
)

type securityToken struct {
	Level         int    `json:"level" mapstructure:"level"`
	SecurityToken string `json:"security_token" mapstructure:"security_token"`
}

type captcha struct {
	Img string `json:"img" mapstructure:"img"`
}

type userInfo struct {
	UID            string `json:"uid" mapstructure:"uid"`
	Token          string `json:"token" mapstructure:"token"`
	BindCardStatus int    `json:"bind_card_status" mapstructure:"bind_card_status"`
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
	err = mapstructure.Decode(resp.Data, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func GetCaptchaImage(deviceId, securityToken string) (*string, error) {
	params := url.Values{}
	Url, err := url.Parse(string(yxyApi.CaptchaImage))
	if err != nil {
		return nil, err
	}
	params.Set("device_id", deviceId)
	params.Set("security_token", securityToken)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := FetchHandleOfGet(yxyApi.YxyApi(urlPath))
	if err != nil {
		return nil, err
	}
	var data captcha
	err = mapstructure.Decode(resp.Data, &data)
	if err != nil {
		return nil, err
	}
	return &data.Img, nil
}

func SendVerificationCode(securityToken, deviceId, phoneNum string) error {
	form := make(map[string]any)
	form["phone_num"] = phoneNum
	form["security_token"] = securityToken
	form["captcha"] = ""
	form["device_id"] = deviceId
	resp, err := FetchHandleOfPost(form, yxyApi.SendVerificationCode)
	if err != nil {
		return err
	}

	if resp.Code == 110005 {
		return apiException.WrongPhoneNum
	} else if resp.Code == 110006 {
		return apiException.SendVerificationCodeLimit
	} else if resp.Code != 0 {
		return apiException.ServerError
	}

	m := resp.Data.(map[string]interface{})
	if m["user_exists"] == false {
		return apiException.NotBindYxy
	}
	return nil
}

func LoginByCode(code, deviceId, phoneNum string) (*string, error) {
	form := make(map[string]any)
	form["phone_num"] = phoneNum
	form["code"] = code
	form["device_id"] = deviceId
	resp, err := FetchHandleOfPost(form, yxyApi.LoginByCode)
	if err != nil {
		return nil, err
	}

	if resp.Code == 110007 || resp.Code == 110008 {
		return nil, apiException.WrongVerificationCode
	} else if resp.Code == 110005 {
		return nil, apiException.WrongPhoneNum
	} else if resp.Code != 0 {
		return nil, apiException.ServerError
	}

	var data userInfo
	err = mapstructure.Decode(resp.Data, &data)
	if err != nil {
		return nil, err
	}
	return &data.UID, nil
}

func SilentLogin(deviceId, uid, phoneNum string) error {
	form := make(map[string]any)
	form["uid"] = uid
	form["device_id"] = deviceId
	form["phone_num"] = phoneNum
	resp, err := FetchHandleOfPost(form, yxyApi.SlientLogin)
	if err != nil {
		return err
	}
	if resp.Code == 100101 || resp.Code == 100102 {
		return apiException.YxySessionExpired
	} else if resp.Code != 0 {
		return apiException.ServerError
	}
	return nil
}
