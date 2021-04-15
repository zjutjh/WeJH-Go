package funnelServices

import (
	"net/url"
	"wejh-go/app/models"
)

func GetClassTable(u *models.User) (interface{}, error) {

	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", u.ZFPassword)
	form.Add("year", "2020")
	form.Add("term", "3")
	return FetchHandle(form, "student/zf/table")
}
