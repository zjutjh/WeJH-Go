package funnelServices

import (
	"net/url"
	"wejh-go/app/models"
)

func GetCurrentBorrow(u *models.User) (interface{}, error) {
	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", u.LibPassword)
	return FetchHandle(form, "student/library/borrow/current")
}

func GetHistoryBorrow(u *models.User) (interface{}, error) {
	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", u.LibPassword)
	return FetchHandle(form, "student/library/borrow/history")
}
