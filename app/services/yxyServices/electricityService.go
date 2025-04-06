package yxyServices

import (
	"net/url"
	"wejh-go/app/apiException"
	"wejh-go/config/api/yxyApi"

	"github.com/mitchellh/mapstructure"
)

type AuthResp struct {
	Token string `json:"token" mapstructure:"token"`
}

type ElecBalance struct {
	DisplayRoomName string  `json:"display_room_name" mapstructure:"display_room_name"`
	RoomStrConcat   string  `json:"room_str_concat" mapstructure:"room_str_concat"`
	Soc             float64 `json:"soc" mapstructure:"surplus"`
}

type RechargeRecords struct {
	List []struct {
		Money    string `json:"money" mapstructure:"money"`
		Datetime string `json:"datetime" mapstructure:"datetime"`
	} `json:"list" mapstructure:"list"`
}

type EleConsumptionRecords struct {
	List []struct {
		Used     string `json:"used" mapstructure:"usage"`
		Datetime string `json:"datetime" mapstructure:"datetime"`
	} `json:"list" mapstructure:"list"`
}

func Auth(uid string) (*string, error) {
	params := url.Values{}
	Url, err := url.Parse(string(yxyApi.Auth))
	if err != nil {
		return nil, err
	}
	params.Set("uid", uid)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := FetchHandleOfGet(yxyApi.YxyApi(urlPath))
	if err != nil {
		return nil, err
	}
	var data AuthResp
	err = mapstructure.Decode(resp.Data, &data)
	if err != nil {
		return nil, err
	}
	return &data.Token, nil
}

func ElectricityBalance(token, campus string) (*ElecBalance, error) {
	params := url.Values{}
	Url, err := url.Parse(string(yxyApi.ElectricityBalance))
	if err != nil {
		return nil, err
	}
	params.Set("token", token)
	params.Set("campus", campus)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := FetchHandleOfGet(yxyApi.YxyApi(urlPath))
	if err != nil {
		return nil, err
	}
	if resp.Code == 100103 {
		return nil, apiException.NotBindCard
	} else if resp.Code == 110102 {
		return nil, apiException.CampusMismatch
	} else if resp.Code != 0 {
		return nil, apiException.ServerError
	}
	var data ElecBalance
	err = mapstructure.Decode(resp.Data, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func ElectricityRechargeRecords(token, campus, page, roomStrConcat string) (*RechargeRecords, error) {
	params := url.Values{}
	Url, err := url.Parse(string(yxyApi.RechargeRecords))
	if err != nil {
		return nil, err
	}
	params.Set("token", token)
	params.Set("campus", campus)
	params.Set("page", page)
	params.Set("room_str_concat", roomStrConcat)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := FetchHandleOfGet(yxyApi.YxyApi(urlPath))
	if err != nil {
		return nil, err
	}
	if resp.Code == 110103 {
		return nil, apiException.CampusMismatch
	} else if resp.Code != 0 {
		return nil, apiException.ServerError
	}
	var data RechargeRecords
	err = mapstructure.Decode(resp.Data, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func ElectricityConsumptionRecords(token, campus, roomStrConcat string) (*EleConsumptionRecords, error) {
	params := url.Values{}
	Url, err := url.Parse(string(yxyApi.ElectricityConsumption))
	if err != nil {
		return nil, err
	}
	params.Set("token", token)
	params.Set("campus", campus)
	params.Set("room_str_concat", roomStrConcat)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := FetchHandleOfGet(yxyApi.YxyApi(urlPath))
	if err != nil {
		return nil, err
	}
	if resp.Code == 110103 {
		return nil, apiException.CampusMismatch
	} else if resp.Code != 0 {
		return nil, apiException.ServerError
	}
	var data EleConsumptionRecords
	err = mapstructure.Decode(resp.Data, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
