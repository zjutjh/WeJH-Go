package yxyServices

import (
	"encoding/json"
	"net/url"
	"wejh-go/app/apiException"
	"wejh-go/app/utils/fetch"
	"wejh-go/config/api/yxyApi"
)

type YxyResponse struct {
	Code int         `json:"code" binding:"required"`
	Msg  string      `json:"message" binding:"required"`
	Data interface{} `json:"data"`
}

func FetchHandleOfPost(form url.Values, url yxyApi.YxyApi) (interface{}, error) {
	f := fetch.Fetch{}
	f.Init()
	res, err := f.PostForm(yxyApi.YxyHost+string(url), form)
	if err != nil {
		return nil, apiException.RequestError
	}
	rc := YxyResponse{}
	err = json.Unmarshal(res, &rc)
	if err != nil {
		return nil, apiException.RequestError
	}
	i := 0
	for rc.Code == 413 && i < 5 {
		i++
		res, err = f.PostForm(yxyApi.YxyHost+string(url), form)
		if err != nil {
			return nil, apiException.RequestError
		}
		rc = YxyResponse{}
		err = json.Unmarshal(res, &rc)
		if err != nil {
			return nil, apiException.RequestError
		}
	}

	if rc.Code == 413 {
		return rc.Data, apiException.ServerError
	}
	if rc.Code == 412 {
		return rc.Data, apiException.NoThatPasswordOrWrong
	}
	return rc.Data, nil
}
func FetchHandleOfGet(url yxyApi.YxyApi) (interface{}, error) {
	f := fetch.Fetch{}
	f.Init()
	res, err := f.Get(yxyApi.YxyHost + string(url))
	if err != nil {
		return nil, apiException.RequestError
	}
	rc := YxyResponse{}
	err = json.Unmarshal(res, &rc)
	if err != nil {
		return nil, apiException.RequestError
	}
	if rc.Code == 400 {
		return nil, apiException.ParamError
	}
	if rc.Code == 401 {
		return nil, apiException.NotLogin
	}
	if rc.Code == 403 {
		return nil, apiException.NotBindYxy
	}
	if rc.Code == 500 {
		return nil, apiException.ServerError
	}
	return rc.Data, nil
}
