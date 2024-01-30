package userCenterServices

import (
	"net/url"
	"wejh-go/app/apiException"
	"wejh-go/config/api/userCenterApi"
)

// ResetPass reset the password in user-center
func ResetPass(stuID, iid, password string) error {
	params := url.Values{}
	Url, err := url.Parse(string(userCenterApi.RePassWithoutEmail))
	if err != nil {
		return err
	}
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	regMap := make(map[string]string)
	regMap["stuid"] = stuID
	regMap["iid"] = iid
	regMap["pwd"] = password
	resp, err := FetchHandleOfPost(regMap, userCenterApi.UserCenterApi(urlPath))
	if err != nil {
		return err
	}
	if resp.Code == 400 {
		return apiException.StudentNumAndIidError
	} else if resp.Code == 401 {
		return apiException.PwdError
	} else if resp.Code != 200 {
		return apiException.ServerError
	}
	return nil
}
