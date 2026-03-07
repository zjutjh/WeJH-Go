package userCenterServices

import (
	"encoding/json"
	"wejh-go/app/apiException"
	"wejh-go/app/utils/fetch"
	"wejh-go/config/api/userCenterApi"
)

type UserCenterResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func FetchHandleOfPost(form map[string]any, url userCenterApi.UserCenterApi) (*UserCenterResponse, error) {
	f := fetch.Fetch{}
	f.Init()
	res, err := f.PostJsonForm(userCenterApi.GetUserCenterHost()+string(url), form)
	if err != nil {
		return nil, apiException.RequestError
	}
	rc := UserCenterResponse{}
	err = json.Unmarshal(res, &rc)
	if err != nil {
		return nil, apiException.RequestError
	}
	return &rc, nil
}

func FetchHandleOfGet(url userCenterApi.UserCenterApi) (*UserCenterResponse, error) {
	f := fetch.Fetch{}
	f.Init()
	res, err := f.Get(userCenterApi.GetUserCenterHost() + string(url))
	if err != nil {
		return nil, apiException.RequestError
	}
	rc := UserCenterResponse{}
	err = json.Unmarshal(res, &rc)
	if err != nil {
		return nil, apiException.RequestError
	}
	return &rc, nil
}
