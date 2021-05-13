package funnelServices

import (
	"net/url"
	"strconv"
	"wejh-go/app/apiExpection"
	"wejh-go/app/models"
	"wejh-go/config/api/funnelApi"
)

func GetCardBalance(u *models.User) (interface{}, error) {
	if u.CardPassword == "" {
		return nil, apiExpection.NoThatPasswordOrWrong
	}
	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", u.CardPassword)
	return FetchHandleOfPost(form, funnelApi.CardBalance)
}

func GetCardToday(u *models.User) (interface{}, error) {
	if u.CardPassword == "" {
		return nil, apiExpection.NoThatPasswordOrWrong
	}
	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", u.CardPassword)

	return FetchHandleOfPost(form, funnelApi.CardToday)
}

func GetCardHistory(u *models.User, year, month int) (interface{}, error) {
	if u.CardPassword == "" {
		return nil, apiExpection.NoThatPasswordOrWrong
	}
	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", u.CardPassword)
	form.Add("year", strconv.Itoa(year))
	form.Add("month", strconv.Itoa(month))
	return FetchHandleOfPost(form, funnelApi.CardHistory)
}
