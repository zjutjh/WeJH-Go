package yxyServices

import (
	"github.com/mitchellh/mapstructure"
	"net/url"
	"wejh-go/app/apiException"
	"wejh-go/config/api/yxyApi"
)

type AuthResp struct {
	Token          string `json:"token"`
	ID             string `json:"id"`
	MobilePhone    string `json:"mobile_phone"`
	Sex            int    `json:"sex"`
	Platform       string `json:"platform"`
	ThirdOpenid    string `json:"third_openid"`
	SchoolCode     string `json:"school_code"`
	SchoolName     string `json:"school_name"`
	UserName       string `json:"user_name"`
	UserType       string `json:"user_type"`
	JobNo          string `json:"job_no"`
	UserIDCard     string `json:"user_id_card"`
	UserClass      string `json:"user_class"`
	BindCardStatus int    `json:"bind_card_status"`
}

type RoomInfo struct {
	ID           string `json:"id"`
	SchoolCode   string `json:"school_code"`
	SchoolName   string `json:"school_name"`
	AreaID       string `json:"area_id"`
	AreaName     string `json:"area_name"`
	BuildingCode string `json:"building_code"`
	BuildingName string `json:"building_name"`
	FloorCode    string `json:"floor_code"`
	FloorName    string `json:"floor_name"`
	RoomCode     string `json:"room_code"`
	RoomName     string `json:"room_name"`
	BindType     string `json:"bind_type"`
	CreateTime   string `json:"create_time"`
}

type EleBalance struct {
	SchoolCode      string  `json:"school_code"`
	AreaID          string  `json:"area_id"`
	BuildingCode    string  `json:"building_code"`
	FloorCode       string  `json:"floor_code"`
	RoomCode        string  `json:"room_code"`
	DisplayRoomName string  `json:"display_room_name"`
	Soc             float64 `json:"soc"`
	SocAmount       float64 `json:"soc_amount"`
	Surplus         float64 `json:"surplus"`
	SurplusAmount   float64 `json:"surplus_amount"`
	Subsidy         int     `json:"subsidy"`
	SubsidyAmount   int     `json:"subsidy_amount"`
	MdType          string  `json:"md_type"`
	MdName          string  `json:"md_name"`
	RoomStatus      string  `json:"room_status"`
}

type RechargeRecords []struct {
	RoomDm    string `json:"room_dm"`
	Datetime  string `json:"datetime"`
	BuyType   string `json:"buy_type"`
	UsingType string `json:"using_type"`
	Money     string `json:"money"`
	IsSend    string `json:"is_send"`
}

type EleConsumptionRecords []struct {
	RoomDm   string `json:"room_dm"`
	Datetime string `json:"datetime"`
	Used     string `json:"used"`
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

func Bind(token string) (*map[string]interface{}, error) {
	params := url.Values{}
	Url, err := url.Parse(string(yxyApi.Bind))
	if err != nil {
		return nil, err
	}
	params.Set("token", token)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := FetchHandleOfGet(yxyApi.YxyApi(urlPath))
	if err != nil {
		return nil, err
	}
	m := resp.Data.(map[string]interface{})
	return &m, nil
}

func ElectricityBalance(token string) (*map[string]interface{}, error) {
	params := url.Values{}
	Url, err := url.Parse(string(yxyApi.ElectricityBalance))
	if err != nil {
		return nil, err
	}
	params.Set("token", token)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := FetchHandleOfGet(yxyApi.YxyApi(urlPath))
	if err != nil {
		return nil, err
	}
	if resp.Data == nil {
		return nil, apiException.ServerError
	}
	m := resp.Data.(map[string]interface{})
	return &m, nil
}

func ElectricityRechargeRecords(token, areaId, buildingCode, floorCode, roomCode, page string) (*[]interface{}, error) {
	params := url.Values{}
	Url, err := url.Parse(string(yxyApi.RechargeRecords))
	if err != nil {
		return nil, err
	}
	params.Set("token", token)
	params.Set("area_id", areaId)
	params.Set("building_code", buildingCode)
	params.Set("floor_code", floorCode)
	params.Set("room_code", roomCode)
	params.Set("page", page)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := FetchHandleOfGet(yxyApi.YxyApi(urlPath))
	if err != nil {
		return nil, err
	}
	if resp.Data == nil {
		return nil, apiException.ServerError
	}
	m := resp.Data.([]interface{})
	return &m, nil
}

func ElectricityConsumptionRecords(token, areaId, buildingCode, floorCode, roomCode, mdType string) (*[]interface{}, error) {
	params := url.Values{}
	Url, err := url.Parse(string(yxyApi.ElectricityConsumption))
	if err != nil {
		return nil, err
	}
	params.Set("token", token)
	params.Set("area_id", areaId)
	params.Set("building_code", buildingCode)
	params.Set("floor_code", floorCode)
	params.Set("room_code", roomCode)
	params.Set("md_type", mdType)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := FetchHandleOfGet(yxyApi.YxyApi(urlPath))
	if err != nil {
		return nil, err
	}
	if resp.Data == nil {
		return nil, apiException.ServerError
	}
	m := resp.Data.([]interface{})
	return &m, nil
}
