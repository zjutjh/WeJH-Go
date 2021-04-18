package funnelServices

import (
	"net/url"
	"wejh-go/app/models"
	"wejh-go/config/api/funnelApi"
)

func GetCurrentBorrow(u *models.User) (interface{}, error) {
	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", u.LibPassword)
	return FetchHandle(form, funnelApi.LibraryCurrent)
}

func GetHistoryBorrow(u *models.User) (interface{}, error) {
	form := url.Values{}
	form.Add("username", u.StudentID)
	form.Add("password", u.LibPassword)
	return FetchHandle(form, funnelApi.LibraryHistory)
}
