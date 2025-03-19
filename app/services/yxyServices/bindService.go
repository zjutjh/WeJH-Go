package yxyServices

import (
	"net/url"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/config/api/yxyApi"
	"wejh-go/config/database"
	r "wejh-go/config/redis"

	"github.com/mitchellh/mapstructure"
)

type securityToken struct {
	Level         int    `json:"level" mapstructure:"level"`
	SecurityToken string `json:"security_token" mapstructure:"security_token"`
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

func LoginByCode(code, deviceId, phoneNum string) (*userInfo, error) {
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
	if data.BindCardStatus == 0 {
		return nil, apiException.NotBindCard
	}
	return &data, nil
}

func SilentLogin(deviceId, uid, phoneNum, token string) error {
	form := make(map[string]any)
	form["uid"] = uid
	form["device_id"] = deviceId
	form["phone_num"] = phoneNum
	form["token"] = token
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

func Unbind(id int, uid string, isNotBindCard bool) error {
	updates := map[string]interface{}{
		"device_id": "",
	}
	if isNotBindCard {
		updates["yxy_uid"] = ""
	}
	if err := database.DB.Model(&models.User{}).
		Where("id = ?", id).
		Updates(updates).Error; err != nil {
		return err
	}
	cacheKey := "card:auth_token:" + uid
	_ = r.RedisClient.Del(ctx, cacheKey)
	return nil
}
