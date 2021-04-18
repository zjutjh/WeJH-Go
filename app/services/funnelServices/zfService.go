package funnelServices

import (
	"net/url"
	"wejh-go/app/models"
	"wejh-go/config/api/funnelApi"
)

func genTermForm(u *models.User, year, term string) url.Values {
	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", u.ZFPassword)
	form.Add("year", year)
	form.Add("term", term)
	return form
}

func GetClassTable(u *models.User, year, term string) (interface{}, error) {
	form := genTermForm(u, year, term)
	return FetchHandle(form, funnelApi.ZFClassTable)
}

func GetScore(u *models.User, year, term string) (interface{}, error) {
	form := genTermForm(u, year, term)
	return FetchHandle(form, funnelApi.ZFScore)
}

func GetExam(u *models.User, year, term string) (interface{}, error) {
	form := genTermForm(u, year, term)
	return FetchHandle(form, funnelApi.ZFExam)
}
