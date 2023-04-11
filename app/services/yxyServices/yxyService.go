package yxyServices

import (
	"encoding/json"
	"fmt"
	"wejh-go/app/apiException"
	"wejh-go/app/utils/fetch"
	"wejh-go/config/api/yxyApi"
)

type YxyResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func FetchHandleOfPost(form map[string]string, url yxyApi.YxyApi) (*YxyResponse, error) {
	f := fetch.Fetch{}
	f.Init()
	res, err := f.PostJsonForm(yxyApi.YxyHost+string(url), form)
	if err != nil {
		fmt.Println(err)
		return nil, apiException.RequestError
	}
	rc := YxyResponse{}
	err = json.Unmarshal(res, &rc)
	if err != nil {
		fmt.Println(err)
		return nil, apiException.RequestError
	}
	return &rc, nil
}

func FetchHandleOfGet(url yxyApi.YxyApi) (*YxyResponse, error) {
	f := fetch.Fetch{}
	f.Init()
	res, err := f.Get(yxyApi.YxyHost + string(url))
	if err != nil {
		fmt.Println(err)
		return nil, apiException.RequestError
	}
	rc := YxyResponse{}
	err = json.Unmarshal(res, &rc)
	if err != nil {
		fmt.Println(err)
		return nil, apiException.RequestError
	}
	return &rc, nil
}
