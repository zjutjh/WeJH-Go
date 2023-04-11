package yxyServices

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"net/url"
	"strings"
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

type userInfo struct {
	UID                string `json:"uid"`
	Token              string `json:"token"`
	DeviceID           string `json:"device_id"`
	Sex                int    `json:"sex"`
	SchoolCode         string `json:"school_code"`
	SchoolName         string `json:"school_name"`
	SchoolClasses      int    `json:"school_classes"`
	SchoolNature       int    `json:"school_nature"`
	UserName           string `json:"user_name"`
	UserType           string `json:"user_type"`
	JobNo              string `json:"job_no"`
	UserIdcard         string `json:"user_idcard"`
	IdentityNo         string `json:"identity_no"`
	UserClass          string `json:"user_class"`
	RealNameStatus     int    `json:"real_name_status"`
	RegiserTime        string `json:"regiser_time"`
	BindCardStatus     int    `json:"bind_card_status"`
	LastLogin          string `json:"last_login"`
	TestAccount        int    `json:"test_account"`
	IsNew              int    `json:"is_new"`
	CreateStatus       int    `json:"create_status"`
	Platform           string `json:"platform"`
	BindCardRate       int    `json:"bind_card_rate"`
	Points             int    `json:"points"`
	SchoolIdentityType int    `json:"school_identity_type"`
	ExtraJSON          string `json:"extra_json"`
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
	err = mapstructure.Decode(resp.Data, &data)
	if err != nil {
		return nil, err
	}
	return &data.Img, nil
}

func SendVerificationCode(securityToken, deviceId, captcha, phoneNum string) error {
	var form map[string]string
	form = make(map[string]string)
	form["phone_num"] = phoneNum
	form["security_token"] = securityToken
	form["captcha"] = captcha
	form["device_id"] = deviceId
	resp, err := FetchHandleOfPost(form, yxyApi.SendVerificationCode)
	if err != nil {
		return err
	}
	if resp.Code != 0 {
		msgSplit := strings.Split(resp.Msg, "; ")
		if len(msgSplit) > 1 && msgSplit[1] == "encryptedDeviceId不一致" {
			fmt.Println(msgSplit[1])
			return apiException.ServerError
		} else {
			return apiException.WrongCaptcha
		}
	}
	m := resp.Data.(map[string]interface{})
	if m["user_exists"] == false {
		return apiException.NotBindYxy
	}
	return nil
}

func LoginByCode(code, deviceId, phoneNum string) (*string, error) {
	var form map[string]string
	form = make(map[string]string)
	form["phone_num"] = phoneNum
	form["code"] = code
	form["device_id"] = deviceId
	resp, err := FetchHandleOfPost(form, yxyApi.LoginByCode)
	if err != nil {
		return nil, err
	}
	if resp.Code == 403 {
		return nil, apiException.WrongVerificationCode
	} else if resp.Code == 500 {
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

func SilentLogin(deviceId, uid string) error {
	var form map[string]string
	form = make(map[string]string)
	form["uid"] = uid
	form["device_id"] = deviceId
	resp, err := FetchHandleOfPost(form, yxyApi.SlientLogin)
	if err != nil {
		return err
	}
	if resp.Code == 403 {
		fmt.Println(resp.Msg)
		return apiException.YxySessionExpired
	} else if resp.Code != 0 {
		return apiException.ServerError
	}
	var data userInfo
	err = mapstructure.Decode(resp.Data, &data)
	return err
}
