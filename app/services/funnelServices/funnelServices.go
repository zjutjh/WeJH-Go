package funnelServices

import (
	"encoding/json"
	"net/url"
	"wejh-go/app/utils/fetch"
	"wejh-go/config/api/funnelApi"
)

type FunnelResponse struct {
	Code int         `json:"code" binding:"required"`
	Msg  string      `json:"message" binding:"required"`
	Data interface{} `json:"data"`
}

func FetchHandle(form url.Values, url string) (interface{}, error) {
	f := fetch.Fetch{}
	f.Init()
	res, err := f.PostForm(funnelApi.FunnelHost+url, form)
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
