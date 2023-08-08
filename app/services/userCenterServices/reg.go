package userCenterServices

import (
	"net/url"
	"wejh-go/app/apiException"
	"wejh-go/config/api/userCenterApi"
)

func RegWithoutVerify(stu_id string, pass string, iid string, email string) error {
	params := url.Values{}
	Url, err := url.Parse(string(userCenterApi.UCRegWithoutVerify))
	if err != nil {
		return err
	}
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	regMap := make(map[string]string)
	regMap["stu_id"] = stu_id
	regMap["password"] = pass
	regMap["iid"] = iid
	regMap["email"] = email
	resp, err := FetchHandleOfPost(regMap, userCenterApi.UserCenterApi(urlPath))
	if err != nil {
		return err
	}
	if resp.Code == 400 || resp.Code == 402 {
		return apiException.StudentNumAndIidError
	} else if resp.Code == 401 {
		return apiException.PwdError
	} else if resp.Code == 403 {
		return apiException.ReactiveError
	}
	return nil
}
