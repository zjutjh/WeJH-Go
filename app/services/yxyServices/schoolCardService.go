package yxyServices

import (
	"github.com/mitchellh/mapstructure"
	"net/url"
	"wejh-go/app/apiException"
	"wejh-go/config/api/yxyApi"
)

const ZJUTSchoolCode = "10337"

type BalanceResp struct {
	Balance string `json:"balance"`
}

type ConsumptionRecords []struct {
	Type     string `json:"type"`
	FeeName  string `json:"fee_name"`
	Time     string `json:"time"`
	SerialNo string `json:"serial_no"`
	Money    string `json:"money"`
	DealTime string `json:"deal_time"`
	Address  string `json:"address"`
}

func GetCardBalance(deviceId, uid string) (*string, error) {
	params := url.Values{}
	Url, err := url.Parse(string(yxyApi.CardBalance))
	if err != nil {
		return nil, err
	}
	params.Set("device_id", deviceId)
	params.Set("uid", uid)
	params.Set("school_code", ZJUTSchoolCode)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	_ = SilentLogin(deviceId, uid)
	resp, err := FetchHandleOfGet(yxyApi.YxyApi(urlPath))
	if err != nil {
		return nil, err
	}
	if resp.Code != 0 {
		return nil, apiException.YxySessionExpired
	}
	var data BalanceResp
	err = mapstructure.Decode(resp.Data, &data)
	if err != nil {
		return nil, err
	}
	return &data.Balance, nil
}

func GetCardConsumptionRecord(deviceId, uid, queryTime string) (ConsumptionRecords, error) {
	params := url.Values{}
	Url, err := url.Parse(string(yxyApi.ConsumptionRecords))
	if err != nil {
		return nil, err
	}
	params.Set("device_id", deviceId)
	params.Set("uid", uid)
	params.Set("query_time", queryTime)
	params.Set("school_code", ZJUTSchoolCode)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	_ = SilentLogin(deviceId, uid)
	resp, err := FetchHandleOfGet(yxyApi.YxyApi(urlPath))
	if err != nil {
		return nil, err
	}
	if resp.Code == 204 {
		return nil, nil
	} else if resp.Code == 500 {
		return nil, apiException.YxySessionExpired
	} else if resp.Code == 403 {
		return nil, apiException.ParamError
	} else if resp.Code != 0 {
		return nil, apiException.ServerError
	}
	var data ConsumptionRecords
	err = mapstructure.Decode(resp.Data, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
