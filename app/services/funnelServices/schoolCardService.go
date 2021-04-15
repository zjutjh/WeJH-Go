package funnelServices

import (
	"net/url"
	"strconv"
	"wejh-go/app/models"
	"wejh-go/config/api/funnelApi"
)

func GetCardBalance(u *models.User) (interface{}, error) {
	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", u.CardPassword)
	return FetchHandle(form, "student/card/balance")
}

func GetCardToday(u *models.User) (interface{}, error) {
	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", u.CardPassword)
	return FetchHandle(form, "student/card/today")
}

func GetCardHistory(u *models.User, year, month int) (interface{}, error) {
	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", u.CardPassword)
	form.Add("year", strconv.Itoa(year))
	form.Add("month", strconv.Itoa(month))
	return FetchHandle(form, funnelApi.CardHistory)
}
