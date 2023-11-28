package funnelServices

import (
	"math/rand"
	"net/url"
	"time"
	"wejh-go/app/apiException"
	"wejh-go/app/models"
	"wejh-go/config/api/funnelApi"
)

func genTermForm(u *models.User, year, term string) url.Values {
	var password, loginType string

	rand.Seed(time.Now().UnixNano())
	oauthVal := rand.Intn(40)
	zfVal := rand.Intn(60)

	if u.OauthPassword != "" && u.ZFPassword != "" {
		if oauthVal > zfVal {
			password = u.OauthPassword
			loginType = "OAUTH"
		} else {
			password = u.ZFPassword
			loginType = "ZF"
		}
	} else if u.OauthPassword != "" {
		password = u.OauthPassword
		loginType = "OAUTH"
	} else {
		password = u.ZFPassword
		loginType = "ZF"
	}
	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", password)
	form.Add("type", loginType)
	form.Add("year", year)
	form.Add("term", term)
	return form
}

func GetClassTable(u *models.User, year, term string) (interface{}, error) {
	if u.ZFPassword == "" && u.OauthPassword == "" {
		return nil, apiException.NoThatPasswordOrWrong
	}
	form := genTermForm(u, year, term)
	return FetchHandleOfPost(form, funnelApi.ZFClassTable)
}

func GetScore(u *models.User, year, term string) (interface{}, error) {
	if u.ZFPassword == "" && u.OauthPassword == "" {
		return nil, apiException.NoThatPasswordOrWrong
	}
	form := genTermForm(u, year, term)
	return FetchHandleOfPost(form, funnelApi.ZFScore)
}

func GetMidTermScore(u *models.User, year, term string) (interface{}, error) {
	if u.ZFPassword == "" && u.OauthPassword == "" {
		return nil, apiException.NoThatPasswordOrWrong
	}
	form := genTermForm(u, year, term)
	return FetchHandleOfPost(form, funnelApi.ZFMidTermScore)
}

func GetExam(u *models.User, year, term string) (interface{}, error) {
	if u.ZFPassword == "" && u.OauthPassword == "" {
		return nil, apiException.NoThatPasswordOrWrong
	}
	form := genTermForm(u, year, term)
	return FetchHandleOfPost(form, funnelApi.ZFExam)
}

func GetRoom(u *models.User, year, term, campus, weekday, week, sections string) (interface{}, error) {
	if u.ZFPassword == "" && u.OauthPassword == "" {
		return nil, apiException.NoThatPasswordOrWrong
	}
	form := genTermForm(u, year, term)
	form.Add("campus", campus)
	form.Add("weekday", weekday)
	form.Add("week", week)
	form.Add("sections", sections)
	return FetchHandleOfPost(form, funnelApi.ZFRoom)
}

func BindPassword(u *models.User, year, term, loginType string) (interface{}, error) {
	var password string
	if loginType == "ZF" {
		password = u.ZFPassword
	} else if loginType == "OAUTH" {
		password = u.OauthPassword
	}
	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", password)
	form.Add("type", loginType)
	form.Add("year", year)
	form.Add("term", term)
	return FetchHandleOfPost(form, funnelApi.ZFExam)
}
