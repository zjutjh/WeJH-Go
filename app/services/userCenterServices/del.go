package userCenterServices

import (
	"net/url"
	"wejh-go/app/apiException"
	"wejh-go/config/api/userCenterApi"
)

// DelAccount delete the account in user-center
func DelAccount(stuID, iid string) error {
	params := url.Values{}
	Url, err := url.Parse(string(userCenterApi.DelAccount))
	if err != nil {
		return err
	}
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	regMap := make(map[string]string)
	regMap["stuid"] = stuID
	regMap["iid"] = iid
	resp, err := FetchHandleOfPost(regMap, userCenterApi.UserCenterApi(urlPath))
	if err != nil {
		return err
	}
	if resp.Code == 400 {
		return apiException.StudentNumAndIidError
	} else if resp.Code != 200 {
		return apiException.ServerError
	}
	return nil
}
