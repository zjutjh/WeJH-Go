package funnelServices

import (
	"net/url"
	"wejh-go/app/models"
	"wejh-go/config/api/funnelApi"
)

func genTermForm(u *models.User, year, term string, loginType funnelApi.LoginType) url.Values {
	var password string

	if loginType == "OAUTH" {
		password = u.OauthPassword
	} else {
		password = u.ZFPassword
	}

	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", password)
	form.Add("type", string(loginType))
	form.Add("year", year)
	form.Add("term", term)
	return form
}

func GetClassTable(u *models.User, year, term, host string, loginType funnelApi.LoginType) (interface{}, error) {
	form := genTermForm(u, year, term, loginType)
	return FetchHandleOfPost(form, host, funnelApi.ZFClassTable)
}

func GetScore(u *models.User, year, term, host string, loginType funnelApi.LoginType) (interface{}, error) {
	form := genTermForm(u, year, term, loginType)
	return FetchHandleOfPost(form, host, funnelApi.ZFScore)
}

func GetMidTermScore(u *models.User, year, term, host string, loginType funnelApi.LoginType) (interface{}, error) {
	form := genTermForm(u, year, term, loginType)
	return FetchHandleOfPost(form, host, funnelApi.ZFMidTermScore)
}

func GetExam(u *models.User, year, term, host string, loginType funnelApi.LoginType) (interface{}, error) {
	form := genTermForm(u, year, term, loginType)
	return FetchHandleOfPost(form, host, funnelApi.ZFExam)
}

func GetRoom(u *models.User, year, term, campus, weekday, week, sections, host string, loginType funnelApi.LoginType) (interface{}, error) {
	form := genTermForm(u, year, term, loginType)
	form.Add("campus", campus)
	form.Add("weekday", weekday)
	form.Add("week", week)
	form.Add("sections", sections)
	return FetchHandleOfPost(form, host, funnelApi.ZFRoom)
}

func BindPassword(u *models.User, year, term, host string, loginType funnelApi.LoginType) (interface{}, error) {
	var password string
	if loginType == "ZF" {
		password = u.ZFPassword
	} else if loginType == "OAUTH" {
		password = u.OauthPassword
	}
	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", password)
	form.Add("type", string(loginType))
	form.Add("year", year)
	form.Add("term", term)
	return FetchHandleOfPost(form, host, funnelApi.ZFExam)
}
