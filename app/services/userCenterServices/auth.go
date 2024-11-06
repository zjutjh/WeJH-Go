package userCenterServices

import (
	"net/url"
	"wejh-go/app/apiException"
	"wejh-go/config/api/userCenterApi"
)

func Login(stu_id string, pass string) error {
	params := url.Values{}
	Url, err := url.Parse(string(userCenterApi.Auth))
	if err != nil {
		return err
	}
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	regMap := make(map[string]string)
	regMap["stu_id"] = stu_id
	regMap["password"] = pass
	resp, err := FetchHandleOfPost(regMap, userCenterApi.UserCenterApi(urlPath))
	if err != nil {
		return apiException.RequestError
	}
	if resp.Code == 404 {
		return apiException.UserNotFind
	} else if resp.Code == 405 {
		return apiException.NoThatPasswordOrWrong
	} else if resp.Code == 200 {
		return nil
	} else {
		return apiException.ServerError
	}
}
