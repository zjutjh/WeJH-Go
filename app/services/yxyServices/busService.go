package yxyServices

import (
	"net/url"
	"strconv"
	"wejh-go/app/apiException"
	"wejh-go/config/api/yxyApi"

	"github.com/mitchellh/mapstructure"
)

type BusInfoResp struct {
	UpdatedAt string `json:"updated_at" mapstructure:"updated_at"`
	List      []struct {
		Name     string   `json:"name" mapstructure:"name"`
		Price    int      `json:"price" mapstructure:"price"`
		Stations []string `json:"stations" mapstructure:"stations"`
		BusTime  []struct {
			DepartureTime string `json:"departure_time" mapstructure:"departure_time"`
			RemainSeats   int    `json:"remain_seats" mapstructure:"remain_seats"`
			OrderedSeats  int    `json:"ordered_seats" mapstructure:"ordered_seats"`
		} `json:"bus_time" mapstructure:"bus_time"`
	} `json:"list" mapstructure:"list"`
}

func GetBusInfo(search string) (*BusInfoResp, error) {
	params := url.Values{}
	Url, err := url.Parse(string(yxyApi.BusInfo))
	if err != nil {
		return nil, err
	}
	params.Set("search", search)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := FetchHandleOfGet(yxyApi.YxyApi(urlPath))
	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, apiException.ServerError
	}

	var data BusInfoResp
	err = mapstructure.Decode(resp.Data, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

type BusAnnouncementResp struct {
	UpdatedAt string `json:"updated_at" mapstructure:"updated_at"`
	Total     int    `json:"total" mapstructure:"total"`
	List      []struct {
		Title       string   `json:"title" mapstructure:"title"`
		Author      string   `json:"author" mapstructure:"author"`
		PublishedAt string   `json:"published_at" mapstructure:"published_at"`
		Abstract    string   `json:"abstract" mapstructure:"abstract"`
		Content     []string `json:"content" mapstructure:"content"`
	} `json:"list" mapstructure:"list"`
}

func GetAnnouncement(page, pageSize int) (*BusAnnouncementResp, error) {
	params := url.Values{}
	Url, err := url.Parse(string(yxyApi.BusAnnouncement))
	if err != nil {
		return nil, err
	}
	params.Set("page", strconv.Itoa(page))
	params.Set("page_size", strconv.Itoa(pageSize))
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := FetchHandleOfGet(yxyApi.YxyApi(urlPath))
	if err != nil {
		return nil, err
	}

	if resp.Code != 0 {
		return nil, apiException.ServerError
	}

	var data BusAnnouncementResp
	err = mapstructure.Decode(resp.Data, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
