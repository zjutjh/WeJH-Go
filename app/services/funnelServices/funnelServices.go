package funnelServices

import (
	"encoding/json"
	"net/url"
	"strings"
	"wejh-go/app/apiException"
	"wejh-go/app/utils/circuitBreaker"
	"wejh-go/app/utils/fetch"
	"wejh-go/config/api/funnelApi"
)

type FunnelResponse struct {
	Code int         `json:"code" binding:"required"`
	Msg  string      `json:"message" binding:"required"`
	Data interface{} `json:"data"`
}

func FetchHandleOfPost(form url.Values, host string, url funnelApi.FunnelApi) (interface{}, error) {
	f := fetch.Fetch{}
	f.Init()

	var rc FunnelResponse
	var res []byte
	var err error
	for i := 0; i < 5; i++ {
		res, err = f.PostForm(host+string(url), form)
		if err != nil {
			err = apiException.RequestError
			break
		}
		if err = json.Unmarshal(res, &rc); err != nil {
			err = apiException.RequestError
			break
		}
		if rc.Code != 413 {
			break
		}
	}

	loginType := funnelApi.LoginType(form.Get("type"))
	zfFlag := strings.Contains(string(url), "zf")
	if err != nil {
		if zfFlag {
			circuitBreaker.CB.Fail(host, loginType)
		}
		return nil, apiException.ServerError
	}

	if zfFlag {
		if rc.Code == 200 || rc.Code == 412 || rc.Code == 416 {
			circuitBreaker.CB.Success(host, loginType)
		} else {
			circuitBreaker.CB.Fail(host, loginType)
		}
	}

	switch rc.Code {
	case 200:
		return rc.Data, nil
	case 413:
		return nil, apiException.ServerError
	case 412:
		return nil, apiException.NoThatPasswordOrWrong
	case 416:
		return nil, apiException.OAuthNotUpdate
	default:
		return nil, apiException.ServerError
	}
}
