package funnelServices

import (
	"encoding/json"
	"net/url"
	"wejh-go/app/apiExpection"
	"wejh-go/app/utils/fetch"
	"wejh-go/config/api/funnelApi"
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
		return nil, apiExpection.RequestError
	}
	rc := FunnelResponse{}
	err = json.Unmarshal(res, &rc)
	if err != nil {
		return nil, apiExpection.RequestError
	}
	i := 0
	for rc.Code == 413 && i < 5 {
		i++
		res, err = f.PostForm(funnelApi.FunnelHost+string(url), form)
		if err != nil {
			return nil, apiExpection.RequestError
		}
		rc = FunnelResponse{}
		err = json.Unmarshal(res, &rc)
		if err != nil {
			return nil, apiExpection.RequestError
		}
	}

	if rc.Code == 413 {
		return rc.Data, apiExpection.ServerError
	}
	if rc.Code == 412 {
		return rc.Data, apiExpection.NoThatPasswordOrWrong
	}
	return rc.Data, nil
}
func FetchHandleOfGet(url funnelApi.FunnelApi) (interface{}, error) {
	f := fetch.Fetch{}
	f.Init()
	res, err := f.Get(funnelApi.FunnelHost + string(url))
	if err != nil {
		return nil, apiExpection.RequestError
	}
	rc := FunnelResponse{}
	err = json.Unmarshal(res, &rc)
	if err != nil {
		return nil, apiExpection.RequestError
	}
	return rc.Data, nil
}
