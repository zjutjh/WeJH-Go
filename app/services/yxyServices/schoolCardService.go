package yxyServices

import (
	"net/url"
	"wejh-go/app/apiException"
	"wejh-go/config/api/yxyApi"

	"github.com/mitchellh/mapstructure"
)

type BalanceResp struct {
	Balance string `json:"balance" mapstructure:"balance"`
}

type ConsumptionRecords struct {
	List []struct {
		Address string `json:"address" mapstructure:"address"`
		Money   string `json:"money" mapstructure:"money"`
		Time    string `json:"time" mapstructure:"time"`
	} `json:"list" mapstructure:"list"`
}

func GetCardBalance(deviceId, uid, phoneNum string) (*string, error) {
	params := url.Values{}
	Url, err := url.Parse(string(yxyApi.CardBalance))
	if err != nil {
		return nil, err
	}
	params.Set("device_id", deviceId)
	params.Set("uid", uid)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := FetchHandleOfGet(yxyApi.YxyApi(urlPath))
	if err != nil {
		return nil, err
	}

	if resp.Code == 100101 || resp.Code == 100102 {
		return nil, apiException.YxySessionExpired
	} else if resp.Code == 100103 {
		return nil, apiException.NotBindCard
	} else if resp.Code != 0 {
		return nil, apiException.ServerError
	}

	var data BalanceResp
	err = mapstructure.Decode(resp.Data, &data)
	if err != nil {
		return nil, err
	}

	_ = SilentLogin(deviceId, uid, phoneNum)

	return &data.Balance, nil
}

func GetCardConsumptionRecord(deviceId, uid, phoneNum, queryTime string) (*ConsumptionRecords, error) {
	params := url.Values{}
	Url, err := url.Parse(string(yxyApi.ConsumptionRecords))
	if err != nil {
		return nil, err
	}
	params.Set("device_id", deviceId)
	params.Set("uid", uid)
	params.Set("query_time", queryTime)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := FetchHandleOfGet(yxyApi.YxyApi(urlPath))
	if err != nil {
		return nil, err
	}

	if resp.Code == 100101 || resp.Code == 100102 {
		return nil, apiException.YxySessionExpired
	} else if resp.Code == 100103 {
		return nil, apiException.NotBindCard
	} else if resp.Code == 100002 {
		return nil, apiException.ParamError
	} else if resp.Code != 0 {
		return nil, apiException.ServerError
	}

	var data ConsumptionRecords
	err = mapstructure.Decode(resp.Data, &data)
	if err != nil {
		return nil, err
	}

	_ = SilentLogin(deviceId, uid, phoneNum)

	return &data, nil
}
