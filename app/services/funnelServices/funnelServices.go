package funnelServices

import (
	"encoding/json"
	"net/url"
	"wejh-go/app/utils/fetch"
	"wejh-go/config/api/funnelApi"
	"wejh-go/errors"
)

type FunnelResponse struct {
	Code int         `json:"code" binding:"required"`
	Msg  string      `json:"message" binding:"required"`
	Data interface{} `json:"data"`
}

func FetchHandleOfPost(form url.Values, url funnelApi.FunnelApi) (interface{}, error) {
	f := fetch.Fetch{}
	f.Init()
	res, err := f.PostForm(funnelApi.FunnelHost+string(url), form)
	if err != nil {
		return nil, err
	}
	rc := FunnelResponse{}
	err = json.Unmarshal(res, &rc)
	if err != nil {
		return nil, err
	}
	i := 0
	for rc.Code == 413 && i < 5 {
		i++
		res, err = f.PostForm(funnelApi.FunnelHost+string(url), form)
		if err != nil {
			return nil, err
		}
		rc = FunnelResponse{}
		err = json.Unmarshal(res, &rc)
		if err != nil {
			return nil, err
		}
	}
	if rc.Code == 412 {
		return rc.Data, errors.PasswordWrong
	}
	if rc.Code == 413 {
		return rc.Data, errors.SystemError
	}
	return rc.Data, nil
}
func FetchHandleOfGet(url funnelApi.FunnelApi) (interface{}, error) {
	f := fetch.Fetch{}
	f.Init()
	res, err := f.Get(funnelApi.FunnelHost + string(url))
	if err != nil {
		return nil, err
	}
	rc := FunnelResponse{}
	err = json.Unmarshal(res, &rc)
	if err != nil {
		return nil, err
	}
	return rc.Data, nil
}
